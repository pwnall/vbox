package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/display.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
)

// The display of a running VM.
type Display struct {
  cdisplay *C.IDisplay
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (display *Display) Release() error {
  if display.cdisplay != nil {
    result := C.GoVboxIDisplayRelease(display.cdisplay)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IDisplay: %x", result))
    }
    display.cdisplay = nil
  }
  return nil
}
