package cerror

/*
#cgo CFLAGS: -I${SRCDIR}/../../third-party
#include<Aqsh.h>
*/
import "C"

// GetError gets error
func GetError() error {
	return nil
}
