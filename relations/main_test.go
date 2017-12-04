package relations

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestFind0DegreesOfSeparation(t *testing.T) {
	// 0 Degrees
	start := "larry"
	target := "larry"
	expected := 0
	success, _, actual := FindDegreesOfSeparation(testPeople(), start, target)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFind1DegreeOfSeparation(t *testing.T) {
	// 1 Degree
	start := "larry"
	target := "malcom"
	expected := 1
	success, journey, actual := FindDegreesOfSeparation(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")
	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFind3DegreesOfSeparation(t *testing.T) {
	// 3 Degrees
	start := "larry"
	target := "betty"
	expected := 3
	success, journey, actual := FindDegreesOfSeparation(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")

	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFindnDegreesOfSeparation(t *testing.T) {
	// 5 Degrees
	start := "groucho"
	target := "betty"
	expected := 5
	success, journey, actual := FindDegreesOfSeparation(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")

	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFind0DegreesOfSeparationAnon(t *testing.T) {
	// 0 Degrees
	start := "larry"
	target := "larry"
	expected := 0
	success, _, actual := FindDegreesOfSeparationAnon(testPeople(), start, target)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFind1DegreeOfSeparationAnon(t *testing.T) {
	// 1 Degree
	start := "larry"
	target := "malcom"
	expected := 1
	success, journey, actual := FindDegreesOfSeparationAnon(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")
	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFind3DegreesOfSeparationAnon(t *testing.T) {
	// 3 Degrees
	start := "larry"
	target := "betty"
	expected := 3
	success, journey, actual := FindDegreesOfSeparationAnon(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")

	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}

func TestFindnDegreesOfSeparationAnon(t *testing.T) {
	// 5 Degrees
	start := "groucho"
	target := "betty"
	expected := 5
	success, journey, actual := FindDegreesOfSeparationAnon(testPeople(), start, target)

	journeyPath := strings.Join(journey, " -> ")

	log.Printf("Success: %t Journey: %s", success, journeyPath)

	assert.True(t, success, "Found it")
	assert.Equal(t, expected, actual, fmt.Sprintf("%d degrees of separation between %s and %s", actual, start, target))

}
