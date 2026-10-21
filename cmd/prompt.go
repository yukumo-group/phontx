package main

import (
	"fmt"
)

// GetAllowsEnginePrompt returns the prompt for whether an engine is allowed
func GetAllowsEnginePrompt(
	engineName string,
) string {
	return fmt.Sprintf(
		"Allows %s engine?",
		engineName,
	)
}
