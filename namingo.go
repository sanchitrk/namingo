package namingo

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"math/rand"
)

const (
	title = "title"
	upper = "upper"
	lower = "lower"
)

func Adjective() string {
	return adjectives[rand.Intn(len(adjectives))]
}

func Name() string {
	return nouns[rand.Intn(len(nouns))]
}

func DefaultCase() string {
	return lower
}

func UpperCase() string {
	return upper
}

func LowerCase() string {
	return lower
}

func TitleCase() string {
	return title
}

// Generate generates a random name
// Takes a number of words and a separator
// If a single word is requested, it will return a name
// If more than one word is requested, it will return an adjective + name
func Generate(words int, sep string, tf string) string {
	var rs string
	if words == 1 {
		rs = Name()
	}
	rs = Adjective() + sep + Name()

	switch tf {
	case "": // if no case is specified, use lowercase as default
		c := cases.Lower(language.English)
		rs = c.String(rs)
	case upper:
		c := cases.Upper(language.English)
		rs = c.String(rs)
	case lower:
		c := cases.Lower(language.English)
		rs = c.String(rs)
	case title:
		c := cases.Title(language.English)
		rs = c.String(rs)
	default: // defaults to lowercase
		c := cases.Lower(language.English)
		rs = c.String(rs)
	}
	return rs
}
