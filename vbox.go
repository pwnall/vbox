package main

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

// This singleton gets initialized by Init().
var client *C.IVirtualBoxClient = nil

var AppVersion uint = 0
var ApiVersion uint = 0

func Init() error {
  // For convenience, Init() is idempotent.
  if client != nil {
    return nil
  }

  result := C.GoVboxCGlueInit()
  if C.GoVboxFAILED(result) != 0 {
    cmessage := C.GoString(&C.g_szVBoxErrMsg[0])
    return errors.New(fmt.Sprintf("VBoxCGlueInit failed: %s", cmessage))
  }

  AppVersion = uint(C.GoVboxGetAppVersion())
  ApiVersion = uint(C.GoVboxGetApiVersion())
  fmt.Printf("VBox App: %d API: %d\n", AppVersion, ApiVersion)

  result = C.GoVboxClientInitialize(&client)
  fmt.Printf("HRESULT: %x %v\n", result, client)
  if C.GoVboxFAILED(result) != 0 {
    fmt.Printf("FAILED: %x\n", C.GoVboxFAILED(result))
    client = nil
    return errors.New(fmt.Sprintf("pfnClientInitialize failed: %x", result))
  }
  fmt.Printf("%#v\n", client)
  return nil
}
func Deinit() error {
  if client == nil {
    return nil
  }

  C.GoVboxClientUninitialize()
  client = nil
  return nil
}

type VirtualBox struct {
  cbox *C.IVirtualBox
  csession *C.ISession
}
type Machine struct {
  Box *VirtualBox
  cmachine *C.IMachine
}

func (vbox *VirtualBox) Init() error {
  C.GoVboxDerpTest()
  return nil

  if err := Init(); err != nil {
    return err
  }

  result := C.GoVboxGetVirtualBox(client, &vbox.cbox)
  if C.GoVboxFAILED(result) != 0 {
    return errors.New("Failed to get IVirtualBox")
  }

  result = C.GoVboxGetSession(client, &vbox.csession)
  if C.GoVboxFAILED(result) != 0 {
    return errors.New("Failed to get ISession")
  }
  return nil
}
func (vbox *VirtualBox) Release() error {
  if vbox.csession != nil {
    result := C.GoVboxISessionRelease(vbox.csession)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New("Failed to release ISession")
    }
    vbox.csession = nil
  }
  if vbox.cbox != nil {
    result := C.GoVboxIVirtualBoxRelease(vbox.cbox)
    if C.GoVboxFAILED(result) != 0 {
      return errors.New("Failed to release IVirtualBox")
    }
    vbox.cbox = nil
  }
  return nil
}
func (vbox* VirtualBox) GetRevision() (int, error) {
  var revision C.ULONG

  result := C.GoVboxGetRevision(vbox.cbox, &revision)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New("Failed to get IVirtualBox revision")
  }

  return int(revision), nil
}

func (vbox* VirtualBox) GetMachines() ([]Machine, error) {
  var cmachinesPtr **C.IMachine
  var machinesCount C.ULONG

  result := C.GoVboxGetMachines(vbox.cbox, &cmachinesPtr, &machinesCount)
  if C.GoVboxFAILED(result) != 0 || cmachinesPtr == nil {
    return nil, errors.New("Failed to get IMachine array")
  }

  sliceHeader := reflect.SliceHeader{
    Data: uintptr(unsafe.Pointer(cmachinesPtr)),
    Len:  int(machinesCount),
    Cap:  int(machinesCount),
  }
  cmachinesSlice := *(*[]*C.IMachine)(unsafe.Pointer(&sliceHeader))

  var machines = make([]Machine, machinesCount)
  for i := range cmachinesSlice {
    machines[i] = Machine{vbox, cmachinesSlice[i]}
  }

  C.free(unsafe.Pointer(cmachinesPtr))

  return machines, nil
}

func (machine* Machine) Release() error {
  if machine.cmachine != nil {
  }
  return nil
}
