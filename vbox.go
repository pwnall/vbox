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
  "reflect"
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
  var revision C.ULONG

  result := C.GoVboxGetRevision(cbox, &revision)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New("Failed to get IVirtualBox revision")
  }

  return int(revision), nil
}

// The description of a VirtualBox machine
type Machine struct {
  cmachine *C.IMachine
}

func GetMachines() ([]Machine, error) {
  var cmachinesPtr **C.IMachine
  var machinesCount C.ULONG

  result := C.GoVboxGetMachines(cbox, &cmachinesPtr, &machinesCount)
  if C.GoVboxFAILED(result) != 0 || cmachinesPtr == nil {
    return nil, errors.New(
        fmt.Sprintf("Failed to get IMachine array: %x", result))
  }

  sliceHeader := reflect.SliceHeader{
    Data: uintptr(unsafe.Pointer(cmachinesPtr)),
    Len:  int(machinesCount),
    Cap:  int(machinesCount),
  }
  cmachinesSlice := *(*[]*C.IMachine)(unsafe.Pointer(&sliceHeader))

  var machines = make([]Machine, machinesCount)
  for i := range cmachinesSlice {
    machines[i] = Machine{cmachinesSlice[i]}
  }

  C.free(unsafe.Pointer(cmachinesPtr))

  return machines, nil
}

// Release frees up the VirtualBox data associated with this machine.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
func (machine* Machine) Release() error {
  if machine.cmachine != nil {
  }
  return nil
}

type Session struct {
  csession *C.ISession
}

func (session *Session) Init() error {
  if err := Init(); err != nil {
    return err
  }

  result := C.GoVboxGetSession(client, &session.csession)
  if C.GoVboxFAILED(result) != 0 || session.csession == nil {
    session.csession = nil
    return errors.New(fmt.Sprintf("Failed to get ISession: %x", result))
  }
  return nil
}
func (session *Session) Release() error {
  if session.csession != nil {
    result := C.GoVboxISessionRelease(session.csession)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New(fmt.Sprintf("Failed to release ISession: %x", result))
    }
    session.csession = nil
  }
  return nil
}

