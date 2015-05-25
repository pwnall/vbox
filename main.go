package main

import (
  "runtime"
)

func main() {
  runtime.LockOSThread()
  vb := VirtualBox{}
  vb.Init()
}

