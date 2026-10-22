package ui

import (
	"errors"
	"fmt"
	"strings"

	"github.com/tinylib/msgp/msgp"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
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

// FetchVal fetches the value through inputing
func FetchVal[T any](
	prompt string,
) (T, error) {
	fmt.Println(prompt)
	var data T
	_, err := fmt.Scan(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}

// DataNeededDefinition provides definition for getting any data
type DataNeededDefinition[T any] struct {
	ParamName    string
	Definition   string
	Ceil         *T
	Floor        *T
	ValuePointer *T
}

// GetToneParams gets all the params for a
func GetToneParams[T any](
	param DataNeededDefinition[T],
) error {
	var showData T
	prompt := fmt.Sprintf(
		"Input %T value for %s, %s",
		showData,
		param.ParamName,
		param.Definition,
	)
	if param.Ceil != nil && param.Floor != nil {
		prompt += fmt.Sprintf(
			" range: (%v-%v)",
			*param.Ceil,
			*param.Floor,
		)
	}
	data, err := FetchVal[T](prompt)
	if err != nil {
		return err
	}
	if param.ValuePointer == nil {
		return errors.New(
			"you cannot pass a nil pointer into the data",
		)
	}
	*param.ValuePointer = data
	return nil
}

// GetAq10CustomToneByUI gets the aquestalk10 custom tone by user interface
func GetAq10CustomToneByUI() (*custom.Aq10CustomTone, error) {
	// Base
	bases, err := custom.GetAllAq10Bases(3)
	if err != nil {
		return nil, err
	}
	for i, base := range bases {
		fmt.Printf(
			"index: %d \n- base name: %s\n\n",
			i,
			base.ToString(),
		)
	}
	baseIdx, err := FetchVal[int](
		"input the index for the base of the custome tone: ",
	)
	if err != nil {
		return nil, err
	}
	selectedBase, err := custom.ToAq10Base(
		baseIdx,
	)
	if err != nil {
		return nil, err
	}
	// LMD
	customLmd, err := FetchVal[int](
		"input the lmd (pitch 1) of the custom tone",
	)
	if err != nil {
		return nil, err
	}
	newCustomTone := custom.NewAq10CustomTone().WithBase(
		selectedBase,
	).WithLMD(
		customLmd,
	)
	return newCustomTone, nil
}
