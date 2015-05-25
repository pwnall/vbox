package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/progress.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
  //"unsafe"
)

// Tracks the progress of a long-running operation.
type Progress struct {
  cprogress *C.IProgress
}

// WaitForCompletion waits for all the operations tracked by this to complete.
// The timeout argument is in milliseconds. -1 is used to wait indefinitely.
// It returns any error encountered.
func (progress* Progress) WaitForCompletion(timeout int) error {
  result := C.GoVboxProgressWaitForCompletion(progress.cprogress,
      C.int(timeout))
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(fmt.Sprintf("Failed to wait on IProgress: %x", result))
  }
  return nil
}

// GetState returns the completion percentage of the tracked operation.
// It returns a number and any error encountered.
func (progress* Progress) GetPercent() (int, error) {
  var cpercent C.PRUint32

  result := C.GoVboxGetProgressPercent(progress.cprogress, &cpercent)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IProgress percent: %x", result))
  }
  return int(cpercent), nil
}

// GetResultCode returns the result code of the tracked operation.
// It returns a number and any error encountered.
func (progress* Progress) GetResultCode() (int, error) {
  var code C.PRUint32

  result := C.GoVboxGetProgressResultCode(progress.cprogress, &code)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IProgress result code: %x", result))
  }
  return int(code), nil
}


// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (progress* Progress) Release() error {
  if progress.cprogress != nil {
    result := C.GoVboxIProgressRelease(progress.cprogress)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(fmt.Sprintf("Failed to release IProgress: %x", result))
    }
    progress.cprogress = nil
  }
  return nil
}
