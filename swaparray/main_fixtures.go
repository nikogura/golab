package swaparray

func initialListEven() []string {
	list := make([]string, 0)
	list = append(list, "a")
	list = append(list, "b")
	list = append(list, "c")
	list = append(list, "d")
	list = append(list, "e")
	list = append(list, "f")

	return list
}

func swappedListEven() []string {
	list := make([]string, 0)
	list = append(list, "f")
	list = append(list, "e")
	list = append(list, "d")
	list = append(list, "c")
	list = append(list, "b")
	list = append(list, "a")

	return list
}

func initialListOdd() []string {
	list := make([]string, 0)
	list = append(list, "a")
	list = append(list, "b")
	list = append(list, "c")
	list = append(list, "d")
	list = append(list, "e")
	list = append(list, "f")
	list = append(list, "g")

	return list
}

func swappedListOdd() []string {
	list := make([]string, 0)
	list = append(list, "g")
	list = append(list, "f")
	list = append(list, "e")
	list = append(list, "d")
	list = append(list, "c")
	list = append(list, "b")
	list = append(list, "a")

	return list
}
