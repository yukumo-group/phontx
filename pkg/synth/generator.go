package synth

import (
	"bytes"
	"context"
	"encoding/binary"
	"errors"
	"fmt"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/yukumo-group/phontx/internal/bts"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// Generator defines a single generator.
// YOU MUST CLOSE GENERATOR AFTER USE!!!!
type Generator struct {
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

// GetWavData gets wav data
func (generator *Generator) GetWavData(
	outPos uint32,
	outLen uint32,
	fileName string,
) ([]byte, error) {
	outData, ok := generator.memory.Read(
		outPos,
		outLen+4,
	)
	if !ok {
		return nil, fmt.Errorf(
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
		return nil, fmt.Errorf(
			"failed to read data from %d with len %d",
			outDataStartPos,
			outDataLength,
		)
	}
	return wavData, nil
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

// Generate generates an audio
func (generator *Generator) Generate(
	ctx context.Context,
	version aqdata.AquesVersion,
	param *Param,
	toneData *tone.Data,
) error {
	const synthFunctionName string = "aq_synthesize"
	// Exports function
	exportedSynthFunction := generator.mod.ExportedFunction(
		synthFunctionName,
	)
	if exportedSynthFunction == nil {
		return fmt.Errorf(
			"function %s does not exist",
			synthFunctionName,
		)
	}
	// Checks if version is allowed
	versionAllowed, err := toneData.IsVersionAllowed(
		version,
	)
	if err != nil {
		return err
	}
	if !versionAllowed {
		return fmt.Errorf(
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
		return fmt.Errorf(
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
		return err
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
		return fmt.Errorf(
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
		return fmt.Errorf(
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
		return err
	}
	if len(returnedData) < 1 {
		return errors.New(
			"no data returned",
		)
	}
	err = generator.GiveError(
		ctx,
		version,
		returnedData[0],
	)
	if err != nil {
		return err
	}
	// get wav data and save
	wavData, err := generator.GetWavData(
		outDataOffset,
		outLen,
		param.FilePath,
	)
	if err != nil {
		return err
	}
	err = os.WriteFile(
		param.FilePath,
		wavData,
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
