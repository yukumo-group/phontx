package main

import (
	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/aqdata"
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
			aq1Preset, err := ui.GetPresetByUI[aqdata.Aq1Preset, *aqdata.Aq1Preset](
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
		byPreset := func() {

		}
		methods := []func(){
			byPreset,
		}
		if allowsAq2 {
			for len(methods) > 0 {

			}
		}
	},
}
