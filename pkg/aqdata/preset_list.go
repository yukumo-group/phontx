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

type presetList[T any] struct {
	presets []T
	byName  map[string]T
	build   func(index int, name string) T
}

func (l *presetList[T]) define(name string) T {
	if l.byName == nil {
		l.byName = make(map[string]T)
	}
	if _, ok := l.byName[name]; ok {
		panic("duplicate preset " + name)
	}
	p := l.build(len(l.presets), name)
	l.presets = append(l.presets, p)
	l.byName[name] = p
	return p
}

func (l *presetList[T]) fromIndex(i int) (T, error) {
	if i < 0 || i >= len(l.presets) {
		var zero T
		return zero, fmt.Errorf("%d is out of range 0-%d", i, len(l.presets)-1)
	}
	return l.presets[i], nil
}

func (l *presetList[T]) fromName(name string) (T, error) {
	p, ok := l.byName[name]
	if !ok {
		var zero T
		return zero, fmt.Errorf("preset %s not found", name)
	}
	return p, nil
}

func (l *presetList[T]) prefix(n int) []T {
	n = min(max(n, 0), len(l.presets))
	out := make([]T, n)
	copy(out, l.presets[:n])
	return out
}
