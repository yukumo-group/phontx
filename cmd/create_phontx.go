package main

import (
	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
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
		tone.NewData()
		ui.GetAnswerForYesNoQuestion(
			"Allows Aquestalk1 engine?",
			false,
		)
	},
}
