package tone

import (
	"testing"

	"github.com/yukumo-group/phontx/pkg/aqdata"
)

// TestYAMLConvertion tests the convertion of yaml data to raw tone
func TestYAMLConvertion(t *testing.T) {
	t.Parallel()
	NewData().SetAq1Preset(
		aqdata.Aq1F1,
	)
	RawYAMLFileToPhontx(
		"testdata/testdata.yaml",
	)
}
