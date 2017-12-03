package formatting

func thing1() Thing {
	return Thing{
		Name: "foo",
		Version: "1.2.3",
		Description: "The quick fox jumped over the lazy brown dog.",
	}
}

func thing2() Thing {
	return Thing{
		Name: "goongala",
		Version: "1.5.3",
		Description: "Default description.",
	}
}

func thing3() Thing {
	return Thing{
		Name: "glortswaggle",
		Version: "2.2.2",
		Description: "Doodle Doodle Dee, Wubba Wubba Wubba.",
	}
}

func thing4() Thing {
	return Thing{
		Name: "freddled-gruntbuggly",
		Version: "9.2.3",
		Description: "Lorem ipsum dolor blah blah blah blah blah blab",
	}
}

func thing5() Thing {
	return Thing{
		Name: "frobnitz",
		Version: "1.0.0",
		Description: "word.",
	}
}

func thingList() []Thing {
	list := make([]Thing, 0)

	list = append(list, thing1())
	list = append(list, thing2())
	list = append(list, thing3())
	list = append(list, thing4())
	list = append(list, thing5())

	return list
}
