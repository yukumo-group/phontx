package tone

import (
	"errors"
	"fmt"

	"github.com/yukumo-group/phontx/pkg/aqdata"
)

// AllowedEngines defines engines allowed by a tone
type AllowedEngines struct {
	Aq1  bool
	Aq2  bool
	Aq10 bool
}

// GetAllowedEngines gets engines that are allowed
func (tone *Data) GetAllowedEngines() *AllowedEngines {
	Aq1Enabled := false
	Aq2Enabled := false
	Aq10Enabled := false
	if tone.Aq1Preset != nil {
		Aq1Enabled = true
	}
	if tone.Aq2Preset != nil || tone.PhontFileData != nil {
		Aq2Enabled = true
	}
	if tone.Aq10CustomTone != nil || tone.Aq10Preset != nil {
		Aq10Enabled = true
	}
	return &AllowedEngines{
		Aq1:  Aq1Enabled,
		Aq2:  Aq2Enabled,
		Aq10: Aq10Enabled,
	}
}

// IsVersionAllowed checks if specific version of engine is allowed
func (tone *Data) IsVersionAllowed(
	version aqdata.AquesVersion,
) (bool, error) {
	allowedEngines := tone.GetAllowedEngines()
	switch version {
	case aqdata.AqVer1:
		return allowedEngines.Aq1, nil
	case aqdata.AqVer2:
		return allowedEngines.Aq2, nil
	case aqdata.AqVer10:
		return allowedEngines.Aq10, nil
	default:
		return false, fmt.Errorf(
			"%d does not exists",
			version,
		)
	}
}

// GetPresetID gets id for preset
func (tone *Data) GetPresetID(
	version aqdata.AquesVersion,
) (int32, error) {
	versionAllowed, err := tone.IsVersionAllowed(version)
	if err != nil {
		return 0, err
	}
	if !versionAllowed {
		return 0, fmt.Errorf(
			"version %s is not supported",
			version.ToString(),
		)
	}
	switch version {
	case aqdata.AqVer1:
		if tone.Aq1Preset == nil {
			return 0, errors.New(
				"the preset is set to nil in aq1 preset field, which is not exxxpected",
			)
		}
		presetIndex, exists := tone.Aq1Preset.GetIndex()
		if !exists {
			return 0, errors.New(
				"aq1 preset does not exists",
			)
		}
		return presetIndex, nil
	case aqdata.AqVer2:
		if tone.Aq2Preset == nil {
			return 0, errors.New(
				"the preset is set to nil in aq2 preset field, which is not exxxpected",
			)
		}
		presetIndex, exists := tone.Aq2Preset.GetIndex()
		if !exists {
			return 0, errors.New(
				"aq2 preset does not exists",
			)
		}
		return presetIndex, nil
	case aqdata.AqVer10:
		if tone.Aq10Preset == nil {
			return 0, errors.New(
				"the preset is set to nil in aq10 preset field, which is not exxxpected",
			)
		}
		presetIndex, exists := tone.Aq10Preset.GetIndex()
		if !exists {
			return 0, errors.New(
				"aq10 preset does not exists",
			)
		}
		return presetIndex, nil
	default:
		return 0, fmt.Errorf(
			"version %s is not supported",
			version.ToString(),
		)
	}
}
