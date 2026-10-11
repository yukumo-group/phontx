package custom

import "fmt"

// Aq10Base defines bases for aquestalk10
type Aq10Base int

const (
	// F1E defines the F1E bases
	F1E Aq10Base = iota
	// F2E defines the F2E bases
	F2E
	// M1E defines the F1E bases
	M1E
)

// ToAq10Base converts int to Aq10Bases.
// Range: 0-2.
func ToAq10Base(
	data int,
) (Aq10Base, error) {
	if data > 2 || data < 0 {
		return F1E, fmt.Errorf(
			"%d is not in the range of 0-2",
			data,
		)
	}
	return Aq10Base(data), nil
}

// GetAllAq10Bases gets all the bases for aquestalk10
// Range: 0-3.
func GetAllAq10Bases(
	numBasesFetched int,
) ([]Aq10Base, error) {
	result := []Aq10Base{}
	realNum := min(max(numBasesFetched, 0), 3)
	for i := range realNum {
		tmpBase, err := ToAq10Base(
			i,
		)
		if err != nil {
			return result, err
		}
		result = append(result, tmpBase)
	}
	return result, nil
}

// ToString converts base to string
func (base Aq10Base) ToString() string {
	switch base {
	case F1E:
		return "F1E"
	case F2E:
		return "F2E"
	case M1E:
		return "M1E"
	default:
		return "NotSupported"
	}
}
