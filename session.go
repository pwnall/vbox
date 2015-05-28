package vbox

/*
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/include
#cgo CFLAGS: -I third_party/VirtualBoxSDK/sdk/bindings/c/glue
#cgo LDFLAGS: -ldl -lpthread

#include <stdlib.h>
#include "c_wrappers/session.c"
*/
import "C"  // cgo's virtual package

import (
  "errors"
  "fmt"
)

// A session gets associated to a VM lock.
type Session struct {
  csession *C.ISession
}

// Init creates the session object on the VirtualBox side.
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

// LockMachine obtains a lock on a VM, so it can be modified or started.
// It returns any error encountered.
func (session *Session) LockMachine(machine Machine, lockType LockType) error {
  result := C.GoVboxLockMachine(machine.cmachine, session.csession,
      C.PRUint32(lockType))
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(fmt.Sprintf("Failed to lock IMachine: %x", result))
  }
  return nil
}

// UnlockMachine releases the VM locked by this session.
// It returns any error encountered.
func (session *Session) UnlockMachine() error {
  result := C.GoVboxUnlockMachine(session.csession)
  if C.GoVboxFAILED(result) != 0 {
    return errors.New(
        fmt.Sprintf("Failed to unlock ISession machine: %x", result))
  }
  return nil
}

// GetConsole obtains the controls for the VM associated with this session.
// The call fails unless the VM associated with this session has started.
// It returns a new Console instance and any error encountered.
func (session *Session) GetConsole() (Console, error) {
  var console Console
  result := C.GoVboxGetSessionConsole(session.csession, &console.cconsole)
  if C.GoVboxFAILED(result) != 0 || console.cconsole == nil {
    return console, errors.New(
        fmt.Sprintf("Failed to get ISession console: %x", result))
  }
  return console, nil
}

// GetMachine obtains the VM associated with this session.
// It returns a new Machine instance and any error encountered.
func (session *Session) GetMachine() (Machine, error) {
  var machine Machine
  result := C.GoVboxGetSessionMachine(session.csession, &machine.cmachine)
  if C.GoVboxFAILED(result) != 0 || machine.cmachine == nil {
    return machine, errors.New(
        fmt.Sprintf("Failed to get ISession machine: %x", result))
  }
  return machine, nil
}

// GetState obtains the current state of this session.
// It returns the SessionState and any error encountered.
func (session *Session) GetState() (SessionState, error) {
  var cstate C.PRUint32
  result := C.GoVboxGetSessionState(session.csession, &cstate)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get ISession state: %x", result))
  }
  return SessionState(cstate), nil
}

// GetType obtains the session's type.
// It returns the SessionType and any error encountered.
func (session *Session) GetType() (SessionType, error) {
  var ctype C.PRUint32
  result := C.GoVboxGetSessionType(session.csession, &ctype)
  if C.GoVboxFAILED(result) != 0 {
    return 0, errors.New(
        fmt.Sprintf("Failed to get ISession type: %x", result))
  }
  return SessionType(ctype), nil
}

// Release frees up the associated VirtualBox data.
// After the call, this instance is invalid, and using it will cause errors.
// It returns any error encountered.
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
