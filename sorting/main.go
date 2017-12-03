package sorting

import "sort"

func SortStuff(strs []string) []string {

	sort.Strings(strs)

	return strs
}

type Thing struct {
	One string
	Two string
	Three string
	Four string
}

type FourLevel []Thing

func (a FourLevel) Len() int {
	return len(a)
}

func (a FourLevel) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a FourLevel) Less(i, j int) bool {
	if a[i].One < a[j].One {
		return true
	}

	if a[i].One > a[j].One {
		return false
	}

	if a[i].Two < a[j].Two {
		return true
	}

	if a[i].Two > a[j].Two {
		return false
	}

	if a[i].Three < a[j].Three {
		return true
	}

	return a[i].Four < a[j].Four
}