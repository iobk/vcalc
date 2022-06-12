package main

import (
	"fmt"
	"vcalc/add"
	"vcalc/add/mul"
	"vcalc/sub"
)

func main() {
	fmt.Println("vcalc main routine ------")
	fmt.Println("vcalc/add funtion add.Add(10,20)=", add.Add(10, 20))
	fmt.Println("vcalc/add funtion add.Add1(10,-20)=", add.Add(10, -20))
	fmt.Println("vcalc/sub funtion sub.Sub(10,20)=", sub.Sub(10, 20))
	fmt.Println("vcalc/sub funtion sub.Sub1(10,-20)=", sub.Sub(10, -20))
	fmt.Println("vcalc/add/mul funtion mul.Mul(10,20)=", mul.Mul(10, 20))
	fmt.Println("vcalc/add/mul funtion mul.Mul1(10,-20)=", mul.Mul1(10, -20))

}
