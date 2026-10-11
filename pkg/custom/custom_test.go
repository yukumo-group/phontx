package custom

import (
	"testing"
)

// TestAq10ConfigSetting tests the setting for aq10config
func TestAq10ConfigSetting(t *testing.T) {
	t.Parallel()
	expectedCfg1 := &Aq10Config{
		Base: 0,
		FSC:  150,
		LMD:  150,
	}
	expectedCfg2 := &Aq10Config{
		Base: 0,
		FSC:  50,
		LMD:  200,
	}
	cfg1 := NewAq10Config().WithBase(
		F1E,
	).WithFSC(
		150,
	).WithLMD(
		150,
	)
	if !cfg1.Equals(expectedCfg1) {
		t.Error(
			"there is a problem in setting config",
		)
	}
	cfg2 := NewAq10Config().WithBase(
		F1E,
	).WithFSC(
		-1,
	).WithLMD(
		114514,
	)
	if !cfg2.Equals(expectedCfg2) {
		t.Error(
			"there is a problem in setting config",
		)
	}
}
