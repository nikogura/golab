package sorting

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"log"
	"sort"
	"testing"
)

func TestSortStuff(t *testing.T) {

	expected := testStringsSorted()

	actual := SortStuff(testStringsUnsorted())

	log.Printf("Unsorted: %s", testStringsUnsorted())
	log.Printf("Sorted: %s", actual)

	assert.Equal(t, expected, actual, "Sort works as expected")
}

func TestFourLevel(t *testing.T) {
	data := testThingsUnsorted()

	log.Printf("Unsorted: %s\n", data)

	sort.Sort(FourLevel(data))

	log.Printf("Sorted: %s\n", data)

	expected := testThingsSorted()

	assert.Equal(t, expected, data, "Sort works as expected.")

}
