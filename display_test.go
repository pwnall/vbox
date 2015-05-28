package vbox

import (
  "bytes"
  "io/ioutil"
  "image"
  "image/png"
  "testing"
  "time"
)

func TestDisplay(t *testing.T) {
  goldImageData, err := ioutil.ReadFile("test_gold/bios_screenshot.bin")
  if err != nil {
    t.Fatal(err)
  }

  WithDvdInVm(t, "", false /* disableBootMenu */,
      func (machine Machine, session Session, console Console) {
    display, err := console.GetDisplay()
    if err != nil {
      t.Fatal(err)
    }
    defer display.Release()

    // Wait until the VM display initializes.
    resolution := Resolution{}
    for {
      if err = display.GetScreenResolution(0, &resolution); err != nil {
        t.Error(err)
        break
      }
      if resolution.BitsPerPixel != 0 {
        break
      }
      t.Logf("%#v", resolution)
      time.Sleep(50 * time.Millisecond)
    }

    // Wait to get the boot failure screen.
    for {
      if err = display.GetScreenResolution(0, &resolution); err != nil {
        t.Error(err)
        break
      }
      imageData, err := display.TakeScreenShotToArray(0, resolution.Width,
          resolution.Height)
      if err != nil {
        t.Error(err)
        break
      }

      if bytes.Equal(imageData, goldImageData) {
        break
      }
      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/bios_screenshot.bin", imageData, 0644)
      time.Sleep(50 * time.Millisecond)
    }

    imageData, err := display.TakeScreenShotPNGToArray(0, resolution.Width,
        resolution.Height)
    if err != nil {
      t.Error(err)
    } else {
      // NOTE: Dumping the screenshot to the filesystem to ease debugging.
      ioutil.WriteFile("test_tmp/bios_screenshot.png", imageData, 0644)

      pngImage, err := png.Decode(bytes.NewReader(imageData))
      if err != nil {
        t.Error(err)
      } else {
        bounds := pngImage.Bounds()
        pngWidth := bounds.Max.X - bounds.Min.X
        pngHeight := bounds.Max.Y - bounds.Min.Y
        if pngWidth != int(resolution.Width) {
          t.Error("Incorrect PNG width: ", pngWidth, " expected:",
              resolution.Width)
        }
        if pngHeight != int(resolution.Height) {
          t.Error("Incorrect PNG height. X: ", pngHeight, " expected:",
              resolution.Height)
        }
        rgbaImage, success := pngImage.(*image.RGBA)
        if !success {
          t.Error("Failed to cast image to image.RGBA")
        } else if !bytes.Equal(rgbaImage.Pix, goldImageData) {
          t.Error("Incorrect PNG pixel data")
        }
      }
    }

    /* NOTE: The fast screenshot method doesn't seem to work.
    bufferSize := int(resolution.Width * resolution.Height * 4 + 1000)
    imageBuffer := make([]byte, bufferSize)
    imageData, err = display.TakeScreenShot(0, imageBuffer, resolution.Width,
        resolution.Height)
    if err != nil {
      t.Error(err)
    } else {
      if len(imageData) != bufferSize {
        t.Error("Incorrect fast screenshot slice length: ", len(imageData),
            " expected: ", bufferSize)
      } else if !bytes.Equal(imageData, goldImageData) {
        t.Error("Incorrect fast screenshot pixel data")
      }
    }
    */
  })
}
