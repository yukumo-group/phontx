package synth

import (
	"context"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
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
	return &Generator{
		mod: generatorModule,
	}, nil
}

// AllocateMemory allocates memory
func (generator *Generator) AllocateMemory(
	startPos int,
) error {
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
