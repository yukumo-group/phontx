package synth

// Param defines the parameter for synthesis
type Param struct {
	Speed      uint32
	Text       string
	textLength uint32
	Volume     uint32
	FilePath   string
}

// NewSynthParam creates new param for synthesis audio
func NewSynthParam(
	text string,
	speed uint32,
	volume uint32,
	filePath string,
) *Param {
	textLength := uint32(len([]byte(text)))
	return &Param{
		Speed:      min(max(speed, 50), 300),
		Text:       text,
		textLength: textLength,
		Volume:     min(max(volume, 0), 300),
		FilePath:   filePath,
	}
}
