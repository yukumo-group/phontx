package synth

import (
	"fmt"
	"testing"

	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// TestRuntime tests the runtime
func TestRuntime(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	const wasmPath string = "testdata/aq_wasm.wasm"
	testRuntime := NewGeneratorRuntime()
	err := testRuntime.InitializeRuntime(
		ctx,
		wasmPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testRuntime.CloseRuntime(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

// TestGenerator tests generator
func TestGenerator(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	const wasmPath string = "testdata/aq_wasm.wasm"
	// temp data
	tempDir := t.TempDir()
	filePath := fmt.Sprintf(
		"%s/test.wav",
		tempDir,
	)
	// Load wasm
	testRuntime := NewGeneratorRuntime()
	err := testRuntime.InitializeRuntime(
		ctx,
		wasmPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	// generation
	testGenerator, err := NewGenerator(
		ctx,
		testRuntime,
	)
	if err != nil {
		t.Fatal(err)
	}
	synthParam := NewSynthParam(
		"ゴーファーくん",
		100,
		100,
		filePath,
	)
	toneUsed := tone.NewData().SetAq1Preset(
		aqdata.Aq1DVD,
	).SetAq2Preset(
		aqdata.Aq2AqDefo1,
	).SetAq10Preset(
		aqdata.Aq10Custom,
	)
	err = testGenerator.Generate(
		ctx,
		aqdata.AqVer1,
		synthParam,
		toneUsed,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testGenerator.CloseGenerator(
		ctx,
	)
	if err != nil {
		t.Fatal(err)
	}
	err = testRuntime.CloseRuntime(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
