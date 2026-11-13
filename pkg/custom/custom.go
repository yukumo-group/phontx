package custom

import (
	"bytes"
	"encoding/binary"
)

// Aq10CustomTone defines the aq10 tone by custom.
// Pitch and Accent depends on base.
//
//go:generate msgp
type Aq10CustomTone struct {
	// 0-2
	Base Aq10Base
	// 0-200, default: 100
	LMD uint32
	// 50-200, default: 100
	FSC uint32
	// 20-200
	Pitch *uint32
	// 0-200
	Accent *uint32
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
	cfg.Base = data
	return cfg
}

// WithPitch sets pitch for config
func (cfg *Aq10CustomTone) WithPitch(
	data uint32,
) *Aq10CustomTone {
	realData := min(max(data, 20), 200)
	cfg.Pitch = &realData
	return cfg
}

// WithAccent sets pitch for config
func (cfg *Aq10CustomTone) WithAccent(
	data uint32,
) *Aq10CustomTone {
	realData := min(max(data, 0), 200)
	cfg.Accent = &realData
	return cfg
}

// WithLMD sets lmd for config
func (cfg *Aq10CustomTone) WithLMD(
	data uint32,
) *Aq10CustomTone {
	realData := min(max(data, 0), 200)
	cfg.LMD = realData
	return cfg
}

// WithFSC sets fsc for config
func (cfg *Aq10CustomTone) WithFSC(
	data uint32,
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

// PackData packs the data to byte buffer.
// Return: Base, Data, error
func (cfg *Aq10CustomTone) PackData() ([]byte, []byte, error) {
	baseBuffer := new(bytes.Buffer)
	// GetBaseBuffer
	err := binary.Write(
		baseBuffer,
		binary.LittleEndian,
		cfg.Base.ToInt(),
	)
	if err != nil {
		return nil, nil, err
	}
	resultBuffer := new(bytes.Buffer)
	// Pitch
	pitch := GetPitch(
		cfg.Base,
		cfg.Pitch,
	)
	err = binary.Write(
		resultBuffer,
		binary.LittleEndian,
		pitch,
	)
	if err != nil {
		return nil, nil, err
	}
	// Accent
	accent := GetAccent(
		cfg.Base,
		cfg.Accent,
	)
	err = binary.Write(
		resultBuffer,
		binary.LittleEndian,
		accent,
	)
	if err != nil {
		return nil, nil, err
	}
	// LMD
	err = binary.Write(
		resultBuffer,
		binary.LittleEndian,
		cfg.LMD,
	)
	if err != nil {
		return nil, nil, err
	}
	// FSC
	err = binary.Write(
		resultBuffer,
		binary.LittleEndian,
		cfg.FSC,
	)
	if err != nil {
		return nil, nil, err
	}
	return baseBuffer.Bytes(), resultBuffer.Bytes(), nil
}
