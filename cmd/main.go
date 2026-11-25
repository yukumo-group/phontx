package main

import (
	"context"

	"github.com/yukumo-group/phontx/pkg/synth"
)

var runtime *synth.GeneratorRuntime

func main() {
	ctx := context.Background()
	runtime = synth.NewGeneratorRuntime()
	err := runtime.InitializeRuntime(
		ctx,
		"./aq_wasm.wasm",
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
