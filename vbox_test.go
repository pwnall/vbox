package vbox

import (
  "fmt"
  "os"
  "path"
  "strings"
  "testing"
  //"image"
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

func TestMain(m *testing.M) {
  if err := Init(); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  cwd, err := os.Getwd()
  if err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }
  testDir := path.Join(cwd, "test_tmp")
  if err = os.MkdirAll(testDir, 0777); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  result := m.Run()
  Deinit()
  os.Exit(result)
}
