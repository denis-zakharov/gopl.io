// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

var palette = []color.Color{
	color.Black,
	color.RGBA{0x0, 0xff, 0x0, 0xff}, // green
	color.RGBA{0xff, 0x0, 0x0, 0xff}, // red
	color.RGBA{0x0, 0x0, 0xff, 0xff}, // blue
}

func main() {
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // http.ResponseWriter implements io.Writer
			if err := r.ParseForm(); err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}
			cyclesParam := r.Form.Get("cycles")
			if len(cyclesParam) == 0 {
				lissajous(w, 5)
				return
			}
			cycles, err := strconv.Atoi(cyclesParam)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			}
			lissajous(w, cycles)

		})
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	fmt.Println("<command> web")
}

func lissajous(out io.Writer, cycles int) {
	const (
		res     = 0.001 // angular resolution
		size    = 100   // image canvas covers [-size..+size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles*2)*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			idx := uint8(rand.Intn(len(palette)-1) + 1)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), idx)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

//!-main
