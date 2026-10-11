package custom

// Aq10Config defines the aq10 config by custom
type Aq10Config struct {
	Base   int
	Volume int
	Pitch  int
	Accent int
	Lmd    int
	Fsc    int
}

// NewAq10Config creates new aquestalk10 config
func NewAq10Config() *Aq10Config {
	return &Aq10Config{
		Volume: 100,
		Lmd:    100,
		Fsc:    100,
	}
}
