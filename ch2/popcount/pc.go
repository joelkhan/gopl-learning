// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

import "fmt"

// pc[i] is the population count of i.
var pc [256]byte

/*
 * 两个init()按顺序执行
 * 第一个init()：
 *   计算0-255范围（即一个字节表示的所有整数）内的每一个整数的
 *   二进制形式所包含的1的个数，使用“整数”作为索引，
 *   “包含的1的个数”作为值，例如：
 *     pc[0] = 0
 *     pc[1] = 1
 *     pc[2] = 1
 *     pc[3] = 2
 *     pc[4] = 1
 *   这个算法要点包括：
 *     (1) 应注意到，pc[]的初始化过程是从0逐一完成的，也就是说，当计算pc[5]时，需要的pc[2]是已知的
       (2) i/2相当于在二进制的层面上把i右移了一位，此时，对i与i/2的位数判断，
           唯一的区别就在于被移走的那一位是不是1
 *     (3) 被移走的那一位，可以通过byte(i&1)来判断
 * 第二个init()：
 *   打印pc[]的初始化结果
 */
func init() {
    fmt.Println("init pc []byte")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}


func init() {
    fmt.Println("show pc init result...")
	for i := range pc {
		fmt.Printf("%d ", pc[i])
	}
	fmt.Println()
}


// PopCount returns the population count (number of set bits) of x.
// 注意： byte(x)的结果是取x的低8位
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// Ex2_3
func PopCount2(x uint64) int {
	var sum int
	for i := 0; i <= 7; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// Ex2_4
func PopCount3(x uint64) int {
	var sum int
	for i := 0; i < 64; i++ {
	    sum += int((x>>i) & 1)
	}
	return sum
}

// Ex2_5
func PopCount4(x uint64) int {
	var sum int
	
	for x!=0 {
		sum++
		x = x & (x-1)
	}
	
	return sum
}

//!-
