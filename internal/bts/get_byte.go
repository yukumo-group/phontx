package bts

import (
	"context"
)

// GetByte defines the function for getting one byte from one position
type GetByte func(
	context.Context,
	uint32,
) (byte, error)

// ToString gets string from null terminate string stored in memory
func ToString(
	ctx context.Context,
	startPos uint32,
	getByteFunc GetByte,
) (string, error) {
	result := ""
	for i := startPos; i < 256; i++ {
		fetchedData, err := getByteFunc(ctx, i)
		if err != nil {
			return result, err
		}
		stringData := string(fetchedData)
		if stringData == "\x00" {
			break
		}
		result += stringData
	}
	return result, nil
}
