package vbox

import (
  "bytes"
  "io/ioutil"
  "testing"
  "time"
)

func TestMouseKeyboard(t *testing.T) {
  bootReadyImageData, err := ioutil.ReadFile(
      "test_gold/tinycore_boot_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  keyPressImageData, err := ioutil.ReadFile(
      "test_gold/tinycore_keypress_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  /*
  mouseMoveImageData, err := ioutil.ReadFile(
      "test_gold/tinycore_mousemove_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  */

  WithDvdInVm(t, "TinyCore-6.2.iso", true /* disableBootMenu */,
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

    // HACK(pwnall): Wait for the VM to start. It gets stuck if we send Enter
    //               too early.
    time.Sleep(500 * time.Millisecond)

    // Wait to get the boot screen.
    resolution := Resolution{}
    width, height := uint(0), uint(0)
    sentEnter := false
    for {
      if err = display.GetScreenResolution(0, &resolution); err != nil {
        t.Fatal(err)
      }
      width, height = resolution.Width, resolution.Height

      t.Logf("%#v", resolution)

      if resolution.BitsPerPixel == 0 {
        // In text mode.
        if sentEnter == false {
          // Send an Enter key to get rid of the bootloader quickly.
          // NOTE: Failures to issue these codes aren't fatal because they
          //       won't bork the test completely. If we can't press Enter,
          //       it'll just take 60 seconds longer.
          codeCount, err := keyboard.PutScancodes([]int{0x1C, 0x9C})
          if err != nil {
            t.Error(err)
          } else if codeCount != 2 {
            t.Error("Failed to send 2 scancodes, only got: ", codeCount)
          } else {
            sentEnter = true
          }
        }
        // Cannot take screenshots in text mode.
        time.Sleep(100 * time.Millisecond)
        continue
      }

      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }
      // Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/tinycore_boot_screenshot.bin", imageData,
          0644)
      pngData, err := display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/tinycore_boot_screenshot.png", pngData,
            0644)
      }

      time.Sleep(100 * time.Millisecond)
      if bytes.Equal(imageData, bootReadyImageData) {
        break
      }
    }

    // Send Alt+Tab to get the switch menu on the screen.
    // NOTE: Failing to send the keys aborts the test, because otherwise we'd
    //       loop forever waiting for screen changes that never come.
    codeCount, err := keyboard.PutScancodes([]int{0x38, 0x0f, 0x8f, 0xb8})
    if err != nil {
      t.Fatal(err)
    } else if codeCount != 4 {
      t.Fatal("Failed to send 4 scancodes, only got: ", codeCount)
    }

    for {
      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }
      if bytes.Equal(imageData, keyPressImageData) {
        break
      }

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/tinycore_keypress_screenshot.bin", imageData,
          0644)
      imageData, err = display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/tinycore_keypress_screenshot.png",
            imageData, 0644)
      }

      time.Sleep(100 * time.Millisecond)
    }

    // Move mouse to the first item in the pop-up menu.
    // NOTE: Failing to send the event aborts the test, because otherwise we'd
    //       loop forever waiting for screen changes that never come.
    hasAbsolute, err := mouse.GetAbsoluteSupported()
    if err != nil {
      t.Fatal(err)
    } else if hasAbsolute == false {
      t.Fatal("Missing absolute mouse events support: ", hasAbsolute)
    }
    err = mouse.PutEventAbsolute(533, 391, 0, 0, MouseButtonState_None)
    if err != nil {
      t.Fatal(err)
    }

    time.Sleep(5 * time.Second)

    for {
      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }

      /*
      if bytes.Equal(imageData, mouseMoveImageData) {
        break
      }*/

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/tinycore_mousemove_screenshot.bin",
          imageData, 0644)
      imageData, err = display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/tinycore_mousemove_screenshot.png",
            imageData, 0644)
      }

      break
      time.Sleep(100 * time.Millisecond)
    }
  })
}
