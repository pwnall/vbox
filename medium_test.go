package vbox

import (
  "os"
  "path"
  "testing"
)

func TestCreateMedium(t *testing.T) {
  cwd, err := os.Getwd()
  if err != nil {
    t.Fatal(err)
  }
  testDir := path.Join(cwd, "test_tmp")

  imageFile := path.Join(testDir, "test_disk.vdi")
  medium, err := CreateHardDisk("VDI", imageFile)
  if err != nil {
    t.Fatal(err)
  }
  defer medium.Release()

  location, err := medium.GetLocation()
  if err != nil {
    t.Error(err)
  } else if location != imageFile {
    t.Error("Wrong medium location: ", location, " expected: ", imageFile)
  }

  state, err := medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_NotCreated {
    t.Error("Wrong medium state: ", state)
  }
}

func TestMedium_CreateBaseStorage_DeleteStorage(t *testing.T) {
  cwd, err := os.Getwd()
  if err != nil {
    t.Fatal(err)
  }
  testDir := path.Join(cwd, "test_tmp")

  imageFile := path.Join(testDir, "test_disk.vdi")
  medium, err := CreateHardDisk("VDI", imageFile)
  if err != nil {
    t.Fatal(err)
  }
  defer medium.Release()

  // NOTE: VirtualBox errors out if we try to create an image that already
  //       exists, so we make sure we clean up any left overs from previous
  //       failed/crashed tests.
  if _, err = os.Stat(imageFile); err == nil {
    if err = os.Remove(imageFile); err != nil {
      t.Fatal(err)
    }
  }

  imageSize := uint64(1 << 24)  // 16MB
  progress, err := medium.CreateBaseStorage(imageSize,
      []MediumVariant{MediumVariant_Standard, MediumVariant_NoCreateDir})
  if err != nil {
    t.Fatal(err)
  }

  if err = progress.WaitForCompletion(10000); err != nil {
    t.Fatal(err)
  }

  percent, err := progress.GetPercent()
  if err != nil {
    t.Error(err)
  }
  if percent < 100 {
    t.Error("Progress percent below 100 after waiting for completion: ",
        percent)
  }

  code, err := progress.GetResultCode()
  if err != nil {
    t.Error(err)
  }
  if code != 0 {
    t.Error("Progress has a non-zero error code:", code)
  }

  state, err := medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_Created {
    t.Error("Unexpected medium state after creation completed: ", state)
  }

  size, err := medium.GetSize()
  if err != nil {
    t.Error(err)
  } else if size <= 0 {
    t.Error("Invalid image file size: ", size)
  }

  progress, err = medium.DeleteStorage()
  if err != nil {
    t.Fatal(err)
  }
  if err = progress.WaitForCompletion(10000); err != nil {
    t.Fatal(err)
  }

  state, err = medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_NotCreated {
    t.Error("Unexpected medium state after deletion completed: ", state)
  }
}

func TestOpenMedium_Medium_Close(t *testing.T) {
  cwd, err := os.Getwd()
  if err != nil {
    t.Fatal(err)
  }
  testDir := path.Join(cwd, "test_tmp")

  imageFile := path.Join(testDir, "TinyCore-6.2.iso")
  imageStat, err := os.Stat(imageFile)
  if err != nil {
    t.Fatal(err)
  }

  medium, err := OpenMedium(imageFile, DeviceType_Dvd, AccessMode_ReadOnly,
      false)
  if err != nil {
    t.Fatal(err)
  }
  defer medium.Release()
  defer medium.Close()

  location, err := medium.GetLocation()
  if err != nil {
    t.Error(err)
  } else if location != imageFile {
    t.Error("Wrong medium location: ", location, " expected: ", imageFile)
  }

  state, err := medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_Created {
    t.Error("Wrong medium state: ", state)
  }

  imageSize := uint64(imageStat.Size())
  size, err := medium.GetSize()
  if err != nil {
    t.Error(err)
  } else if size != imageSize {
    t.Error("Invalid image file size: ", size, " expected: ", imageSize)
  }

  if err = medium.Close(); err != nil {
    t.Error(err)
  }
}
