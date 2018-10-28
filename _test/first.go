package main

import "fmt"

func main() {


	//list := []int{		1,		2,		3,		4}
	//fmt.Println(list)

	x := 5

	cls(&x)

	fmt.Println(x)

}

func cls(i *int) {


	*i = *i * 10
}