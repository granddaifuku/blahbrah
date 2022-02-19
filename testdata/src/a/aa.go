package a

type cat struct {
	name string
	age  int
}

func foo() {
	_ = cat{
		// want "ineffectual blank line after the left brace"
		name: "Haru",
		age:  2,
	}

	_ = cat{
		name: "Hime",
		age:  1,
	}
}
