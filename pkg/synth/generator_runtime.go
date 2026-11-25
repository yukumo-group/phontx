package synth

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero"
)

// GeneratorRuntime defines the runtime for generator
type GeneratorRuntime struct {
	runtime                 wazero.Runtime
	compiledGeneratorModule wazero.CompiledModule
}

// NewGeneratorRuntime creates new generator runtime
func NewGeneratorRuntime() *GeneratorRuntime {
	return &GeneratorRuntime{}
}

// InitializeRuntimeFromBinary initializes runtime from binary data
func (generator *GeneratorRuntime) InitializeRuntimeFromBinary(
	ctx context.Context,
	binaryData []byte,
) error {
	compiledModule, err := generator.runtime.CompileModule(
		ctx,
		binaryData,
	)
	if err != nil {
		return err
	}
	generator.compiledGeneratorModule = compiledModule
	return nil
}

// InitializeRuntime creates runtime
func (generator *GeneratorRuntime) InitializeRuntime(
	ctx context.Context,
	generatorWasmFilePath string,
) error {
	generator.runtime = wazero.NewRuntime(ctx)
	_, err := os.Stat(generatorWasmFilePath)
	if err != nil {
		return err
	}
	wasmGeneratorByte, err := os.ReadFile(
		generatorWasmFilePath,
	)
	if err != nil {
		return err
	}
	return generator.InitializeRuntimeFromBinary(
		ctx,
		wasmGeneratorByte,
	)
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
