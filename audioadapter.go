package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/audioadapter.c"
*/
import "C"
import (
	"errors"
	"fmt"
) // cgo's virtual package

// The description of a audio adapter
type AudioAdapter struct {
	caudioadapter *C.IAudioAdapter
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (a *AudioAdapter) Release() error {
	if a.caudioadapter != nil {
		result := C.GoVboxIAudioAdapterRelease(a.caudioadapter)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(fmt.Sprintf("Failed to release IAudioAdapter: %x", result))
		}
		a.caudioadapter = nil
	}
	return nil
}
