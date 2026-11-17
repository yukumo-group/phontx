package synth

import (
	"context"
	"errors"
	"fmt"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
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

// Generate generates an audio
func (generator *Generator) Generate(
	ctx context.Context,
	version aqdata.AquesVersion,
	param *Param,
	toneData *tone.Data,
) error {
	const phontDataName string = "phont_data"
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
	if !ok {
		return fmt.Errorf(
			"data of length %d out of range",
			len(phontData),
		)
	}
	return nil
}

// CloseGenerator closes generator in memory
func (generator *Generator) CloseGenerator(
	ctx context.Context,
) error {
	return generator.mod.Close(
		ctx,
	)
}
