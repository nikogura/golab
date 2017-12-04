package relations

import (
	"container/list"
	"log"
)

// Given two Persons, each having a Name, and a list of relations (other people), find the degree to which they are related.
// I.e. Walk the (possibly cyclic) web of relations and return the shortest path between persion A and person B without needing to traverse the entire web.
// Do it quickly and efficiently of course.

// Person stores a person.  A name- which must be unique and a list of relations that are people.
type Person struct {
	// names are assumed to be unique for the purposes of this exercise
	Name string
	// every person has 0 or more relations
	Relations []string
}

// TERSE VERSION (anonymous funcs)

// FindDegreesOfSeparationAnon returns the shortest list of relations between two people.
// Same as the FindDegreesOfSeparation below, but using anonymous functions.  Terser, but a lot less readable methinks.
func FindDegreesOfSeparationAnon(people map[string]Person, start string, target string) (found bool, journey []string, degrees int) {

	// need a list of people we've seen, so we don't get into an endless loop
	// It's a map so lookups are in constant time
	seen := make(map[string]bool)

	// Need a struct with a list to track the path from Start to Finish
	journey = []string{start}

	// return early if someone is playing silly buggers
	if start == target {
		log.Printf("Found em! 0 degrees of separation.")
		return true, journey, 0
	}

	// make a processing stack
	stack := list.New()

	// add this node to the stack
	stack.PushFront(journey)

	// Parse the stack
	// check each node's children, and if we find the target, return success and the journey.
	// if not, push the child on to the processing stack because we want to evaluate all children of this node before descending a generation. (breadth first search)
	for e := stack.Front(); e != nil; e = e.Next() {
		// copy the journey
		thisJourney := e.Value.([]string)

		// start the next iteration with the child, which happens to be the end of the list
		relativeStart := thisJourney[len(thisJourney)-1]

		success, childJourney := func(people map[string]Person, start string, target string, seen map[string]bool, journey []string, stack *list.List) (success bool, childJourney []string) {
			for _, child := range people[start].Relations {
				success, childJourney = func(name string, target string, seen map[string]bool, journey []string) (success bool, myJourney []string) {
					// copy the track to this point
					myJourney = []string(journey)

					// add this person to the track
					myJourney = append(myJourney, name)

					// add the person to the seen list
					seen[name] = true

					// if we found the goal
					if name == target {
						return true, myJourney
					}

					return false, myJourney
				}(child, target, seen, journey)

				// if we find the target, return it and the journey that got us here
				if success {
					return success, childJourney
				}

				// if not, put it on the end of the stack for further processing
				stack.PushBack(childJourney)
			}

			return false, childJourney
		}(people, relativeStart, target, seen, thisJourney, stack)

		// if we found it in the children, return it
		if success {
			return success, childJourney, len(childJourney) - 1
		}
	}

	// have to return something
	return false, journey, -1
}

// READABLE VERSION

// FindDegreesOfSeparation returns the shortest list of relations between two people.
func FindDegreesOfSeparation(people map[string]Person, start string, target string) (found bool, journey []string, degrees int) {

	// need a list of people we've seen, so we don't get into an endless loop
	// It's a map so lookups are in constant time
	seen := make(map[string]bool)

	// Need a struct with a list to track the path from Start to Finish
	journey = []string{start}

	// return early if someone is playing silly buggers
	if start == target {
		log.Printf("Found em! 0 degrees of separation.")
		return true, journey, 0
	}

	// make a processing stack
	stack := list.New()

	// add this node to the stack
	stack.PushFront(journey)

	// Parse the stack
	// check each node's children, and if we find the target, return success and the journey.
	// if not, push the child on to the processing stack because we want to evaluate all children of this node before descending a generation. (breadth first search)
	for e := stack.Front(); e != nil; e = e.Next() {
		// copy the journey
		thisJourney := e.Value.([]string)

		// start the next iteration with the child, which happens to be the end of the list
		relativeStart := thisJourney[len(thisJourney)-1]

		success, childJourney := CheckChildren(people, relativeStart, target, seen, thisJourney, stack)
		// if we found it in the children, return it
		if success {
			return success, childJourney, len(childJourney) - 1
		}
	}

	// have to return something
	return false, journey, -1
}

func CheckChildren(people map[string]Person, start string, target string, seen map[string]bool, journey []string, stack *list.List) (success bool, childJourney []string) {

	// check each child of this node
	for _, child := range people[start].Relations {
		success, childJourney = CheckPerson(child, target, seen, journey)

		// if we find the target, return it and the journey that got us here
		if success {
			return success, childJourney
		}

		// if not, put it on the end of the stack for further processing
		stack.PushBack(childJourney)
	}

	return false, childJourney
}

func CheckPerson(name string, target string, seen map[string]bool, journey []string) (success bool, myJourney []string) {
	// copy the track to this point
	myJourney = []string(journey)

	// add this person to the track
	myJourney = append(myJourney, name)

	// add the person to the seen list
	seen[name] = true

	// if we found the goal
	if name == target {
		return true, myJourney
	}

	return false, myJourney
}
