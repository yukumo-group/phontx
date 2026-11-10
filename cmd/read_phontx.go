package main

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// readPhontxCommand reads the phontx
var readPhontxCommand = &cobra.Command{
	Use:   "readPhontx",
	Short: "read a phontx file and displays the content",
	Long:  "read a phontx file and displays the content",
	Run: func(cmd *cobra.Command, args []string) {
		fetchedTone, err := tone.NewDataFromFile(
			PhontxFilePath,
		)
		if err != nil {
			panic(err)
		}
		convertedData, err := ui.ToneToString(
			fetchedTone,
		)
		if err != nil {
			panic(err)
		}
		fmt.Println(convertedData)
	},
}
