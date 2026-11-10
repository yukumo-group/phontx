package synth

import (
	"context"
	"os"

	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
)

// Generator defines the data struct for generating audio
type Generator struct {
	wasmFilePath string
	runtime      wazero.Runtime
	mod          api.Module
}

// NewGenerator creates new generator
func NewGenerator(
	wasmFilePath string,
) *Generator {
	return &Generator{
		wasmFilePath: wasmFilePath,
	}
}

// CreateRuntime creates runtime
func (generator *Generator) CreateRuntime(
	ctx context.Context,
) error {
	_, err := os.Stat(generator.wasmFilePath)
	if err != nil {
		return err
	}
	wasmData, err := os.ReadFile(generator.wasmFilePath)
	if err != nil {
		return err
	}
	generator.runtime = wazero.NewRuntime(ctx)
	generator.mod, err = generator.runtime.Instantiate(
		ctx,
		wasmData,
	)
	return err
}

// CloseRuntime closes the runtime
func (generator *Generator) CloseRuntime(
	ctx context.Context,
) {
	generator.CloseRuntime(
		ctx,
	)
}
