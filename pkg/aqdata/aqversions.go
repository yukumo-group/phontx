package aqdata

// AquesVersion defines the version of aquestalke
type AquesVersion int

const (
	// AqVer1 defines version of aquestalk1
	AqVer1 AquesVersion = 1
	// AqVer2 defines the version of aquestalk2
	AqVer2 AquesVersion = 2
	// AqVer10 defines the version of aquestalk10
	AqVer10 AquesVersion = 10
)

// ToAquesVersion converts integer to aques version
func ToAquesVersion(
	data int,
) AquesVersion {
	switch data {
	case 1:
		return AqVer1
	case 2:
		return AqVer2
	case 10:
		return AqVer10
	default:
		return AqVer2
	}
}

// ToString converts the version to string
func (version AquesVersion) ToString() string {
	switch version {
	case AqVer1:
		return "AquesTalk1"
	case AqVer2:
		return "AquesTalk2"
	case AqVer10:
		return "AquesTalk10"
	default:
		return "Not Defined"
	}
}

// ToInt converts the version to int
func (version AquesVersion) ToInt() int {
	switch version {
	case AqVer1:
		return 1
	case AqVer2:
		return 2
	case AqVer10:
		return 10
	default:
		return 2
	}
}

// GetAllVersions gets all the possible versions
func GetAllVersions() []AquesVersion {
	return []AquesVersion{
		AqVer1,
		AqVer2,
		AqVer10,
	}
}
