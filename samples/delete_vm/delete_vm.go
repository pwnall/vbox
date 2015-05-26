package main

// This sample demonstrates the removal of a VM. After running it, the
// "vbox-sample-vm" machine should disappear from the VirtualBox UI.

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

  vms, err := vbox.GetMachines()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  foundVm := false
  vm := vbox.Machine{}
  for _, ivm := range vms {
    name, err := ivm.GetName()
    if err != nil {
      fmt.Printf("%v\n", err)
    }
    if name == "vbox-sample-vm" {
      foundVm = true
      vm = ivm
    } else {
      ivm.Release()
    }
  }

  if foundVm == false {
    fmt.Printf("Could not find vbox-sample-vm on the VM list\n")
    os.Exit(1)
  }
  defer vm.Release()

  // NOTE: The exact cleanup mode is inconsequential for this sample, but we
  //       want to show the generally recommended value.
  media, err := vm.Unregister(vbox.CleanupMode_DetachAllReturnHardDisksOnly)
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  progress, err := vm.DeleteConfig(media)
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  if err = progress.WaitForCompletion(-1); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  percent, err := progress.GetPercent()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  if percent != 100 {
    fmt.Printf("Config deletion stopped at %d%%\n", percent)
    os.Exit(1)
  }

  result, err := progress.GetResultCode()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  if result != 0 {
    fmt.Printf("Config deletion failed with code %x\n", percent)
    os.Exit(1)
  }
}
