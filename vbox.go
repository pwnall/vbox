package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/vbox.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
  "unsafe"
)

// These singletons get initialized by Init().
var client *C.IVirtualBoxClient = nil
var cbox *C.IVirtualBox = nil
var glueInitialized = false

var AppVersion uint = 0
var ApiVersion uint = 0

// Init initializes the VirtualBox global data structures.
//
// Due to VirtualBox oddness, Init should ideally be called in the
// application's main thread. The odds of this happening are maximized by
// calling Init() from the application's main goroutine.
//
// It returns any error encountered.
func Init() error {
  // For convenience, Init() is idempotent.

  if glueInitialized == false {
    result := C.GoVboxCGlueInit()
    if C.GoVboxFAILED(result) != 0 {
      cmessage := C.GoString(&C.g_szVBoxErrMsg[0])
      return errors.New(fmt.Sprintf("VBoxCGlueInit failed: %s", cmessage))
    }

    glueInitialized = true
    AppVersion = uint(C.GoVboxGetAppVersion())
    ApiVersion = uint(C.GoVboxGetApiVersion())
  }

  if client == nil {
    result := C.GoVboxClientInitialize(&client)
    if C.GoVboxFAILED(result) != 0 || client == nil {
      client = nil
      return errors.New(fmt.Sprintf("pfnClientInitialize failed: %x", result))
    }
  }

  if cbox == nil {
    result := C.GoVboxGetVirtualBox(client, &cbox)
    if C.GoVboxFAILED(result) != 0 || cbox == nil {
      cbox = nil
      return errors.New(fmt.Sprintf("Failed to get IVirtualBox: %x", result))
    }
  }

  return nil
}

// Deinit cleans up the VirtualBox global state.
// After this method is called, all VirtualBox-related objects are invalid.
// It returns any error encountered.
func Deinit() error {
  if cbox != nil {
    result := C.GoVboxIVirtualBoxRelease(cbox)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(
          fmt.Sprintf("Failed to release IVirtualBox: %x", result))
    }
    cbox = nil
  }

  if client != nil {
    C.GoVboxClientUninitialize()
    client = nil
  }

  return nil
}

// GetRevision returns VirtualBox's SVN revision as a number.
func GetRevision() (int, error) {
  if err := Init(); err != nil {
    return 0, err
  }

  var revision C.ULONG
  result := C.GoVboxGetRevision(cbox, &revision)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get IVirtualBox revision: %x", result))
  }

  return int(revision), nil
}

// ComposeMachineFilename returns a default VM config file path.
// If baseFolder is empty, VirtualBox's default machine folder will be used.
// It returns a string and any error encountered.
func ComposeMachineFilename(
    name string, flags string, baseFolder string) (string, error) {
  if err := Init(); err != nil {
    return "", err
  }

  var cpath *C.char
  cname := C.CString(name)
  cflags := C.CString(flags)
  cbaseFolder := C.CString(baseFolder)
  result := C.GoVboxComposeMachineFilename(cbox, cname, cflags, cbaseFolder,
      &cpath)
  C.free(unsafe.Pointer(cname))
  C.free(unsafe.Pointer(cflags))
  C.free(unsafe.Pointer(cbaseFolder))

  if C.GoVboxFAILED(result) != 0 || cpath == nil {
    return "", errors.New(
        fmt.Sprintf("IVirtualBox failed to compose machine name: %x", result))
  }

  path := C.GoString(cpath)
  C.GoVboxUtf8Free(cpath)
  return path, nil
}
