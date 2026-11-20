package main

import (
	"context"

	"github.com/yukumo-group/phontx/pkg/synth"
)

var runtime *synth.GeneratorRuntime

func main() {
	ctx := context.Background()
	runtime = synth.NewGeneratorRuntime(
		"./aq_wasm.wasm",
	)
	err := runtime.InitializeRuntime(
		ctx,
	)
	if err != nil {
		panic(err)
	}
	err = Execute()
	if err != nil {
		panic(err)
	}
	err = runtime.CloseRuntime(
		ctx,
	)
	if err != nil {
		panic(err)
	}
}
