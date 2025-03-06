package fx

func Fibonacci(n uint64) []uint64 {
  sequence := []uint64{1, 1}

  for range n {
    sequence = append(sequence, sequence[len(sequence)-1]+sequence[len(sequence)-2])
  }

  if n == 1 {
    return []uint64{1}
  } else if n == 0 {
    return []uint64{}
  }

  return sequence
}
