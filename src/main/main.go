package main

import (
	"fmt"
	"talker"
)

func main() {

	fmt.Println("Start")
	t := talker.Talker{"Mike", 30}
	fmt.Println(t.Talk())
}
