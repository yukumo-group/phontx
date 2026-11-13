package synth

// SynthParam defines the parameter for synthesis
type SynthParam struct {
	Speed      uint32
	Text       string
	textLength uint32
	Volume     uint32
}

// NewSynthParam creates new param for synthesis audio
func NewSynthParam(
	text string,
	speed uint32,
	volume uint32,
) *SynthParam {
	textLength := uint32(len([]byte(text)))
	return &SynthParam{
		Speed:      speed,
		Text:       text,
		textLength: textLength,
		Volume:     volume,
	}
}
