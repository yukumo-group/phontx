package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// createPhontxCommand creates phontx
var createPhontxCommand = &cobra.Command{
	Use:   "createPhontx",
	Short: "createPhontx command creates phontx, which is a tone data allowed for three versions of aquestalk",
	Long: `
createPhontx command creates phontx, which is a tone data allowed for three versions of aquestalk
	`,
	Run: func(cmd *cobra.Command, args []string) {
		newTone := tone.NewData()
		allowsAq1, err := ui.GetAnswerForYesNoQuestion(
			GetAllowsEnginePrompt(
				aqdata.Aq1Presets.GetListName(),
			),
			false,
		)
		if err != nil {
			panic(err)
		}
		// Aquestalk1
		if allowsAq1 {
			aq1Preset, err := ui.GetPresetByUI(
				aqdata.Aq1Presets,
			)
			if err != nil {
				panic(err)
			}
			newTone = newTone.SetAq1Preset(
				aq1Preset,
			)
		}
		// Aquestalk2
		allowsAq2, err := ui.GetAnswerForYesNoQuestion(
			GetAllowsEnginePrompt(
				aqdata.Aq2Presets.GetListName(),
			),
			false,
		)
		if err != nil {
			panic(err)
		}
		if allowsAq2 {
			usePreset, err := ui.GetAnswerForYesNoQuestion(
				"do you want to use preset as custom tone instead of phont file?",
				false,
			)
			if err != nil {
				panic(err)
			}
			if usePreset {
				// By presets
				aq2Preset, err := ui.GetPresetByUI(
					aqdata.Aq2Presets,
				)
				if err != nil {
					panic(err)
				}
				newTone = newTone.SetAq2Preset(
					aq2Preset,
				)
			} else {
				// By phont file path
				usedPhontFilePath, err := ui.FetchVal[string](
					"input the path of the phont file you want to use: ",
				)
				if err != nil {
					panic(err)
				}
				_, err = os.Stat(
					usedPhontFilePath,
				)
				if err != nil {
					panic(err)
				}
				err = newTone.SetPhontDataFromFile(
					usedPhontFilePath,
				)
				if err != nil {
					panic(err)
				}
			}
		}
		// Aquestalk10
		allowsAq10, err := ui.GetAnswerForYesNoQuestion(
			GetAllowsEnginePrompt(
				aqdata.Aq10Presets.GetListName(),
			),
			false,
		)
		if err != nil {
			panic(err)
		}
		if allowsAq10 {
			usePreset, err := ui.GetAnswerForYesNoQuestion(
				"do you want to use presets instead of custom tone",
				false,
			)
			if err != nil {
				panic(err)
			}
			if usePreset {
				// By presets
				aq10Preset, err := ui.GetPresetByUI(
					aqdata.Aq10Presets,
				)
				if err != nil {
					panic(err)
				}
				newTone = newTone.SetAq10Preset(
					aq10Preset,
				)
			} else {
				// Set custom tone
				customTone := custom.NewAq10CustomTone()
				newTone = newTone.SetAq10CustomTone(
					customTone,
				)
			}
		}
		// Saves file
		fmt.Println(
			"new tone configuration done, input the directory you want to save the phontx file in:",
		)
		var directorySavePhont string
		_, err = fmt.Scan(&directorySavePhont)
		if err != nil {
			panic(err)
		}
		fmt.Println(
			"input the name of the tone: ",
		)
		var toneName string
		_, err = fmt.Scan(&toneName)
		if err != nil {
			panic(err)
		}
		fileName, err := newTone.SaveFile(
			directorySavePhont,
			toneName,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"file saved at %s\n",
			fileName,
		)
	},
}
