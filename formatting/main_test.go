package formatting

import (
	"log"
	"testing"
)

func TestFixedColumnOutput(t *testing.T) {
	actual := FixedColumnOutput(thingList())

	for _, line := range actual {
		log.Printf(line)

	}

}
