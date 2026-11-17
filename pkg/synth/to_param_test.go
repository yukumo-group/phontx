package synth

import (
	"encoding/hex"
	"testing"

	"github.com/yukumo-group/phontx/internal/bts"
	"github.com/yukumo-group/phontx/internal/python"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// TestToParam tests the packing of param to byte buffer
func TestToParam(t *testing.T) {
	t.Parallel()
	// Execute python script and get the data wanted
	const scriptPath string = "testdata/testdata.py"
	pythonResult, err := python.ExecutePython(
		scriptPath,
	)
	if err != nil {
		t.Fatal(err)
	}
	decodedData, err := hex.DecodeString(
		string(pythonResult),
	)
	if err != nil {
		t.Fatal(err)
	}
	// Prepare data
	synthParam := NewSynthParam(
		"114514",
		60,
		100,
	)
	customTone := custom.NewAq10CustomTone().WithAccent(
		30,
	).WithPitch(
		70,
	).WithFSC(
		100,
	).WithLMD(
		10,
	).WithBase(
		custom.F1E,
	)
	phontData := tone.NewData().SetAq10CustomTone(
		customTone,
	).SetAq10Preset(
		aqdata.Aq10F1,
	)
	packedResult, err := PackData(
		0,
		aqdata.AqVer10,
		phontData,
		synthParam,
	)
	if err != nil {
		t.Fatal(err)
	}
	isTwoBytesSame := bts.IsBytesSame(
		packedResult,
		decodedData,
	)
	if !isTwoBytesSame {
		t.Errorf(
			"\n%v (p)\n%v (d)",
			packedResult,
			decodedData,
		)
		t.Fatal(
			"fetched data is not the same",
		)
	}
}
