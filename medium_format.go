package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/medium_format.c"
*/
import "C" // cgo's virtual package

import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

// The description of a supported storage medium format
type MediumFormat struct {
	cformat *C.IMediumFormat
}

// GetId returns the string used to identify this format in other API calls.
// It returns a string and any error encountered.
func (format *MediumFormat) GetId() (string, error) {
	var cid *C.char
	result := C.GoVboxGetMediumFormatId(format.cformat, &cid)
	if C.GoVboxFAILED(result) != 0 || cid == nil {
		return "", errors.New(
			fmt.Sprintf("Failed to get IMediumFormat id: %x", result))
	}

	id := C.GoString(cid)
	C.GoVboxUtf8Free(cid)
	return id, nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (format *MediumFormat) Release() error {
	if format.cformat != nil {
		result := C.GoVboxIMediumFormatRelease(format.cformat)
		if C.GoVboxFAILED(result) != 0 {
			return errors.New(
				fmt.Sprintf("Failed to release IMediumFormat: %x", result))
		}
		format.cformat = nil
	}
	return nil
}

// Initialized returns true if there is VirtualBox data associated with this.
func (format *MediumFormat) Initialized() bool {
	return format.cformat != nil
}

// GetMediumFormats returns the guest OS formats supported by VirtualBox.
// It returns a slice of MediumFormat instances and any error encountered.
func (props *SystemProperties) GetMediumFormats() ([]MediumFormat, error) {
	var cformatsPtr **C.IMediumFormat
	var formatCount C.ULONG

	result := C.GoVboxGetMediumFormats(props.cprops, &cformatsPtr, &formatCount)
	if C.GoVboxFAILED(result) != 0 || (cformatsPtr == nil && formatCount > 0) {
		return nil, errors.New(
			fmt.Sprintf("Failed to get IMediumFormat array: %x", result))
	}

	sliceHeader := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cformatsPtr)),
		Len:  int(formatCount),
		Cap:  int(formatCount),
	}
	cformatsSlice := *(*[]*C.IMediumFormat)(unsafe.Pointer(&sliceHeader))

	var formats = make([]MediumFormat, formatCount)
	for i := range cformatsSlice {
		formats[i] = MediumFormat{cformatsSlice[i]}
	}

	C.GoVboxArrayOutFree(unsafe.Pointer(cformatsPtr))
	return formats, nil
}
