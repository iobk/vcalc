package main

import (
	"fmt"
	"vcalc/add"
)

func main() {
	fmt.Println("vcalc main routine ------")
	fmt.Println("vcalc/add funtion add.Add(10,20)=", add.Add(10, 20))
	fmt.Println("vcalc/add funtion add.Add1(10,-20)=", add.Add(10, -20))

}
