//https://gobyexample.com/string-formatting
package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("01-%v\n", p)

	fmt.Printf("02-%+v\n", p)

	fmt.Printf("03-%#v\n", p)

	fmt.Printf("04-%T\n", p)

	fmt.Printf("05-%t\n", true)

	fmt.Printf("06-%d\n", 123)

	fmt.Printf("07-%b\n", 14)

	fmt.Printf("08-%c\n", 33)

	fmt.Printf("09-%x\n", 456)

	fmt.Printf("10-%f\n", 78.9)

	fmt.Printf("11-%e\n", 123400000.0)
	fmt.Printf("12-%E\n", 123400000.0)

	fmt.Printf("13-%s\n", "\"string\"")

	fmt.Printf("14-%q\n", "\"string\"")

	fmt.Printf("15-%x\n", "hex this")

	fmt.Printf("16-%p\n", &p)

	fmt.Printf("17-|%6d|%6d|\n", 12, 345)

	fmt.Printf("18-|%6.2f|%6.2f|\n", 1.2, 3.45)

	fmt.Printf("19-|%-6.2f|%-6.2f|\n", 1.2, 3.45)

	fmt.Printf("20-|%6s|%6s|\n", "foo", "b")

	fmt.Printf("21-|%-6s|%-6s|\n", "foo", "b")

	s := fmt.Sprintf("22-a %s", "string")
	fmt.Println(s)

	fmt.Fprintf(os.Stderr, "23-an %s\n", "error")
}
