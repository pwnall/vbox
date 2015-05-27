package vbox

import (
  //"bytes"
  "io/ioutil"
  "testing"
  "time"
)

func TestMouseKeyboard(t *testing.T) {
  t.Skip("WiP")

  /*
  goldImageData, err := ioutil.ReadFile("test_gold/bios_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }*/

  WithDvdInVm(t, "TinyCore-6.2.iso",
      func (machine Machine, session Session, console Console) {
    display, err := console.GetDisplay()
    if err != nil {
      t.Fatal(err)
    }
    defer display.Release()

    keyboard, err := console.GetKeyboard()
    if err != nil {
      t.Fatal(err)
    }
    defer keyboard.Release()

    mouse, err := console.GetMouse()
    if err != nil {
      t.Fatal(err)
    }
    defer mouse.Release()

    time.Sleep(10 * time.Second)

    // Wait to get the boot screen.
    resolution := Resolution{}
    for {
      if err = display.GetScreenResolution(0, &resolution); err != nil {
        t.Error(err)
        break
      }
      if resolution.BitsPerPixel == 0 {
        time.Sleep(50 * time.Millisecond)
        continue
      }

      imageData, err := display.TakeScreenShotToArray(0, resolution.Width,
          resolution.Height)
      if err != nil {
        t.Error(err)
        break
      }

      /*
      if bytes.Equal(imageData, goldImageData) {
        break
      }*/

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/tinycore_boot_screenshot.bin", imageData,
          0644)
      imageData, err = display.TakeScreenShotPNGToArray(0, resolution.Width,
          resolution.Height)
      if err == nil {
        ioutil.WriteFile("test_tmp/tinycore_boot_screenshot.png", imageData,
            0644)
      }

      break
      time.Sleep(100 * time.Millisecond)
    }
  })
}
