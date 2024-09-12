package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	sum := Sum(numbers[:])

	assert.Equal(t, 15, sum, "they should be equal")

	numbersFour := [4]int{1, 2, 3, 4}
	sumFour := Sum(numbersFour[:])
	assert.Equal(t, 10, sumFour, "they should be equal")
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5, 6})
	want := 21
	assert.Equal(t, got, want)

	t.Run("more than 3 arrays", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{4, 5, 6}, []int{1, 2, 3}, []int{4, 5, 6})
		want := 42
		assert.Equal(t, got, want)
	})
}
