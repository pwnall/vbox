package vbox

// This file doesn't have any test cases. It only contains TestMain, which has
// setup and teardown code for all the tests.

import (
  "fmt"
  "io/ioutil"
  "net/http"
  "os"
  "path"
  "testing"
)

func TestMain(m *testing.M) {
  if err := Init(); err != nil {
    fmt.Printf("%v\n", err)
    os.Exit(1)
  }

  // The test images and VM files will be stored in test_tmp.
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

  // We use TinyCore 6.2 because it is a relatively small VM that hosts a GUI,
  // which is necessary to test the Console and related classes.
  tinyCoreFile := path.Join(testDir, "TinyCore-6.2.iso")
  if _, err = os.Stat(tinyCoreFile); err != nil {
    response, err := http.Get("http://distro.ibiblio.org/tinycorelinux/6.x/" +
        "x86/release/TinyCore-6.2.iso")
    if err != nil {
      fmt.Printf("%v\n", err)
      os.Exit(1)
    }
    defer response.Body.Close()
    bodyBytes, err := ioutil.ReadAll(response.Body)
    err = ioutil.WriteFile(tinyCoreFile, bodyBytes, 0644)
    if err != nil {
      fmt.Printf("%v\n", err)
      os.Exit(1)
    }
  }

  result := m.Run()
  Deinit()
  os.Exit(result)
}
