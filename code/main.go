package main

import "fmt"

func main() {
	a()
	//fmt.Println(a)

}

func a() {
	a := 2
	b := []string{"1", "2", "3"}
	b = append(b, "4")

	fmt.Println(a)
	fmt.Println(cap(b))
}

