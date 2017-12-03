package swaparray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapArrayInPlace(t *testing.T) {
	expected := swappedListEven()
	actual := SwapArrayInPlace(initialListEven())

	assert.Equal(t, expected, actual, "Even List Successfully Swapped.")

	expected = swappedListOdd()
	actual = SwapArrayInPlace(initialListOdd())

	assert.Equal(t, expected, actual, "Odd List Successfully Swapped.")
}
