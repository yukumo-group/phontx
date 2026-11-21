package tone

import (
	"errors"
	"os"

	"github.com/tinylib/msgp/msgp"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"gopkg.in/yaml.v3"
)

// RawPreset defines the raw preset for generation
type RawPreset struct {
	PresetIdx  *int    `yaml:"preset_idx"`
	PresetName *string `yaml:"preset_name"`
}

// GetPresetFromRawPreset gets preset from raw preset
func (raw *RawPreset) GetPresetFromRawPreset[T aqdata.PresetLike, PT msgp.RTFor[T]](
	presetList *aqdata.PresetList[T, PT],
) (*T, error) {
	if raw.PresetName != nil {
		fetchedPreset, err := presetList.FromName(
			*raw.PresetName,
		)
		if err != nil {
			return nil, err
		}
		return &fetchedPreset, nil
	}
	if raw.PresetIdx != nil {
		fetchedPreset, err := presetList.FromIndex(
			*raw.PresetIdx,
		)
		if err != nil {
			return nil, err
		}
		return &fetchedPreset, nil
	}
	return nil, errors.New(
		"cannot convert to preset as no preset can be found through index or name",
	)
}

// RawTone defines the raw tone from yaml to phontx
type RawTone struct {
	Aq1Preset           *RawPreset             `yaml:"aq1_preset"`
	Aq2Preset           *RawPreset             `yaml:"aq2_preset"`
	Aq10Preset          *RawPreset             `yaml:"aq10_preset"`
	Aq10CustomTone      *custom.Aq10CustomTone `yaml:"aq1_custom_tone"`
	SourcePhontFilePath *string                `yaml:"source_phont_file_path"`
}

// NewRawToneFromYAML converts data in yaml file to raw tone
func NewRawToneFromYAML(
	filePath string,
) (*RawTone, error) {
	_, err := os.Stat(filePath)
	if err != nil {
		return nil, err
	}
	bytesData, err := os.ReadFile(
		filePath,
	)
	if err != nil {
		return nil, err
	}
	return NewRawToneFromBinary(
		bytesData,
	)
}

// FromBinary converts yaml binary data to raw tone
func NewRawToneFromBinary(
	bytesData []byte,
) (*RawTone, error) {
	rawData := &RawTone{}
	err := yaml.Unmarshal(
		bytesData,
		&rawData,
	)
	return rawData, err
}

// FromRawToneToPhontx converts raw tone to phontx tone
func (raw *RawTone) FromRawToneToPhontx() (*Data, error) {
	data := NewData()
	if raw.Aq1Preset != nil {
		aq1PresetSelected, err := raw.Aq1Preset.GetPresetFromRawPreset(
			aqdata.Aq1Presets,
		)
		if err != nil {
			return nil, err
		}
		if aq1PresetSelected != nil {
			data.SetAq1Preset(
				*aq1PresetSelected,
			)
		}
	}
	if raw.Aq2Preset != nil {
		aq2PresetSelected, err := raw.Aq2Preset.GetPresetFromRawPreset(
			aqdata.Aq2Presets,
		)
		if err != nil {
			return nil, err
		}
		if aq2PresetSelected != nil {
			data.SetAq2Preset(
				*aq2PresetSelected,
			)
		}
	}
	if raw.Aq10Preset != nil {
		aq10PresetSelected, err := raw.Aq10Preset.GetPresetFromRawPreset(
			aqdata.Aq10Presets,
		)
		if err != nil {
			return nil, err
		}
		if aq10PresetSelected != nil {
			data.SetAq10Preset(
				*aq10PresetSelected,
			)
		}
	}
	data.SetAq10CustomTone(
		raw.Aq10CustomTone,
	)
	return data, nil
}

// RawYAMLToPhontx converts raw yaml binary data to phontx
func RawYAMLBinaryToPhontx(
	binaryData []byte,
) (*Data, error) {
	rawTone, err := NewRawToneFromBinary(
		binaryData,
	)
	if err != nil {
		return nil, err
	}
	return rawTone.FromRawToneToPhontx()
}

// RawYAMLFileToPhontx converts raw yaml file to phontxx
func RawYAMLFileToPhontx(
	filePath string,
) (*Data, error) {
	rawTone, err := NewRawToneFromYAML(
		filePath,
	)
	if err != nil {
		return nil, err
	}
	return rawTone.FromRawToneToPhontx()
}
