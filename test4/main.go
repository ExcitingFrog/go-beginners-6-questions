package main

//#include<stdio.h>
import "C"
import "fmt"

func main() {
	set := newmyset()
	cgo_insert(set, 2)
	cgo_insert(set, 1)
	fmt.Println(cgo_size(set))
	fmt.Println(cgo_find(set, 1))
	fmt.Println(cgo_find(set, 3))
	cgo_erase(set, 1)
	fmt.Println(cgo_size(set))

	//fmt.Println(set)

}
