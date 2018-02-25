package main

import (
	"fmt"
	"stringutil"
	//"bytes"
)

func main() {
	var x, y, z int
	var str string
	//var buffer bytes.Buffer
	z = 3
	y = 4
	x = y + z
	fmt.Printf("%b \n", x)
	str = stringutil.Reverse("1TPreffuB")
	fmt.Println(str)
}