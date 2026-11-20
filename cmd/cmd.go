package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

// PhontxFilePath defines the file path of phontx file
var PhontxFilePath string

// TargetText defines the target text for generation
var TargetText string

// TargetVolume defines the volume of target audio generated
var TargetVolume int32 = -1

// TargetSpeed defines the volume of target speed generated
var TargetSpeed int32 = -1

// TargetAudioFileName defines the name for audio file
var TargetAudioFileName string

// TargetAudioDirectory defines the directory that stores audio file
var TargetAudioDirectory string

// rootCommand defines the root command
var rootCommand = &cobra.Command{
	Use:   "phontx",
	Short: "The cli tool for generating phontx file tool",
	Long: `
PhontX is a tool for generating phontx file. 
Phontx file is a file defines the tone for aquestalk TTS engine across three versions (1,2,10). 
You can: 
- Create and save phontx file. 
- Listen to generated examples. 
- Generate audios using phontx file
	`,
	Run: func(cmd *cobra.Command, args []string) {
		DisplayWelcome()
		fmt.Println(
			"Welcome to phontx",
		)
	},
}

// init adds subcommands
func init() {
	// Set arguments for readPhontx
	readPhontxCommand.PersistentFlags().StringVarP(
		&PhontxFilePath,
		"filePath",
		"f",
		"",
		"path to phontx file",
	)
	readPhontxCommand.MarkFlagRequired("filePath")
	// Set arguments for synth audio
	synthAudioCommand.PersistentFlags().StringVarP(
		&PhontxFilePath,
		"filePath",
		"f",
		"",
		"path to phontx file",
	)
	synthAudioCommand.MarkFlagRequired("filePath")
	synthAudioCommand.PersistentFlags().StringVarP(
		&TargetText,
		"targetText",
		"t",
		"",
		"text to generate",
	)
	synthAudioCommand.MarkFlagRequired("targetText")
	synthAudioCommand.PersistentFlags().StringVarP(
		&TargetAudioFileName,
		"targetAudioFileName",
		"a",
		"",
		"the target name of the audio file (no .wav suffix needed)",
	)
	synthAudioCommand.MarkFlagRequired("targetAudioFileName")
	synthAudioCommand.PersistentFlags().StringVarP(
		&TargetAudioDirectory,
		"targetAudioDirectory",
		"d",
		"",
		"the target directory to store the audio",
	)
	synthAudioCommand.PersistentFlags().Int32VarP(
		&TargetVolume,
		"targetVolume",
		"v",
		-1,
		"text to generate",
	)
	synthAudioCommand.PersistentFlags().Int32VarP(
		&TargetSpeed,
		"targetSpeed",
		"s",
		-1,
		"volume of the generated audio",
	)
	// Add subcommands for a command
	readPhontxCommand.AddCommand(
		showEngineInfo,
	)
	rootCommand.AddCommand(
		createPhontxCommand,
		readPhontxCommand,
		synthAudioCommand,
	)
}

// Execute executes the command
func Execute() error {
	return rootCommand.Execute()
}
