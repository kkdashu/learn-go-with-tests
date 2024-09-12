package iteration

func Repeat(c string) string {
  count := 5
  var repeated string
  for range(count) {
    repeated += c
  }
  return repeated
}
