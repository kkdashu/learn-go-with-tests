package array

func Sum(arr []int) int {
  sum := 0
  for _, num := range arr {
    sum += num
  }
  return sum
}

func SumAll(arrs ...[]int) int {
  var sum int
  for _, arr := range arrs {
    sum += Sum(arr)
  }
  return sum
}
