package formatting

import (
	"testing"
	"log"
)

func TestFixedColumnOutput(t *testing.T) {
	actual := FixedColumnOutput(thingList())

	for _, line := range actual {
		log.Printf(line)

	}

}