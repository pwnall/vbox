package vbox

import (
  "strings"
  "testing"
)

func TestAppVersion(t *testing.T) {
  if AppVersion <= 4003000 {
    t.Error("AppVersion below 4.3: ", AppVersion)
  }
}

func TestGetRevision(t *testing.T) {
  revision, err := GetRevision()
  if err != nil {
    t.Fatal(err)
  }
  if revision <= 100000 {
    t.Error("Revision below 100000: ", revision)
  }
}

func TestComposeMachineFilename(t *testing.T) {
  path, err := ComposeMachineFilename("TestVM", "", "/test/vm/path")
  if err != nil {
    t.Fatal(err)
  }
  if path != "/test/vm/path/TestVM/TestVM.vbox" {
    t.Error("Wrong VM filename when given baseFolder: ", path)
  }

  path, err = ComposeMachineFilename("TestVM", "", "")
  if err != nil {
    t.Fatal(err)
  }
  if !strings.Contains(path, "VirtualBox") {
    t.Error("VM filename without baseFolder doesn't have VirtualBox: ", path)
  }
  if !strings.Contains(path, "TestVM.vbox") {
    t.Error("VM filename without baseFolder doesn't have TestVM.vbox: ", path)
  }
}

