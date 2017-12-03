package template

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestSimpleTemplate(t *testing.T) {
	output := SimpleTemplate()

	log.Printf("Output: %s", output)

	assert.Equal(t, "Dot:map[1:S 2:H 3:C 4:D]", output, "Output matches expectations")
}
