package synth

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/yukumo-group/phontx/internal/bts"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// Generator defines a single generator.
// YOU MUST CLOSE GENERATOR AFTER USE!!!!
type Generator struct {
	lock   sync.RWMutex
	mod    api.Module
	memory api.Memory
}

// NewGenerator creates new generator
func NewGenerator(
	ctx context.Context,
	generatorRuntime *GeneratorRuntime,
) (*Generator, error) {
	generatorModule, err := generatorRuntime.runtime.InstantiateModule(
		ctx,
		generatorRuntime.compiledGeneratorModule,
		wazero.NewModuleConfig().WithName(""),
	)
	if err != nil {
		return nil, err
	}
	generatorMemory := generatorModule.Memory()
	if generatorMemory == nil {
		return nil, errors.New(
			"memory not exported",
		)
	}
	return &Generator{
		mod:    generatorModule,
		memory: generatorMemory,
	}, nil
}

// GetByteInMemory gets single byte in memory
func (generator *Generator) GetByteInMemory(
	ctx context.Context,
	i uint32,
) (byte, error) {
	data, ok := generator.memory.ReadByte(i)
	if !ok {
		return byte(0), fmt.Errorf(
			"failed to get byte in memory position %d",
			i,
		)
	}
	return data, nil
}

// FreeWavData frees .wav data
func (generator *Generator) FreeWavData(
	ctx context.Context,
	wavOffset uint64,
) error {
	const freeWavFunctionName string = "aq_free"
	freeWavFunction := generator.mod.ExportedFunction(
		freeWavFunctionName,
	)
	if freeWavFunction == nil {
		return fmt.Errorf(
			"failed to get function %s",
			freeWavFunctionName,
		)
	}
	_, err := freeWavFunction.Call(
		ctx,
		wavOffset,
	)
	return err
}

// GetWavData gets wav data
func (generator *Generator) GetWavData(
	outPos uint32,
	outLen uint32,
	fileName string,
) (uint32, []byte, error) {
	outData, ok := generator.memory.Read(
		outPos,
		outLen+4,
	)
	if !ok {
		return 0, nil, fmt.Errorf(
			"failed to read data from %d with len %d",
			outPos,
			outLen+4,
		)
	}
	newBuffer := bytes.NewBuffer(outData)
	nextData := newBuffer.Next(8)
	outDataStartPosInBytes := nextData[:4]
	outDataLengthInBytes := nextData[4:9]
	outDataStartPos := binary.LittleEndian.Uint32(
		outDataStartPosInBytes,
	)
	outDataLength := binary.LittleEndian.Uint32(
		outDataLengthInBytes,
	)
	wavData, ok := generator.memory.Read(
		outDataStartPos,
		outDataLength,
	)
	if !ok {
		return outDataStartPos, nil, fmt.Errorf(
			"failed to read data from %d with len %d",
			outDataStartPos,
			outDataLength,
		)
	}
	return outDataStartPos, wavData, nil
}

// GiveError gives error from the function given by wasm module
func (generator *Generator) GiveError(
	ctx context.Context,
	version aqdata.AquesVersion,
	code uint64,
) error {
	if code == 0 {
		return nil
	}
	const errorMessageFunctionName string = "aq_error_message"
	giveErrorExportedFunction := generator.mod.ExportedFunction(
		errorMessageFunctionName,
	)
	if giveErrorExportedFunction == nil {
		return fmt.Errorf(
			"%s is not successfully exported",
			errorMessageFunctionName,
		)
	}
	datas, err := giveErrorExportedFunction.Call(
		ctx,
		uint64(version),
		code,
	)
	if err != nil {
		return err
	}
	if len(datas) < 1 {
		return errors.New(
			"returned no data",
		)
	}
	stringData, err := bts.ToString(
		ctx,
		uint32(datas[0]),
		generator.GetByteInMemory,
	)
	if err != nil {
		return err
	}
	return errors.New(
		stringData,
	)
}

// GenerateBinaryData generates binary data
func (generator *Generator) GenerateBinaryData(
	ctx context.Context,
	version aqdata.AquesVersion,
	param *Param,
	toneData *tone.Data,
) ([]byte, error) {
	generator.lock.Lock()
	defer generator.lock.Unlock()
	const synthFunctionName string = "aq_synthesize"
	// Exports function
	exportedSynthFunction := generator.mod.ExportedFunction(
		synthFunctionName,
	)
	if exportedSynthFunction == nil {
		return nil, fmt.Errorf(
			"function %s does not exist",
			synthFunctionName,
		)
	}
	// Checks if version is allowed
	versionAllowed, err := toneData.IsVersionAllowed(
		version,
	)
	if err != nil {
		return nil, err
	}
	if !versionAllowed {
		return nil, fmt.Errorf(
			"version %s is not allowed",
			version.ToString(),
		)
	}
	// Save phont data
	var phontData []byte
	if toneData.PhontFileData == nil {
		phontData = []byte{}
	} else {
		phontData = toneData.PhontFileData
	}
	phontDataOffset := uint32(0)
	ok := generator.memory.Write(
		phontDataOffset,
		phontData,
	)
	lengthOfPhontData := len(phontData)
	if !ok {
		return nil, fmt.Errorf(
			"data of length %d out of range",
			lengthOfPhontData,
		)
	}
	// Pack param
	packedData, err := PackData(
		phontDataOffset,
		version,
		toneData,
		param,
	)
	if err != nil {
		return nil, err
	}
	packedDataOffset := bts.NextOffset(
		phontDataOffset,
		uint32(lengthOfPhontData),
	)
	ok = generator.memory.Write(
		packedDataOffset,
		packedData,
	)
	lengthOfPackedData := len(packedData)
	if !ok {
		return nil, fmt.Errorf(
			"data, length of %d, starting from offset %d, is out of range",
			lengthOfPackedData,
			packedDataOffset,
		)
	}
	// text
	textConvertedToBytes := []byte(param.Text)
	textOffset := bts.NextOffset(
		packedDataOffset,
		uint32(lengthOfPackedData),
	)
	ok = generator.memory.Write(
		textOffset,
		textConvertedToBytes,
	)
	lengthOfText := len(textConvertedToBytes)
	if !ok {
		return nil, fmt.Errorf(
			"data, length of %d, starting from offset %d, is out of range",
			lengthOfText,
			textOffset,
		)
	}
	// Out data definition
	outDataOffset := bts.NextOffset(
		textOffset,
		uint32(lengthOfText),
	)
	outLen := outDataOffset + 4
	// Call Generation function
	returnedData, err := exportedSynthFunction.Call(
		ctx,
		uint64(textOffset),
		uint64(lengthOfText),
		uint64(packedDataOffset),
		uint64(lengthOfPhontData),
		uint64(outDataOffset),
		uint64(outLen),
	)
	if err != nil {
		return nil, err
	}
	if len(returnedData) < 1 {
		return nil, errors.New(
			"no data returned",
		)
	}
	err = generator.GiveError(
		ctx,
		version,
		returnedData[0],
	)
	if err != nil {
		return nil, err
	}
	// get wav data and save
	outOffset, wavData, err := generator.GetWavData(
		outDataOffset,
		outLen,
		param.FilePath,
	)
	if err != nil {
		return wavData, err
	}
	err = generator.FreeWavData(
		ctx,
		uint64(outOffset),
	)
	if err != nil {
		return wavData, err
	}
	return wavData, nil
}

// Generate generates an audio
func (generator *Generator) Generate(
	ctx context.Context,
	version aqdata.AquesVersion,
	param *Param,
	toneData *tone.Data,
) error {
	binaryData, err := generator.GenerateBinaryData(
		ctx,
		version,
		param,
		toneData,
	)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		param.FilePath,
		binaryData,
		0644,
	)
	return err
}

// CloseGenerator closes generator in memory
func (generator *Generator) CloseGenerator(
	ctx context.Context,
) error {
	return generator.mod.Close(
		ctx,
	)
}
