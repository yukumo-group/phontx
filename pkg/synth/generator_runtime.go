package synth

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero"
)

// GeneratorRuntime defines the runtime for generator
type GeneratorRuntime struct {
	generatorWasmFilePath   string
	runtime                 wazero.Runtime
	compiledGeneratorModule wazero.CompiledModule
}

// NewGeneratorRuntime creates new generator runtime
func NewGeneratorRuntime(
	generatorWasmFilePath string,
) *GeneratorRuntime {
	return &GeneratorRuntime{
		generatorWasmFilePath: generatorWasmFilePath,
	}
}

// InitializeRuntime creates runtime
func (generator *GeneratorRuntime) InitializeRuntime(
	ctx context.Context,
) error {
	generator.runtime = wazero.NewRuntime(ctx)
	_, err := os.Stat(generator.generatorWasmFilePath)
	if err != nil {
		return err
	}
	wasmGeneratorByte, err := os.ReadFile(
		generator.generatorWasmFilePath,
	)
	if err != nil {
		return err
	}
	compiledModule, err := generator.runtime.CompileModule(
		ctx,
		wasmGeneratorByte,
	)
	if err != nil {
		return err
	}
	generator.compiledGeneratorModule = compiledModule
	return nil
}

// CloseRuntime closes the runtime
func (generator *GeneratorRuntime) CloseRuntime(
	ctx context.Context,
) error {
	err := generator.compiledGeneratorModule.Close(
		ctx,
	)
	if err != nil {
		return err
	}
	return generator.runtime.Close(
		ctx,
	)
}
