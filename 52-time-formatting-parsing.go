//https://gobyexample.com/time-formatting-parsing
package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println

	t := time.Now()

	p("01-", t.Format(time.RFC3339))

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p("02-", t1)

	p("03-", t.Format("3:04PM"))
	p("04-", t.Format("Mon Jan _2 15:04:05 2006"))
	p("05-", t.Format("2006-01-02T15:04:05.999999-07:00"))
	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM")
	p("06-", t2)

	fmt.Printf("07- %d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM")
	p("08-", e)
}
