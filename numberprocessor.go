package numberprocessor

type Number struct {
  Value int
}

func ProcessData(data []*Number) int {
  var sum int
  for _, number := range data {
    sum += number.Value
  }
  return sum
}
