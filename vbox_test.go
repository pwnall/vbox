package main

import (
  "fmt"
  "os"
  "runtime"
  "testing"
  //"image"
)

func TestVirtualBox_GetRevision(t *testing.T) {
  fmt.Print("Will init\n")

  var vb VirtualBox
  if err := vb.Init(); err != nil {
    t.Fatal("VirtualBox.Init() failed", err)
  }
  fmt.Print("Passed init\n")

  revision, err := vb.GetRevision()
  if err != nil {
    t.Fatal("GetRevision failed", err)
  }
  t.Log("VirtualBox revision: ", revision)
  if revision == 0 {
    t.Error("GetRevision failed")
  }

  vb.Release()
}

func TestVirtualBox_GetMachines(t *testing.T) {
  t.Skip("crashes")

  var vb VirtualBox
  if err := vb.Init(); err != nil {
    t.Error("VirtualBox.Init() failed", err)
  }

  machines, err := vb.GetMachines()
  if err != nil {
    t.Error("GetMachines failed", err)
  }
  if len(machines) == 0 {
    t.Error("GetMachines failed")
  }

  vb.Release()
}

func TestMain(m *testing.M) {
  _ = runtime.LockOSThread
  //runtime.LockOSThread()
  vb := VirtualBox{}
  vb.Init()

  result := m.Run()
  Deinit()
  os.Exit(result)
}
