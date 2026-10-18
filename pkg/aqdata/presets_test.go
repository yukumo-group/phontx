package aqdata

import (
	"testing"
)

// TestPresetsCompare tests the comparison of presets
func TestPresetsCompare(t *testing.T) {
	t.Parallel()
	if IsPresetsEqual(&Aq10F1, &Aq10F2) {
		t.Errorf(
			"expected %s and %s not to be the same",
			Aq10F1.ToString(),
			Aq10F2.ToString(),
		)
	}
	if !IsPresetsEqual(&Aq2AqF1c, &Aq2AqF1c) {
		t.Errorf(
			"%s expected to be equals to %s",
			Aq2AqF1c.ToString(),
			Aq2AqF1c.ToString(),
		)
	}
}
