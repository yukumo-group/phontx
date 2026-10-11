package custom

// Aq10Config defines the aq10 config by custom.
// Pitch and Accent depends on base.
type Aq10Config struct {
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

// NewAq10Config creates new aquestalk10 config
func NewAq10Config() *Aq10Config {
	return &Aq10Config{
		LMD: 100,
		FSC: 100,
	}
}

// WithBase sets base for config
func (cfg *Aq10Config) WithBase(
	data Aq10Base,
) *Aq10Config {
	cfg.Base = int(data)
	return cfg
}

// WithPitch sets pitch for config
func (cfg *Aq10Config) WithPitch(
	data int,
) *Aq10Config {
	realData := min(max(data, 20), 200)
	cfg.Pitch = &realData
	return cfg
}

// WithAccent sets pitch for config
func (cfg *Aq10Config) WithAccent(
	data int,
) *Aq10Config {
	realData := min(max(data, 0), 200)
	cfg.Accent = &realData
	return cfg
}

// WithLMD sets lmd for config
func (cfg *Aq10Config) WithLMD(
	data int,
) *Aq10Config {
	realData := min(max(data, 0), 200)
	cfg.LMD = realData
	return cfg
}

// WithFSC sets fsc for config
func (cfg *Aq10Config) WithFSC(
	data int,
) *Aq10Config {
	realData := min(max(data, 50), 200)
	cfg.FSC = realData
	return cfg
}

// Equals checks if two configs are the same
func (cfg *Aq10Config) Equals(
	another *Aq10Config,
) bool {
	return cfg.Base == another.Base &&
		cfg.Accent == another.Accent &&
		cfg.Pitch == another.Pitch &&
		cfg.LMD == another.LMD &&
		cfg.FSC == another.FSC
}
