package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// toPhontxCommand converts yaml file to phontx data
var toPhontxCommand *cobra.Command = &cobra.Command{
	Use:   "toPhontx",
	Short: "Converts yaml file to phontx data",
	Long:  "Converts yaml file to phontx data",
	Run: func(cmd *cobra.Command, args []string) {
		phontxFile, err := tone.RawYAMLFileToPhontx(
			TargetYAMLPath,
		)
		if err != nil {
			panic(err)
		}
		convertedContent, err := ui.ToneToString(
			phontxFile,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(convertedContent)
		filePath, err := phontxFile.SaveFile(
			TargetPhontxDirectoryPath,
			TargetPhontxFileName,
		)
		if err != nil {
			panic(err)
		}
		fmt.Printf(
			"file saved at %s: \n",
			filePath,
		)
	},
}
