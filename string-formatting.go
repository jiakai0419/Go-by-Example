package main

import "fmt"
import "os"

type point struct {
	x, y int
}

var P = fmt.Printf

func main() {
	p := point{1, 2}
	P("%v\n", p)
	P("%+v\n", p)
	P("%#v\n", p)
	P("%T\n", p)

	P("%t\n", true)

	P("%d\n", 123)
	P("%b\n", 14)
	P("%c\n", 33)
	P("%x\n", 456)

	P("%f\n", 78.9)
	P("%e\n", 123400000.0)
	P("%E\n", 123400000.0)

	P("%s\n", "\"string\"")
	P("%q\n", "\"string\"")
	P("%x\n", "hex this")

	P("%p\n", &p)

	P("|%6d|%6d|\n", 12, 345)
	P("|%6.2f|%6.2f|\n", 1.2, 3.45)
	P("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	P("|%6s|%6s|\n", "foo", "b")
	P("|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("a %s\n", "string")
	P(s)

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
