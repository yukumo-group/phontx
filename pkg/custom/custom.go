package custom

// CustomAq10Config defines the aq10 config by custom
type CustomAq10Config struct {
	Base   int
	Volume int
	Pitch  int
	Accent int
	Lmd    int
	Fsc    int
}

// NewCustomAq10Config creates new aquestalk10 config
func NewCustomAq10Config() *CustomAq10Config {
	return &CustomAq10Config{
		Volume: 100,
		Lmd:    100,
		Fsc:    100,
	}
}
