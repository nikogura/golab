package sorting

func testStringsUnsorted() []string {
	strs := []string{"f", "w", "a"}

	return strs
}

func testStringsSorted() [] string {
	sorted := []string{"a", "f", "w"}

	return sorted
}

func thing1() Thing {
	return Thing{One: "alpha", Two: "beta", Three: "zed", Four: "monkey"}
}

func thing2() Thing {
	return Thing{One: "alpha", Two: "beta", Three: "kappa", Four: "delta"}

}

func thing3() Thing {
	return Thing{One: "alpha", Two: "alpha", Three: "yellow", Four: "green"}

}

func thing4() Thing {
	return Thing{One: "alpha", Two: "beta", Three: "zed", Four: "four"}

}

func testThingsUnsorted() []Thing {
	things := make([]Thing, 0)
	things = append(things, thing1())
	things = append(things, thing2())
	things = append(things, thing3())
	things = append(things, thing4())

	return things
}

func testThingsSorted() []Thing {
	things := make([]Thing, 0)
	things = append(things, thing3())
	things = append(things, thing2())
	things = append(things, thing4())
	things = append(things, thing1())

	return things
}
