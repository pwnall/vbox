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

func TestMedium_CreateBaseStorage(t *testing.T) {
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

  // NOTE: The err value is ignored on purpose. We clean up as well as we can.
  //defer os.Remove(imageFile)

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
}
