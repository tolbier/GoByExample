//https://gobyexample.com/channel-buffering
//also https://tour.golang.org/concurrency/2
package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
