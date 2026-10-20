package aqdata

import (
	"fmt"

	"github.com/tinylib/msgp/msgp"
)

// Preset define single preset
type Preset struct {
	Index int
	Name  string
}

// ToString gets the name of the preset
func (p Preset) ToString() string {
	return p.Name
}

// GetIndex gets the index for preset
func (p Preset) GetIndex() (int, bool) {
	if p.Name == "" {
		return 0, false
	}
	return p.Index, true
}

// PresetLike defines type constraint for preset
type PresetLike interface {
	~struct{ Preset }
	ToString() string
	GetIndex() (int, bool)
}

// IsPresetsEqual checks if two presets are the same
func IsPresetsEqual[T PresetLike](
	preset1 *T,
	preset2 *T,
) bool {
	if preset1 == nil && preset2 == nil {
		return true
	}
	if preset1 != nil && preset2 != nil {
		index1, hasName1 := (*preset1).GetIndex()
		index2, hasName2 := (*preset1).GetIndex()
		return (*preset1).ToString() == (*preset2).ToString() &&
			index1 == index2 &&
			hasName1 == hasName2
	}
	return false
}

// PresetList defines list of presets accepting generic preset type
type PresetList[T PresetLike, _ msgp.RTFor[T]] struct {
	Presets []T
	ByName  map[string]T
}

// define a new preset
func (l *PresetList[T, _]) define(name string) T {
	if l.ByName == nil {
		// init byName map if it is nil
		l.ByName = make(map[string]T)
	}
	if _, ok := l.ByName[name]; ok {
		// enough to panic
		panic("duplicate preset " + name)
	}
	// create a new preset based on the preset type
	p := T{Preset{len(l.Presets), name}}
	l.Presets = append(l.Presets, p)
	l.ByName[name] = p
	return p
}

// Length gets the length of the preset list
func (l *PresetList[T, _]) Length() int { return len(l.Presets) }

// FromIndex gets the preset by index
func (l *PresetList[T, _]) FromIndex(i int) (T, error) {
	// range check
	if i < 0 || i >= l.Length() {
		var zero T
		return zero, fmt.Errorf("%d is out of range 0-%d", i, l.Length()-1)
	}
	return l.Presets[i], nil
}

// FromName gets the preset by its name
func (l *PresetList[T, _]) FromName(name string) (T, error) {
	// lookup the preset by name
	p, ok := l.ByName[name]
	if !ok {
		var zero T
		return zero, fmt.Errorf("preset %s not found", name)
	}
	return p, nil
}

// GetPresets get the presets up to the limit
func (l *PresetList[T, _]) GetPresets(limit int) []T {
	limit = min(max(limit, 0), l.Length())
	return l.Presets[:limit]
}
