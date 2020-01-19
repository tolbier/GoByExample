//https://gobyexample.com/regular-expressions
package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func main() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println("01-", r.MatchString("peach"))

	fmt.Println("02-", r.FindString("peach punch"))

	fmt.Println("03-", r.FindStringIndex("peach punch"))

	fmt.Println("04-", r.FindStringSubmatch("peach punch"))

	fmt.Println("05-", r.FindStringSubmatchIndex("peach punch"))

	fmt.Println("06-", r.FindAllString("peach punch pinch", -1))

	fmt.Println("07-", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1))

	fmt.Println("08-", r.FindAllString("peach punch pinch", 2))

	fmt.Println("09-", r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("10-", r)

	fmt.Println("11-", r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println("12-", string(out))
}
