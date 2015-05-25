package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/os_type.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
  "reflect"
  "unsafe"
)

// The description of a supported guest OS type
type GuestOsType struct {
  ctype *C.IGuestOSType
}

// GetId returns the string used to identify this OS type in other API calls.
// It returns a string and any error encountered.
func (osType *GuestOsType) GetId() (string, error) {
  var cid *C.char
  result := C.GoVboxGetGuestOSTypeId(osType.ctype, &cid)
  if C.GoVboxFAILED(result) != 0 || cid == nil {
    return "", errors.New(
        fmt.Sprintf("Failed to get IGuestOSType id: %x", result))
  }

  id := C.GoString(cid)
  C.free(unsafe.Pointer(cid))
  return id, nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (osType *GuestOsType) Release() error {
  if osType.ctype != nil {
    result := C.GoVboxIGuestOSTypeRelease(osType.ctype)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IGuestOSType: %x", result))
    }
    osType.ctype = nil
  }
  return nil
}


// GetGuestOsTypes returns the guest OS types supported by VirtualBox.
// It returns a slice of GuestOsType instances and any error encountered.
func GetGuestOsTypes() ([]GuestOsType, error) {
  if err := Init(); err != nil {
    return nil, err
  }

  var ctypesPtr **C.IGuestOSType
  var typeCount C.ULONG

  result := C.GoVboxGetGuestOSTypes(cbox, &ctypesPtr, &typeCount)
  if C.GoVboxFAILED(result) != 0 || ctypesPtr == nil {
    return nil, errors.New(
        fmt.Sprintf("Failed to get IGuestOSType array: %x", result))
  }

  sliceHeader := reflect.SliceHeader{
    Data: uintptr(unsafe.Pointer(ctypesPtr)),
    Len:  int(typeCount),
    Cap:  int(typeCount),
  }
  ctypesSlice := *(*[]*C.IGuestOSType)(unsafe.Pointer(&sliceHeader))

  var types = make([]GuestOsType, typeCount)
  for i := range ctypesSlice {
    types[i] = GuestOsType{ctypesSlice[i]}
  }

  C.free(unsafe.Pointer(ctypesPtr))
  return types, nil
}

