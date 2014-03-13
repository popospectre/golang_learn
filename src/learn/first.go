package main

import (
	"fmt"
	//"bytes"
	//"strconv"
	"strings"
)

type Person struct {
	Name string
	Age  int
	Sex  int
}

func main() {
	fmt.Println("hello world, go")

	p1 := &Person{"polaris", 28, 0}
	fmt.Println(p1)

	fmt.Println(strings.Contains("in failure", "s g"))
	const intSize = 32 << uint(^uint(0)>>63)
	fmt.Printf("%v", intSize)
}
