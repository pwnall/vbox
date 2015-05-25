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
  defer os.Remove(imageFile)

  imageSize := uint64(1 << 24)  // 16MB
  progress, err := medium.CreateBaseStorage(imageSize,
      []MediumVariant{MediumVariant_Standard})
  if err != nil {
    t.Fatal(err)
  }

  state, err := medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_Creating && state != MediumState_Created {
    t.Error("Unexpected medium state after creation request: ", state)
  }

  if err = progress.WaitForCompletion(-1); err != nil {
    t.Fatal(err)
  }
  state, err = medium.GetState()
  if err != nil {
    t.Error(err)
  } else if state != MediumState_Created {
    t.Error("Unexpected medium state after creation completed: ", state)
  }
}
