package main

import "fmt"

type data struct {
	name string
}

func main() {
	m := map[string]data{"x": {"one"}}
	r := m["x"]
	r.name = "two"
	m["x"] = r
	fmt.Printf("%v\n", m)  //prints: map[x:{two}]
	fmt.Printf("%+v\n", m) //prints: map[x:{name:two}]
	fmt.Println(m)         //prints: map[x:{two}]

}
