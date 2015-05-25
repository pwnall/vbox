package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/medium.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
  "unsafe"
)

// Enumeration of MediumVariant values
type MediumVariant uint

const (
  // Default image options.
  MediumVariant_Standard MediumVariant = C.MediumVariant_Standard
  // Entire image is allocated at creation time.
  MediumVariant_Fixed MediumVariant = C.MediumVariant_Fixed
  // The image's directory is not created.
  MediumVariant_NoCreateDir MediumVariant = C.MediumVariant_NoCreateDir
)

// Enumeration of MediumState values
type MediumState uint
const (
  // The medium's backing image was not created or was deleted.
  MediumState_NotCreated = C.MediumState_NotCreated
  // The medium's backing image was not created or was deleted.
  MediumState_Created = C.MediumState_Created
  // The medium's backing image is locked with a shared reader lock.
  MediumState_LockedRead = C.MediumState_LockedRead
  // The medium's backing image is locked with an exclusive writer lock.
  MediumState_LockedWrite = C.MediumState_LockedWrite
  // The medium's backing image cannot / was not accessed.
  MediumState_Inaccessible = C.MediumState_Inaccessible
  // The medium's backing image is being built.
  MediumState_Creating = C.MediumState_Creating
  // The medium's backing image is being deleted.
  MediumState_Deleting = C.MediumState_Deleting
)

// The description of a VirtualBox storage medium
type Medium struct {
  cmedium *C.IMedium
}

// GetLocation returns the path to the image file backing the storage medium.
// It returns a string and any error encountered.
func (medium *Medium) GetLocation() (string, error) {
  var clocation *C.char
  result := C.GoVboxGetMediumLocation(medium.cmedium, &clocation)
  if C.GoVboxFAILED(result) != 0 || clocation == nil {
    return "", errors.New(
        fmt.Sprintf("Failed to get IMedium location: %x", result))
  }

  id := C.GoString(clocation)
  C.free(unsafe.Pointer(clocation))
  return id, nil
}

// GetState returns the last known medium state.
// It returns a MediumState enum instance and any error encountered.
func (medium* Medium) GetState() (MediumState, error) {
  var cstate C.PRUint32

  result := C.GoVboxGetMediumState(medium.cmedium, &cstate)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IMedium state: %x", result))
  }
  return MediumState(cstate), nil
}

// CreateBaseStorage starts building a hard disk image.
// It returns a Progress and any error encountered.
func (medium *Medium) CreateBaseStorage(
    size uint64, variants []MediumVariant) (Progress, error) {
  cvariants := make([]C.PRUint32, len(variants))
  for i, variant := range variants {
    cvariants[i] = C.PRUint32(variant)
  }

  var progress Progress
  result := C.GoVboxMediumCreateBaseStorage(medium.cmedium, C.PRInt64(size),
      C.PRUint32(len(variants)), &cvariants[0], &progress.cprogress)
  if C.GoVboxFAILED(result) != 0 || progress.cprogress == nil {
    return progress, errors.New(
        fmt.Sprintf("Failed to create IMedium storage: %x", result))
  }
  return progress, nil
}


// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (medium* Medium) Release() error {
  if medium.cmedium != nil {
    result := C.GoVboxIMediumRelease(medium.cmedium)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(fmt.Sprintf("Failed to release IMedium: %x", result))
    }
    medium.cmedium = nil
  }
  return nil
}


// CreateHardDisk creates a VirtualBox storage medium for a hard disk image.
// The disk's contents must be created by calling createBaseStorage.
// It returns the created medium and any error encountered.
func CreateHardDisk(formatId string, location string) (Medium, error) {
  var medium Medium
  if err := Init(); err != nil {
    return medium, err
  }

  cformatId := C.CString(formatId)
  clocation := C.CString(location)
  result := C.GoVboxCreateHardDisk(cbox, cformatId, clocation, &medium.cmedium)
  C.free(unsafe.Pointer(cformatId))
  C.free(unsafe.Pointer(clocation))

  if C.GoVboxFAILED(result) != 0 || medium.cmedium == nil {
    return medium, errors.New(
        fmt.Sprintf("Failed to create hard disk IMedium: %x", result))
  }
  return medium, nil
}

