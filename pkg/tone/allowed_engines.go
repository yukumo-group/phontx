package tone

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
