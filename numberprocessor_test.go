package numberprocessor_test

import (
  "github.com/karngyan/numberprocessor"
  "testing"
)

func TestProcessData(t *testing.T) {
  data := make([]*numberprocessor.Number, 5)
  data[0] = &numberprocessor.Number{Value: 1}
  data[1] = &numberprocessor.Number{Value: 2}
  // we deliberately leave data[2], data[3] and data[4] as nil

  result := numberprocessor.ProcessData(data)
  if result != 3 {
    t.Errorf("Sum was incorrect, got: %d, want: %d.", result, 3)
  }
}
