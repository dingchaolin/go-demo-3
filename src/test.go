package main

import "fmt"

func add( a int, b int)int{
	var sum int
	sum = a + b
	return sum
}

func main(){
	var c int
	c = add(100, 300)
	fmt.Print("你好世界", c)
}