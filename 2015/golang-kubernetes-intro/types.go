package main

import "fmt"

func main() {
	var name1 string
	name1 = "bob"
	fmt.Println(name1)

	name2 := "mary"
	fmt.Println(name2)

	first, last, age := "chris", "whitcombe", 30
	fmt.Printf("%s, %s (%d)\n", first, last, age)
}
