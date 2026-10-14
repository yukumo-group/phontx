package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

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

// Execute executes the command
func Execute() error {
	return rootCommand.Execute()
}
