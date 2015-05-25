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
