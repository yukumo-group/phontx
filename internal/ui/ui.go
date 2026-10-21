package ui

import (
	"fmt"
	"strings"

	"github.com/tinylib/msgp/msgp"
	"github.com/yukumo-group/phontx/pkg/aqdata"
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
		"%s %s:\n",
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

// GetPresetByUI gets the preset by ui
func GetPresetByUI[T aqdata.PresetLike, PT msgp.RTFor[T]](
	presetList *aqdata.PresetList[T, PT],
) (T, error) {
	// Get data
	presetsList := presetList.Length()
	allPresets := presetList.GetPresets(
		presetsList,
	)
	// Presents all the data
	fmt.Printf(
		"Here are the available presets for engine %s: \n",
		presetList.GetListName(),
	)
	for _, preset := range allPresets {
		fmt.Println(
			preset.ToString(),
		)
	}
	// Input & get the result
	fmt.Println(
		"Input the preset you want to use: ",
	)
	var selectedPreset string
	fmt.Scan(&selectedPreset)
	result, err := presetList.FromName(selectedPreset)
	return result, err
}
