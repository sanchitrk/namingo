package namingo

import "testing"

func TestGenerateAll(t *testing.T) {
	for i := 0; i < 100; i++ {
		name := Generate(1, " ", "")
		if name == "" {
			t.Error("Name should not be empty")
		}
		t.Log(name)
		name = Generate(2, " ", "")
		if name == "" {
			t.Error("Name should not be empty")
		}
		t.Log(name)
	}
}

func TestGenerateWithNoCase(t *testing.T) {
	// default case is lowercase
	name := Generate(1, " ", "")
	if name == "" {
		t.Error("Name should not be empty")
	}
	//t.Log(name)
	name = Generate(2, " ", "")
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
}

func TestGenerateWithDefaultCase(t *testing.T) {
	name := Generate(1, " ", DefaultCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	name = Generate(2, " ", DefaultCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
}

func TestGenerateWithUpperCase(t *testing.T) {
	name := Generate(1, " ", UpperCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
	name = Generate(2, " ", UpperCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
}

func TestGenerateWithLowerCase(t *testing.T) {
	name := Generate(1, " ", LowerCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
	name = Generate(2, " ", LowerCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
}

func TestGenerateWithTitleCase(t *testing.T) {
	name := Generate(1, " ", TitleCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
	name = Generate(2, " ", TitleCase())
	if name == "" {
		t.Error("Name should not be empty")
	}
	t.Log(name)
}
