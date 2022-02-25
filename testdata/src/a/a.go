package a

func add(a, b int) int {
	// want "ineffectual blank line after the left brace"
	return a + b
	// want "ineffectual blank line before the right brace"
}

func none() {
	// want "ineffectual blank line before the right brace"
}

func success() {
	// This is a successful comment
}

func foo() {
	/*
	   These are successful comments
	*/
	a := 5
	b := 10

	if a < b {
		/*
		   These are successful comments
		*/
		b = add(b, 10)
		/*

		 */
		// want "ineffectual blank line before the right brace"
	}
}

func bar() {
	a := 5
	b := 10

	if a < b {
		// want "ineffectual blank line after the left brace"
		a = add(a, 10)
		// want "ineffectual blank line before the right brace"
	}

	if a > b {
		a = add(a, 10)
	}

	if a == b {
		// This is a successful comment
		a = add(a, 10)
		// This is a successful comment
	}

	// want "ineffectual blank line before the right brace"
}
