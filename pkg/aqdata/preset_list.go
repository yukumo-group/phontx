package aqdata

import "fmt"

type preset struct {
	index int
	name  string
}

func (p preset) ToString() string { return p.name }

func (p preset) Index() (int, bool) {
	if p.name == "" {
		return 0, false
	}
	return p.index, true
}

// type constraint for preset
type presetLike interface {
	~struct{ preset }
	ToString() string
	Index() (int, bool)
}

// list of presets accepting generic preset type
type presetList[T presetLike] struct {
	presets []T
	byName  map[string]T
}

// define a new preset
func (l *presetList[T]) define(name string) T {
	if l.byName == nil {
		// init byName map if it is nil
		l.byName = make(map[string]T)
	}
	if _, ok := l.byName[name]; ok {
		// enough to panic
		panic("duplicate preset " + name)
	}
	// create a new preset based on the preset type
	p := T{preset{len(l.presets), name}}
	l.presets = append(l.presets, p)
	l.byName[name] = p
	return p
}

// get the length of the preset list
func (l *presetList[T]) Length() int { return len(l.presets) }

// get the preset by index
func (l *presetList[T]) FromIndex(i int) (T, error) {
	// range check
	if i < 0 || i >= l.Length() {
		var zero T
		return zero, fmt.Errorf("%d is out of range 0-%d", i, l.Length()-1)
	}
	return l.presets[i], nil
}

// get the preset by its name
func (l *presetList[T]) FromName(name string) (T, error) {
	// lookup the preset by name
	p, ok := l.byName[name]
	if !ok {
		var zero T
		return zero, fmt.Errorf("preset %s not found", name)
	}
	return p, nil
}

// get the presets up to the limit
func (l *presetList[T]) GetPresets(limit int) []T {
	limit = min(max(limit, 0), l.Length())
	return l.presets[:limit]
}
