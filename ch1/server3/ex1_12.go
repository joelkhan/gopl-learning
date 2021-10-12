// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"strconv"
)

//!-main
// Packages not needed by version in book.
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{color.White, color.Black} // slice

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			cycles := 5.0 // 默认值
			// 从URL中解析cycles
			if err := r.ParseForm(); err != nil {
				log.Print(err)
			}
        	for k, v := range r.Form {
        		//fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
				if k == "cycles" {
					cValue, err := strconv.Atoi(v[0]) // 字符串转整数
        			if err != nil {
        				log.Print(err)
        			}
					cycles = float64(cValue) // 整数转浮点
					break
				}
        	}
			lissajous(w, cycles)
		}
		
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout, 5.0)
}

// 
func lissajous(out io.Writer, cycles float64) {
	const (  // 定义常量
//		cycles  = 20     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes} // anim是一个结构体，有64帧
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1 // y方向的相位随每帧增加
		anim.Delay = append(anim.Delay, delay) // anim的每帧延迟8*10ms
		anim.Image = append(anim.Image, img) // anim的每帧是一幅img
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}


//!-main

