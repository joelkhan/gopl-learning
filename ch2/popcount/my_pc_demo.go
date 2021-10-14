package main

import (
	"fmt"

	"gopl-learning/ch2/popcount"
)

const (
    demoNum =  52258//  1100 1100 0010 0010
)

func main() {
    fmt.Printf("sample code:\n  %d %d\n", demoNum, popcount.PopCount(demoNum))
	//fmt.Printf("%b\n", byte(demoNum>>(0*8)))
	//showByteSlice()
	fmt.Printf("Ex2_3:\n  %d %d\n", demoNum, popcount.PopCount2(demoNum))
	fmt.Printf("Ex2_4:\n  %d %d\n", demoNum, popcount.PopCount3(demoNum))
	fmt.Printf("Ex2_5:\n  %d %d\n", demoNum, popcount.PopCount4(demoNum))
}


func showByteSlice() {
	var pc [256]byte
	for i := range pc {
		fmt.Printf("%d %d\n", i, pc[i])
	}
}





