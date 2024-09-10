# Namingo

## Usage

Namingo is a [package](https://golang.org/doc/code#ImportingRemote), so all you need to do is import it into your code:
```go
package main

import (
	"fmt"
	"github.com/sanchitrk/namingo"
)

func main() {
    fmt.Println(namingo.Generate(1, " ", namingo.TitleCase()))
    fmt.Println(namingo.Generate(2, " ", namingo.LowerCase()))
}
```
