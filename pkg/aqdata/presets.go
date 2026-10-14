package aqdata

import (
	"fmt"
)

// Aq1Preset defines presets for aquestalk 1
type Aq1Preset int

const (
	// Aq1F1 -> f1 of aquestalk 1 library
	Aq1F1 Aq1Preset = iota
	// Aq1F2 -> f2 of aquestalk 1 library
	Aq1F2
	// Aq1F3 -> f3 of aquestalk 1 library
	Aq1F3
	// Aq1M1 -> m1 of aquestalk 1 library
	Aq1M1
	// Aq1M2 -> m2 of aquestalk 1 library
	Aq1M2
	// Aq1R1 -> r1 of aquestalk 1 library
	Aq1R1
	// Aq1DVD -> dvd of aquestalk 1 library
	Aq1DVD
	// Aq1IMD1 -> imd1 of aquestalk 1 library
	Aq1IMD1
	// Aq1JGR -> jgr of aquestalk 1 library
	Aq1JGR
)

// ToAq1Preset converts int to preset of aquestalk1
// It ranges from 0 to 8
func ToAq1Preset(data int) (Aq1Preset, error) {
	if data < 0 || data > 8 {
		return Aq1F1, fmt.Errorf(
			"%d is not in range of 0-8",
			data,
		)
	}
	return Aq1Preset(data), nil
}

// ToString converts Aq1Preset to string
func (preset Aq1Preset) ToString() string {
	switch preset {
	case Aq1F1:
		return "f1"
	case Aq1F2:
		return "f2"
	case Aq1F3:
		return "f3"
	case Aq1M1:
		return "m1"
	case Aq1M2:
		return "m2"
	case Aq1R1:
		return "r1"
	case Aq1DVD:
		return "dvd"
	case Aq1IMD1:
		return "imd1"
	case Aq1JGR:
		return "jgr"
	default:
		return "not_supported"
	}
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

// Aq2Preset defines preset for questalk2
type Aq2Preset int

const (
	// Aq2AqYukkuri -> aq_yukkuri
	Aq2AqYukkuri Aq2Preset = iota
	// Aq2AqDefo1 -> aq_defo1
	Aq2AqDefo1
	// Aq2AqF1c -> aq_f1c
	Aq2AqF1c
	// Aq2AqF3a -> aq_f3a
	Aq2AqF3a
	// Aq2AqHuskey -> aq_huskey
	Aq2AqHuskey
	// Aq2AqM4b -> aq_m4b
	Aq2AqM4b
	// Aq2AqMf1 -> aq_mf1
	Aq2AqMf1
	// Aq2AqRb2 -> aq_rb2
	Aq2AqRb2
	// Aq2AqRb3 -> aq_rb3
	Aq2AqRb3
	// Aq2AqRm -> aq_rm
	Aq2AqRm
	// Aq2AqRobo -> aq_robo
	Aq2AqRobo
	// Aq2AqTeto1 -> aq_teto1
	Aq2AqTeto1
	// Aq2ArF4 -> ar_f4
	Aq2ArF4
	// Aq2ArM5 -> ar_m5
	Aq2ArM5
	// Aq2ArMF2 -> ar_mf2
	Aq2ArMF2
	// Aq2ArRM3 -> ar_rm3
	Aq2ArRM3
)

// ToAq2Preset converts in to preset of aquestalk2.
// It ranges from 0 to 15.
func ToAq2Preset(
	data int,
) (Aq2Preset, error) {
	if data > 15 || data < 0 {
		return Aq2AqYukkuri, fmt.Errorf(
			"%d is not in range of 0-15",
			data,
		)
	}
	return Aq2Preset(data), nil
}

// ToString converts presets to string
func (preset Aq2Preset) ToString() string {
	switch preset {
	case Aq2AqYukkuri:
		return "aq_yukkuri"
	case Aq2AqDefo1:
		return "aq_defo1"
	case Aq2AqF1c:
		return "aq_f1c"
	case Aq2AqF3a:
		return "aq_f3a"
	case Aq2AqHuskey:
		return "aq_huskey"
	case Aq2AqM4b:
		return "aq_m4b"
	case Aq2AqMf1:
		return "aq_mf1"
	case Aq2AqRb2:
		return "aq_rb2"
	case Aq2AqRb3:
		return "aq_rb3"
	case Aq2AqRm:
		return "aq_rm"
	case Aq2AqRobo:
		return "aq_robo"
	case Aq2AqTeto1:
		return "aq_teto1"
	case Aq2ArF4:
		return "ar_f4"
	case Aq2ArM5:
		return "ar_m5"
	case Aq2ArMF2:
		return "ar_mf2"
	case Aq2ArRM3:
		return "ar_rm3"
	default:
		return "not_supported"
	}
}

// GetAq2Presets gets presets for aquestalk2.
// numPresetsFetched is the total number of presets fetched
// its range is from 0 to 16
func GetAq2Presets(
	numPresetsFetched int,
) ([]Aq2Preset, error) {
	realNum := min(max(numPresetsFetched, 0), 16)
	result := []Aq2Preset{}
	for i := range realNum {
		data, err := ToAq2Preset(i)
		if err != nil {
			return result, err
		}
		result = append(result, data)
	}
	return result, nil
}

// Aq10Preset defines presets for aquestalk10
type Aq10Preset int

const (
	// Aq10Custom uses custom data
	Aq10Custom Aq10Preset = iota
	// Aq10F1 -> f1
	Aq10F1
	// Aq10F2 -> f2
	Aq10F2
	// Aq10F3 -> f3
	Aq10F3
	// Aq10M1 -> m1
	Aq10M1
	// Aq10M2 -> m2
	Aq10M2
	// Aq10R1 -> r1
	Aq10R1
	// Aq10R2 -> r2
	Aq10R2
)

// ToAq10Preset converts int to preset of aq10
// Ranges from 0-7
func ToAq10Preset(
	data int,
) (Aq10Preset, error) {
	if data < 0 || data > 7 {
		return Aq10Custom, fmt.Errorf(
			"%d is out of range 0-7",
			data,
		)
	}
	return Aq10Preset(data), nil
}

// ToString converts Aq10Preset to string
func (preset Aq10Preset) ToString() string {
	switch preset {
	case Aq10Custom:
		return "custom"
	case Aq10F1:
		return "f1"
	case Aq10F2:
		return "f2"
	case Aq10F3:
		return "f3"
	case Aq10M1:
		return "m1"
	case Aq10M2:
		return "m2"
	case Aq10R1:
		return "r1"
	case Aq10R2:
		return "r2"
	default:
		return "not_supported"
	}
}

// GetAq10Presets gets presets for aquestalk2.
// numPresetsFetched is the total number of presets fetched
// its range is from 0 to 8
func GetAq10Presets(
	numPresetsFetched int,
) ([]Aq10Preset, error) {
	realNum := min(max(numPresetsFetched, 0), 8)
	result := []Aq10Preset{}
	for i := range realNum {
		data, err := ToAq10Preset(i)
		if err != nil {
			return result, err
		}
		result = append(result, data)
	}
	return result, nil
}
