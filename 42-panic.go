//https://gobyexample.com/panic
package main

import "os"

func main() {

	panic("a problem")

	_, err := os.Create("/tmp/file3")
	if err != nil {
		panic(err)
	}
}
