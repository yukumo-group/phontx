package bts

import (
	"context"
	"errors"
	"testing"
)

// TestGetStringFromBytes tests getting string from bytes
func TestGetStringFromBytes(t *testing.T) {
	t.Parallel()
	ctx := t.Context()
	const originalText string = "hello\x00"
	const notNullText string = "hello"
	convertedText := []byte(originalText)
	convertFunc := func(
		ctx context.Context,
		i uint32,
	) (byte, error) {
		if i >= uint32(len(convertedText)) {
			return byte(0), errors.New(
				"out of range",
			)
		}
		return convertedText[i], nil
	}
	fetchedString, err := ToString(
		ctx,
		0,
		convertFunc,
	)
	if err != nil {
		t.Fatal(err)
	}
	if fetchedString != notNullText {
		t.Fatalf(
			"expected string to be %s, got %s",
			notNullText,
			fetchedString,
		)
	}
}
