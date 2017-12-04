package relations

func testPeople() map[string]Person {
	people := make(map[string]Person)

	people["larry"] = Person{
		Name: "larry",
		Relations: []string{
			"moe",
			"curly",
			"shemp",
			"malcom",
			"kaylee",
		},
	}

	people["moe"] = Person{
		Name: "moe",
		Relations: []string{
			"curly",
			"barney",
			"harpo",
			"fred",
		},
	}

	people["curly"] = Person{
		Name: "curly",
		Relations: []string{
			"moe",
			"chico",
		},
	}

	people["shemp"] = Person{
		Name: "shemp",
		Relations: []string{
			"wash",
			"zoe",
			"malcom",
		},
	}

	people["groucho"] = Person{
		Name: "groucho",
		Relations: []string{
			"zoe",
			"simon",
		},
	}

	people["chico"] = Person{
		Name: "chico",
		Relations: []string{
			"groucho",
			"simon",
			"wilma",
			"fred",
		},
	}

	people["harpo"] = Person{
		Name: "Harpo",
		Relations: []string{
			"fred",
			"larry",
			"river",
		},
	}

	people["zeppo"] = Person{
		Name: "zeppo",
		Relations: []string{
			"kaylee",
			"moe",
		},
	}

	people["malcom"] = Person{
		Name: "malcom",
		Relations: []string{
			"river",
			"kaylee",
			"zoe",
		},
	}

	people["zoe"] = Person{
		Name: "zoe",
		Relations: []string{
			"wash",
			"river",
			"kaylee",
			"malcom",
		},
	}

	people["wash"] = Person{
		Name: "wash",
		Relations: []string{
			"zoe",
		},
	}

	people["kaylee"] = Person{
		Name: "kaylee",
		Relations: []string{
			"malcom",
			"chico",
			"zeppo",
			"curly",
		},
	}

	people["simon"] = Person{
		Name: "simon",
		Relations: []string{
			"river",
			"jayne",
		},
	}

	people["river"] = Person{
		Name: "river",
		Relations: []string{
			"simon",
			"malcom",
			"curly",
			"wilma",
		},
	}

	people["jayne"] = Person{
		Name: "jayne",
		Relations: []string{
			"kaylee",
			"river",
			"zeppo",
		},
	}

	people["wilma"] = Person{
		Name: "Wilma",
		Relations: []string{
			"fred",
			"barney",
		},
	}

	people["fred"] = Person{
		Name: "fred",
		Relations: []string{
			"barney",
			"wilma",
			"groucho",
			"jayne",
		},
	}

	people["barney"] = Person{
		Name: "barney",
		Relations: []string{
			"larry",
			"zeppo",
			"malcom",
			"betty",
		},
	}

	people["betty"] = Person{
		Name: "betty",
		Relations: []string{
			"barney",
		},
	}

	return people
}
