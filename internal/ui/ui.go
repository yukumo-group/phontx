package ui

import (
	"fmt"
	"strings"
)

// GetTrueFalsePrompt gets the prompt for true/false
func GetTrueFalsePrompt(
	defaultChoice bool,
) string {
	yesPrompt := ""
	noPrompt := ""
	if defaultChoice {
		yesPrompt = "Y"
	} else {
		yesPrompt = "y"
	}
	if !defaultChoice {
		noPrompt = "N"
	} else {
		noPrompt = "n"
	}
	return fmt.Sprintf(
		"(%s/%s)",
		yesPrompt,
		noPrompt,
	)
}

// GetAnswerForYesNoQuestion gets the answer for yes or no question
func GetAnswerForYesNoQuestion(
	prompt string,
	defaultValue bool,
) (bool, error) {
	trueFalsePrompt := GetTrueFalsePrompt(
		defaultValue,
	)
	fmt.Printf(
		"%s %s:",
		prompt,
		trueFalsePrompt,
	)
	var ans string
	_, err := fmt.Scan(&ans)
	if err != nil {
		return false, err
	}
	lowerAns := strings.ToLower(ans)
	switch lowerAns {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		return false, fmt.Errorf(
			"%s is not a valid answer",
			lowerAns,
		)
	}
}
