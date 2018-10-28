package main

import (
	"fmt"
	"time"
)

func main() {


	for i := 0; i < 10; i++ {

		go f(i)
	}

	fmt.Println("End")

	var in string

	fmt.Scanln(&in)


}

func f(in int) {

	for i := 0; i < 10; i++ {

		fmt.Println(i, "=", in)

		amt := time.Duration(250)
		time.Sleep(time.Millisecond * amt)
	}
}
