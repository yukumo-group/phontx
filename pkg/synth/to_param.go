package synth

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// PackData packs synth param to byte buffer
func PackData(
	version aqdata.AquesVersion,
	tone *tone.Data,
	synthParam *SynthParam,
) ([]byte, error) {
	if tone == nil || synthParam == nil {
		return nil, errors.New(
			"you cannot input nil tone or synth param",
		)
	}
	result := new(bytes.Buffer)
	// pack version
	err := binary.Write(
		result,
		binary.LittleEndian,
		version.ToInt(),
	)
	if err != nil {
		return nil, err
	}
	// preset id
	presetID, err := tone.GetPresetID(
		version,
	)
	if err != nil {
		return nil, err
	}
	err = binary.Write(
		result,
		binary.LittleEndian,
		presetID,
	)
	if err != nil {
		return nil, err
	}
	// speed
	err = binary.Write(
		result,
		binary.LittleEndian,
		synthParam.Volume,
	)
	// Tone
	if tone.Aq10CustomTone != nil {
		tone.Aq10CustomTone.PackData()
	}
	return result.Bytes(), nil
}
