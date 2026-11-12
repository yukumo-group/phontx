package synth

// SynthParam defines the parameter for synthesis
type SynthParam struct {
	Speed      int32
	Text       string
	textLength int32
	Volume     int32
}

// NewSynthParam creates new param for synthesis audio
func NewSynthParam(
	text string,
	speed int32,
	volume int32,
) *SynthParam {
	textLength := int32(len([]rune(text)))
	return &SynthParam{
		Speed:      speed,
		Text:       text,
		textLength: textLength,
		Volume:     volume,
	}
}
