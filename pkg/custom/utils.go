package custom

// GetDefaultPitchAndAccent gets default pitch and accent.
// pitch + accent.
func GetDefaultPitchAndAccent(
	base Aq10Base,
) (uint32, uint32) {
	switch base {
	case F1E:
		return 100, 100
	case F2E:
		return 77, 150
	case M1E:
		return 30, 100
	default:
		return 100, 100
	}
}

// GetPitch gets the default pitch or value set
func GetPitch(
	base Aq10Base,
	pitch *uint32,
) uint32 {
	if pitch == nil {
		pitchVal, _ := GetDefaultPitchAndAccent(
			base,
		)
		return pitchVal
	}
	return *pitch
}

// GetAccent gets the default accent or value set
func GetAccent(
	base Aq10Base,
	accent *uint32,
) uint32 {
	if accent == nil {
		_, accentVal := GetDefaultPitchAndAccent(
			base,
		)
		return accentVal
	}
	return *accent
}
