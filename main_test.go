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

  // We use Lubuntu because it's not huge and has working VirtualBox mouse
  // support.
  isoFile := path.Join(testDir, "lubuntu-15.04.iso")
  if _, err = os.Stat(isoFile); err != nil {
    fmt.Printf("Lubuntu ISO not found, downloading\n")
    response, err := http.Get("http://cdimage.ubuntu.com/lubuntu/releases/" +
        "15.04/release/lubuntu-15.04-desktop-i386.iso")
    if err != nil {
      fmt.Printf("%v\n", err)
      os.Exit(1)
    }
    defer response.Body.Close()
    bodyBytes, err := ioutil.ReadAll(response.Body)
    err = ioutil.WriteFile(isoFile, bodyBytes, 0644)
    if err != nil {
      fmt.Printf("%v\n", err)
      os.Exit(1)
    }
    fmt.Printf("Done downloading Lubuntu ISO\n")
  }

  result := m.Run()
  Deinit()
  os.Exit(result)
}
