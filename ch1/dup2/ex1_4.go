// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	recordFiles := make(map[string]string) // 记录行出现的文件
	files := os.Args[1:]
	for _, arg := range files {
		f, err := os.Open(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
			continue
		}
		countLines(f, counts, arg, recordFiles)
		f.Close()
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%-2d  %-90s  %s\n", n, line, stat(recordFiles[line]))
		}
	}
	//stat("f1.txt,f2.txt,f2.txt,f3.txt,")
}

// 记录行出现的次数及文件
func countLines(f *os.File, counts map[string]int, filename string, recordFiles map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
	    line := input.Text()
		counts[line]++
		recordFiles[line] += filename + ","
	}
	// NOTE: ignoring potential errors from input.Err()
}

// 对行出现的文件进行统计
func stat(filenames string) string {
	res := ""
	//fmt.Printf("%s\n", filenames)
	statFilename := make(map[string]int)
	for _, filename := range strings.Split(strings.Trim(filenames, ","), ",") {
		//fmt.Printf("%s\n", filename)
		statFilename[filename]++
	}
	for f, n := range statFilename {
		//fmt.Printf("%s(%d) ", f, n)
		res += fmt.Sprintf("%s(%d) ", f, n)
	}
	//fmt.Printf("\n")
	return res
}


//!-

