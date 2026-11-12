package ui

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/1Vewton/structdisplay/pkg/display"
	"github.com/1Vewton/structdisplay/pkg/ignore"
	"github.com/tinylib/msgp/msgp"
	"github.com/yukumo-group/phontx/pkg/aqdata"
	"github.com/yukumo-group/phontx/pkg/custom"
	"github.com/yukumo-group/phontx/pkg/tone"
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

// NewDataNeededDefinitionWithCF creates new data needed definition
func NewDataNeededDefinitionWithCF[T any](
	paramName string,
	definition string,
	ceil T,
	floor T,
	valuePointer *T,
) DataNeededDefinition[T] {
	return DataNeededDefinition[T]{
		ParamName:    paramName,
		Definition:   definition,
		Ceil:         &ceil,
		Floor:        &floor,
		ValuePointer: valuePointer,
	}
}

// GetData gets all the params for a
func (param DataNeededDefinition[T]) GetData() error {
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

// GetMultipleParams gets multiple params
func GetMultipleParams[T any](
	params ...DataNeededDefinition[T],
) error {
	for _, param := range params {
		err := param.GetData()
		if err != nil {
			return err
		}
	}
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
	baseIdx, err := FetchVal[int32](
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
	// Set lmd and fsc
	customLmd := 0
	customFsc := 50
	err = GetMultipleParams(
		NewDataNeededDefinitionWithCF(
			"LMD",
			"pitch 1",
			0,
			200,
			&customLmd,
		),
		NewDataNeededDefinitionWithCF(
			"FSC",
			"pitch 2",
			50,
			200,
			&customFsc,
		),
	)
	newCustomTone := custom.NewAq10CustomTone().WithBase(
		selectedBase,
	).WithLMD(
		int32(customLmd),
	).WithFSC(
		int32(customFsc),
	)
	// Adjust accent or pitch
	adjustAccentOrPitch, err := GetAnswerForYesNoQuestion(
		"Do you want to adjust accent and pitch value which depends on the base setted?",
		false,
	)
	if err != nil {
		return newCustomTone, err
	}
	if adjustAccentOrPitch {
		customPitch := 20
		customAccent := 0
		err = GetMultipleParams(
			NewDataNeededDefinitionWithCF(
				"Pitch",
				"pitch of the tone",
				20,
				200,
				&customPitch,
			),
			NewDataNeededDefinitionWithCF(
				"Accent",
				"the strength of the tone",
				0,
				200,
				&customAccent,
			),
		)
		if err != nil {
			return newCustomTone, err
		}
		newCustomTone.WithPitch(
			int32(customPitch),
		).WithAccent(
			int32(customAccent),
		)
	}
	return newCustomTone, nil
}

// ToneToString converts tone to string
func ToneToString(
	target *tone.Data,
) (string, error) {
	sampleBytes := []byte{}
	sampleType := reflect.TypeOf(sampleBytes)
	ignoredList := ignore.NewTermsIgnoredSlice()
	ignoredList.AddTypeByType(sampleType)
	return display.NewConverter().WithRecursionLimit(
		99,
	).WithIgnoredTerms(
		ignoredList,
	).ConvertAnyData(
		"ToneData",
		"",
		target,
		0,
	)
}

// ShowEngineSupportOrNot shows whether a version of engine is supported or not
func ShowEngineSupportOrNot(
	version int,
	isSupported bool,
) {
	var supportedText string
	if isSupported {
		supportedText = "Supported"
	} else {
		supportedText = "Not Supported"
	}
	fmt.Printf(
		"Aquestalk%d: %s\n",
		version,
		supportedText,
	)
}

// ShowEngineInfo shows the engine info for a phont file
func ShowEngineInfo(
	target *tone.Data,
) {
	allowedEngine := target.GetAllowedEngines()
	ShowEngineSupportOrNot(
		1,
		allowedEngine.Aq1,
	)
	ShowEngineSupportOrNot(
		2,
		allowedEngine.Aq2,
	)
	ShowEngineSupportOrNot(
		10,
		allowedEngine.Aq10,
	)
}
