package main

// This sample demonstrates posting mouse events to a VM. It assumes that the
// "vbox-sample-vm" VM is up and running.

import (
  "fmt"
  "os"
  "github.com/pwnall/vbox"
)

func main() {
  if err := vbox.Init(); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  machine, err := vbox.FindMachine("vbox-sample-vm")
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer machine.Release()

  session := vbox.Session{}
  if err := session.Init(); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer session.Release()

  if err := session.LockMachine(machine, vbox.LockType_Shared); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer func() {
    if err := session.UnlockMachine(); err != nil {
      fmt.Printf("%v\n", err)
    }
  }()

  console, err := session.GetConsole()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer console.Release()

  mouse, err := console.GetMouse()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer mouse.Release()

  ss, err := mouse.GetRelativeSupported()
  fmt.Printf("Relative support: %v %v\n", ss, err)

  ss, err = mouse.GetAbsoluteSupported()
  fmt.Printf("Absolute support: %v %v\n", ss, err)

  err = mouse.PutEventAbsolute(200, 200, 0, 0, vbox.MouseButtonState_None)
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
}
