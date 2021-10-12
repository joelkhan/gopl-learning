// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Ex1_8
		if !strings.HasPrefix(url, "http") {
			url = "http://" + url
		}
		fmt.Printf("\nLet's fetch >>>%s<<<\n", url)
		resp, err := http.Get(url) // resp响应结构体
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		b, err := ioutil.ReadAll(resp.Body) // Body是resp响应结构体中一个服务端响应的可读取数据流
		// Ex1_9
		fmt.Printf("HTTP Status Code: %s\n", resp.Status)
		resp.Body.Close() // 关闭数据流
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", b)
	}
}

//!-
