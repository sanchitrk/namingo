package main

import (
	"flag"
	"fmt"
	"github.com/sanchitrk/namingo"
)

var (
	words     = flag.Int("w", 1, "Number of words in the name")
	sep       = flag.String("sep", " ", "Separator between words")
	transform = flag.String("case", "title", "Case of the generated name")
)

func main() {
	flag.Parse()
	fmt.Println(namingo.Generate(*words, *sep, *transform))
}
