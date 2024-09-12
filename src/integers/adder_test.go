package integers

import (
	"testing"
  "fmt"

  "github.com/stretchr/testify/assert"
)

func TestAdder(t *testing.T) {
	sum := Add(3, 4)

  assert.Equal(t, 7, sum, "they should be equal")
}

func ExampleAdd() {
  sum := Add(1, 5)
  fmt.Println(sum)
  // Output: 6
}
