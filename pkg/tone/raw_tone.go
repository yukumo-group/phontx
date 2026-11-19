package tone

// RawPreset defines the raw preset for generation
type RawPreset struct {
	PresetIdx  *int
	RresetName *string
}

// RawTone defines the raw tone from yaml to phontx
type RawTone struct {
	Aq1Preset  *RawPreset
	Aq2Preset  *RawPreset
	Aq10Preset *RawPreset
}
