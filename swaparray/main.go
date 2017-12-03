package swaparray

import "log"

func SwapArrayInPlace(list []string) (swapped []string) {
	for i, j := 0, len(list)-1; i < len(list); i, j = i+1, j-1 {

		log.Printf("Swapping %d and %d", i, j)
		list[i], list[j] = list[j], list[i]

		if i == j-1 || i == j {
			break
		}
	}

	return list
}
