package vbox

import (
  "bytes"
  "io/ioutil"
  "testing"
  "time"
)

func TestMouseKeyboard(t *testing.T) {
  bootReadyImageData, err := ioutil.ReadFile(
      "test_gold/lubuntu_boot_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  mouseMoveImageData, err := ioutil.ReadFile(
      "test_gold/lubuntu_mousemove_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  mouseClickImageData, err := ioutil.ReadFile(
      "test_gold/lubuntu_mouseclick_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }
  keyPressImageData, err := ioutil.ReadFile(
      "test_gold/lubuntu_keypress_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }

  WithDvdInVm(t, "lubuntu-15.04.iso", true /* disableBootMenu */,
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
        // Cannot take screenshots in text mode.
        time.Sleep(100 * time.Millisecond)
        continue
      }

      if resolution.BitsPerPixel == 16 {
        // In 16-bit graphics mode.
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
            sentEnter = false  // HACK
          }
        }
        // Don't bother taking screenshots while booting.
        time.Sleep(100 * time.Millisecond)
        continue
      }

      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }

      // Lubuntu displays the clock at the bottom of the screen, so we trim it.
      imageData = imageData[0:((len(imageData) * 7) / 8)]

      // Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/lubuntu_boot_screenshot.bin", imageData,
          0644)
      pngData, err := display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/lubuntu_boot_screenshot.png", pngData,
            0644)
      }

      if bytes.Equal(imageData, bootReadyImageData) {
        break
      }
      time.Sleep(200 * time.Millisecond)
    }

    // Move mouse to the left of the screen.
    // NOTE: Failing to send the event aborts the test, because otherwise we'd
    //       loop forever waiting for screen changes that never come.
    hasAbsolute, err := mouse.GetAbsoluteSupported()
    if err != nil {
      t.Fatal(err)
    } else if hasAbsolute == false {
      t.Fatal("Missing absolute mouse events support: ", hasAbsolute)
    }
    err = mouse.PutEventAbsolute(400, 200, 0, 0, MouseButtonState_None)
    if err != nil {
      t.Fatal(err)
    }

    for {
      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }
      // Lubuntu displays the clock at the bottom of the screen, so we trim it.
      imageData = imageData[0:((len(imageData) * 7) / 8)]

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/lubuntu_mousemove_screenshot.bin",
          imageData, 0644)
      pngData, err := display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/lubuntu_mousemove_screenshot.png",
            pngData, 0644)
      }

      if bytes.Equal(imageData, mouseMoveImageData) {
        break
      }
      time.Sleep(100 * time.Millisecond)
    }

    // Move the mouse slightly and click.
    // NOTE: Failing to send the event aborts the test, because otherwise we'd
    //       loop forever waiting for screen changes that never come.
    hasRelative, err := mouse.GetRelativeSupported()
    if err != nil {
      t.Fatal(err)
    } else if hasRelative == false {
      t.Fatal("Missing absolute mouse events support: ", hasRelative)
    }
    err = mouse.PutEvent(-5, 2, 0, 0, MouseButtonState_RightButton)
    if err != nil {
      t.Fatal(err)
    }

    for {
      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }
      // Lubuntu displays the clock at the bottom of the screen, so we trim it.
      imageData = imageData[0:((len(imageData) * 7) / 8)]

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/lubuntu_mouseclick_screenshot.bin",
          imageData, 0644)
      pngData, err := display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/lubuntu_mouseclick_screenshot.png",
            pngData, 0644)
      }

      if bytes.Equal(imageData, mouseClickImageData) {
        break
      }
      time.Sleep(100 * time.Millisecond)
    }

    // Send Ctrl+A to get stuff highlighted.
    // NOTE: Failing to send the keys aborts the test, because otherwise we'd
    //       loop forever waiting for screen changes that never come.
    codeCount, err := keyboard.PutScancodes([]int{0x01, 0x91})
    if err != nil {
      t.Fatal(err)
    } else if codeCount != 2 {
      t.Fatal("Failed to send 2 scancodes, only got: ", codeCount)
    }

    for {
      imageData, err := display.TakeScreenShotToArray(0, width, height)
      if err != nil {
        t.Error(err)
        break
      }
      // Lubuntu displays the clock at the bottom of the screen, so we trim it.
      imageData = imageData[0:((len(imageData) * 7) / 8)]

      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/lubuntu_keypress_screenshot.bin", imageData,
          0644)
      pngData, err := display.TakeScreenShotPNGToArray(0, width, height)
      if err == nil {
        ioutil.WriteFile("test_tmp/lubuntu_keypress_screenshot.png",
            pngData, 0644)
      }

      if bytes.Equal(imageData, keyPressImageData) {
        break
      }
      time.Sleep(100 * time.Millisecond)
    }
  })
}
