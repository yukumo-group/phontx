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
var TargetVolume int32 = 100

// TargetSpeed defines the volume of target speed generated
var TargetSpeed int32 = 100

// TargetAudioFileName defines the name for audio file
var TargetAudioFileName string

// TargetAudioDirectory defines the directory that stores audio file
var TargetAudioDirectory string

// TargetPhontxDirectoryPath defines the target path of the exported phontx file
var TargetPhontxDirectoryPath string

// TargetPhontxFileName defines the file name of the target phontx file
var TargetPhontxFileName string

// TargetYAMLPath defines the path of the yaml file to use
var TargetYAMLPath string

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
	err := readPhontxCommand.MarkPersistentFlagRequired("filePath")
	if err != nil {
		panic(err)
	}
	// Set arguments for synth audio
	synthAudioCommand.PersistentFlags().StringVarP(
		&PhontxFilePath,
		"filePath",
		"f",
		"",
		"path to phontx file",
	)
	err = synthAudioCommand.MarkPersistentFlagRequired("filePath")
	if err != nil {
		panic(err)
	}
	synthAudioCommand.PersistentFlags().StringVarP(
		&TargetText,
		"targetText",
		"t",
		"",
		"text to generate",
	)
	err = synthAudioCommand.MarkPersistentFlagRequired("targetText")
	if err != nil {
		panic(err)
	}
	synthAudioCommand.PersistentFlags().StringVarP(
		&TargetAudioFileName,
		"targetAudioFileName",
		"a",
		"",
		"the target name of the audio file (no .wav suffix needed)",
	)
	err = synthAudioCommand.MarkPersistentFlagRequired("targetAudioFileName")
	if err != nil {
		panic(err)
	}
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
	// Set argument for converting yaml file
	toPhontxCommand.PersistentFlags().StringVarP(
		&TargetPhontxDirectoryPath,
		"targetPhontxDirectoryPath",
		"d",
		"./",
		"the path for directory to store phontx file",
	)
	err = toPhontxCommand.MarkPersistentFlagRequired(
		"targetPhontxDirectoryPath",
	)
	if err != nil {
		panic(err)
	}
	toPhontxCommand.PersistentFlags().StringVarP(
		&TargetYAMLPath,
		"targetYAMLPath",
		"y",
		"",
		"the path for target yaml file for inputting",
	)
	err = toPhontxCommand.MarkPersistentFlagRequired(
		"targetYAMLPath",
	)
	if err != nil {
		panic(err)
	}
	toPhontxCommand.PersistentFlags().StringVarP(
		&TargetPhontxFileName,
		"targetPhontxFileName",
		"f",
		"",
		"the name of the target phontx file, no suffix is needed",
	)
	err = toPhontxCommand.MarkPersistentFlagRequired(
		"targetPhontxFileName",
	)
	if err != nil {
		panic(err)
	}
	// Add subcommands for a command
	readPhontxCommand.AddCommand(
		showEngineInfo,
	)
	rootCommand.AddCommand(
		createPhontxCommand,
		readPhontxCommand,
		synthAudioCommand,
		toPhontxCommand,
	)
}

// Execute executes the command
func Execute() error {
	return rootCommand.Execute()
}
