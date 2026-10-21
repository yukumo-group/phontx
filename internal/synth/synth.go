//go:build !test

package synth

/*
#cgo windows LDFLAGS: -L${SRCDIR}/../../third-party/win64 -lAqsh
#cgo linux LDFLAGS: -ldl
#cgo CFLAGS: -I${SRCDIR}/../../third-party
#include<stdlib.h>
#include "Aqsh.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

// InitializeSynth initializes the synthesis program
func InitializeSynth(
	envFilePath string,
) error {
	cEnvFilePath := C.CString(envFilePath)
	initializeResult := C.AqShim_Init(
		cEnvFilePath,
	)
	if initializeResult != 0 {
		return fmt.Errorf(
			"failed to load env file %s",
			envFilePath,
		)
	}
	C.free(unsafe.Pointer(cEnvFilePath))
	return nil
}
