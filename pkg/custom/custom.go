//go:generate msgp
package custom

// Aq10CustomTone defines the aq10 tone by custom.
// Pitch and Accent depends on base.
type Aq10CustomTone struct {
	// 0-2
	Base int
	// 0-200, default: 100
	LMD int
	// 50-200, default: 100
	FSC int
	// 20-200
	Pitch *int
	// 0-200
	Accent *int
}

// NewAq10CustomTone creates new aquestalk10 config
func NewAq10CustomTone() *Aq10CustomTone {
	return &Aq10CustomTone{
		LMD: 100,
		FSC: 100,
	}
}

// WithBase sets base for config
func (cfg *Aq10CustomTone) WithBase(
	data Aq10Base,
) *Aq10CustomTone {
	cfg.Base = int(data)
	return cfg
}

// WithPitch sets pitch for config
func (cfg *Aq10CustomTone) WithPitch(
	data int,
) *Aq10CustomTone {
	realData := min(max(data, 20), 200)
	cfg.Pitch = &realData
	return cfg
}

// WithAccent sets pitch for config
func (cfg *Aq10CustomTone) WithAccent(
	data int,
) *Aq10CustomTone {
	realData := min(max(data, 0), 200)
	cfg.Accent = &realData
	return cfg
}

// WithLMD sets lmd for config
func (cfg *Aq10CustomTone) WithLMD(
	data int,
) *Aq10CustomTone {
	realData := min(max(data, 0), 200)
	cfg.LMD = realData
	return cfg
}

// WithFSC sets fsc for config
func (cfg *Aq10CustomTone) WithFSC(
	data int,
) *Aq10CustomTone {
	realData := min(max(data, 50), 200)
	cfg.FSC = realData
	return cfg
}

// Equals checks if two configs are the same
func (cfg *Aq10CustomTone) Equals(
	another *Aq10CustomTone,
) bool {
	return cfg.Base == another.Base &&
		cfg.Accent == another.Accent &&
		cfg.Pitch == another.Pitch &&
		cfg.LMD == another.LMD &&
		cfg.FSC == another.FSC
}
