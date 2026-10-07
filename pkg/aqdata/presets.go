package aqdata

import (
	"fmt"
)

// Aq1Preset defines presets for aquestalk 1
type Aq1Preset int

const (
	// F1aq1 -> f1 of aquestalk 1 library
	F1aq1 Aq1Preset = iota
	// F2aq1 -> f2 of aquestalk 1 library
	F2aq1
	// F3aq1 -> f3 of aquestalk 1 library
	F3aq1
	// M1aq1 -> m1 of aquestalk 1 library
	M1aq1
	// M2aq1 -> m2 of aquestalk 1 library
	M2aq1
	// R1aq1 -> r1 of aquestalk 1 library
	R1aq1
	// DVDaq1 -> dvd of aquestalk 1 library
	DVDaq1
	// IMD1aq1 -> imd1 of aquestalk 1 library
	IMD1aq1
	// JGRaq1 -> jgr of aquestalk 1 library
	JGRaq1
)

// ToAq1Preset converts int to preset of aquestalk1
func ToAq1Preset(data int) (Aq1Preset, error) {
	if data < 0 || data > 8 {
		return F1aq1, fmt.Errorf(
			"%d is not in range of 0-8",
			data,
		)
	}
	return Aq1Preset(data), nil
}

// GetAq1Presets gets the intended number of presets
// numPresetsFetched is the total number of presets fetched
// its range is from 0 to 9
func GetAq1Presets(
	numPresetsFetched int,
) ([]Aq1Preset, error) {
	effectiveNumPresets := max(min(numPresetsFetched, 9), 0)
	result := []Aq1Preset{}
	for i := range effectiveNumPresets {
		data, err := ToAq1Preset(i)
		if err != nil {
			return result, err
		}
		result = append(result, data)
	}
	return result, nil
}
