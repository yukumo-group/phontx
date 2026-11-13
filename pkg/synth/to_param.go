package synth

import (
	"bytes"
	"encoding/binary"
	"errors"

	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// PackData packs synth param to byte buffer
func PackData(
	phontData uint32,
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
		synthParam.Speed,
	)
	if err != nil {
		return nil, err
	}
	// phont_data
	err = binary.Write(
		result,
		binary.LittleEndian,
		phontData,
	)
	if err != nil {
		return nil, err
	}
	// Tone
	var base []byte
	var customData []byte
	var convertTone *custom.Aq10CustomTone
	if tone.Aq10CustomTone != nil {
		convertTone = tone.Aq10CustomTone
	} else {
		convertTone = custom.NewAq10CustomTone().WithFSC(
			100,
		).WithLMD(
			100,
		)
	}
	base, customData, err = convertTone.PackData()
	if err != nil {
		return nil, err
	}
	// base
	err = binary.Write(
		result,
		binary.LittleEndian,
		base,
	)
	if err != nil {
		return nil, err
	}
	// Volume
	err = binary.Write(
		result,
		binary.LittleEndian,
		synthParam.Volume,
	)
	if err != nil {
		return nil, err
	}
	// tone
	err = binary.Write(
		result,
		binary.LittleEndian,
		customData,
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
	if err != nil {
		return nil, err
	}
	return result.Bytes(), nil
}
