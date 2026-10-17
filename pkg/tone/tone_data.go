//go:generate msgp
package tone

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
)

// Data defines data to store in .phontx data.
// It defines tone for multiple tones.
type Data struct {
	// Aquestalk1
	Aq1Preset *aqdata.Aq1Preset
	// Aquestalk2
	Aq2Preset     *aqdata.Aq2Preset
	PhontFileData []byte
	// Aquestalk10
	Aq10Preset *aqdata.Aq10Preset
	// Defines the tone by custom for aquestalk10
	Aq10CustomTone *custom.Aq10CustomTone
}

// NewDataFromFile reads new data from file
func NewDataFromFile(
	filePath string,
) (*Data, error) {
	return nil, nil
}

// NewData creates new tone
func NewData() *Data {
	return &Data{}
}

// SetAq1Preset sets the preset for aquestalk1
func (tone *Data) SetAq1Preset(
	preset aqdata.Aq1Preset,
) *Data {
	tone.Aq1Preset = &preset
	return tone
}

// SetAq2Preset sets the preset for aquestalk2
func (tone *Data) SetAq2Preset(
	preset aqdata.Aq2Preset,
) *Data {
	tone.Aq2Preset = &preset
	return tone
}

// SetPhontFileData sets the byte data of phont file for aquestalk2
func (tone *Data) SetPhontFileData(
	data []byte,
) *Data {
	tone.PhontFileData = data
	return tone
}

// SetAq10Preset sets the preset of phont file for aquestalk10
func (tone *Data) SetAq10Preset(
	preset aqdata.Aq10Preset,
) *Data {
	tone.Aq10Preset = &preset
	return tone
}

// SetAq10CustomTone sets the custom tone for aquestalk10
func (tone *Data) SetAq10CustomTone(
	customTone *custom.Aq10CustomTone,
) *Data {
	tone.Aq10CustomTone = customTone
	return tone
}

// SetPhontDataFromFile reads phont data and sets the PhontsFileData
func (tone *Data) SetPhontDataFromFile(
	phontFilePath string,
) error {
	_, err := os.Stat(phontFilePath)
	if err != nil {
		return err
	}
	fileName := filepath.Base(
		phontFilePath,
	)
	extension := filepath.Ext(
		fileName,
	)
	if extension != ".phont" {
		return fmt.Errorf(
			"the file extension should be %s, got %s",
			".phont",
			extension,
		)
	}
	data, err := os.ReadFile(phontFilePath)
	if err != nil {
		return err
	}
	tone.SetPhontFileData(data)
	return nil
}
