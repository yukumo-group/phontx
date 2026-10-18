package bts

import (
	"bytes"
)

// IsBytesSame compares multiple bytes to see if these bytes are the same
func IsBytesSame(
	bytesData ...[]byte,
) bool {
	for i := 1; i < len(bytesData); i++ {
		isEqual := bytes.Compare(
			bytesData[i],
			bytesData[i-1],
		)
		if isEqual != 0 {
			return false
		}
	}
	return true
}
