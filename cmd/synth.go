package main

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/yukumo-group/phontx/internal/ui"
	"github.com/yukumo-group/phontx/pkg/synth"
	"github.com/yukumo-group/phontx/pkg/tone"
)

// synthAudioCommand synthesis audio command
var synthAudioCommand *cobra.Command = &cobra.Command{
	Use:   "synthAudio",
	Short: "synthesis audio from phontx",
	Long:  "synthesis audio from phontx file",
	Run: func(cmd *cobra.Command, args []string) {
		ctx := context.Background()
		// Get data from file
		loadedTone, err := tone.NewDataFromFile(
			PhontxFilePath,
		)
		if err != nil {
			panic(err)
		}
		// get version
		selectedVersion, err := ui.SelectVersion(
			"select the version for generation: ",
		)
		if err != nil {
			panic(err)
		}
		// generation
		var filePath string
		if TargetAudioDirectory == "" {
			filePath = fmt.Sprintf(
				"./%s.wav",
				TargetAudioFileName,
			)
		} else {
			filePath = fmt.Sprintf(
				"%s/%s.wav",
				TargetAudioDirectory,
				TargetAudioFileName,
			)
		}
		synthParam := synth.NewSynthParam(
			TargetText,
			uint32(TargetSpeed),
			uint32(TargetVolume),
			filePath,
		)
		generator, err := synth.NewGenerator(
			ctx,
			runtime,
		)
		if err != nil {
			panic(err)
		}
		err = generator.Generate(
			ctx,
			selectedVersion,
			synthParam,
			loadedTone,
		)
		if err != nil {
			panic(err)
		}
		// show generation info
		fmt.Printf(
			"audio saved at %s\n",
			filePath,
		)
		// close generator
		err = generator.CloseGenerator(
			ctx,
		)
		if err != nil {
			panic(err)
		}
	},
}
