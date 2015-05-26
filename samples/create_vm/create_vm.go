package main

// This sample demonstrates the creation of a VM. After running it, you should
// see a "vbox-sample-vm" machine in the VirtualBox UI.

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

  vm, err := vbox.CreateMachine("vbox-sample-vm", "Linux", "")
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  defer vm.Release()

  if err = vm.Register(); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
}
