//https://gobyexample.com/json
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type response1 struct {
	Page   int
	Fruits []string
}

type response2 struct {
	Page       int      `json:"page"`
	Fruits     []string `json:"fruits"`
	NotVisible string   `json:"-"`
}

func main() {

	bolB, _ := json.Marshal(true)
	fmt.Print("01-")
	fmt.Println(string(bolB))

	intB, _ := json.Marshal(1)
	fmt.Print("02-")
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Print("03-")
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Print("04-")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Print("05-")
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple": 5, "lettuce": 7}
	mapB, _ := json.Marshal(mapD)
	fmt.Print("06-")
	fmt.Println(string(mapB))

	res1D := &response1{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Print("07-")
	fmt.Println(string(res1B))

	res2D := &response2{
		Page:   1,
		Fruits: []string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Print("08-")
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)

	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		panic(err)
	}
	fmt.Print("09-")
	fmt.Println(dat)

	num := dat["num"].(float64)
	fmt.Print("10-")
	fmt.Println(num)

	strs := dat["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Print("11-")
	fmt.Println(str1)

	str := `{"page": 1, "fruits": ["apple", "peach"]}`
	res := response2{}
	json.Unmarshal([]byte(str), &res)
	fmt.Print("12-")
	fmt.Println(res)
	fmt.Print("13-")
	fmt.Println(res.Fruits[0])

	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"apple": 5, "lettuce": 7}
	fmt.Print("14-")
	enc.Encode(d)
}
