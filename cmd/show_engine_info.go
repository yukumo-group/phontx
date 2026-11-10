package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// showEngineInfo shows engine info
var showEngineInfo = &cobra.Command{
	Use:   "showEngineInfo",
	Short: "shows which engine it supports",
	Long:  "shows which engine it supports",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(
			"Engines supported by %s: \n",
			PhontxFilePath,
		)
		toneData, err := tone.NewDataFromFile(
			PhontxFilePath,
		)
		if err != nil {
			panic(err)
		}
		ui.ShowEngineInfo(
			toneData,
		)
	},
}
