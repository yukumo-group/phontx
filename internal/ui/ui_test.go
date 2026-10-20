package ui

import (
	"testing"
)

// TestTrueFalsePrompt tests the generation for prompt for true/false question
func TestTrueFalsePrompt(t *testing.T) {
	t.Parallel()
	const defaultTrueResult string = "(Y/n)"
	const defaultFalseResult string = "(y/N)"
	trueResult := GetTrueFalsePrompt(
		true,
	)
	falseResult := GetTrueFalsePrompt(
		false,
	)
	if trueResult != defaultTrueResult {
		t.Errorf(
			"expected %s, got %s",
			defaultTrueResult,
			trueResult,
		)
	}
	if falseResult != defaultFalseResult {
		t.Errorf(
			"expected %s, got %s",
			defaultFalseResult,
			falseResult,
		)
	}
}
