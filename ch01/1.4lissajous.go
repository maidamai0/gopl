package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
)

var palette = []color.Color{
	color.RGBA{100, 100, 100, 255},
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
	color.RGBA{0, 255, 255, 255},
	color.RGBA{255, 0, 255, 255},
}

const (
	whiteIndex = 0
	blackIndex = 1
)

// execise 1.5 and 1.6 also inclued in this file
func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cicles  = 5    //	number of complete oscillator revolutons in x axis
		res     = 0.01 // resolution in angular
		size    = 100  // image size, [-size...size]
		nframes = 64   // number of frames in a animation
		delay   = 8    // delay between images in an animation, unit 10ms
	)

	freq := 0.0
	anim := gif.GIF{LoopCount: 100}
	phase := 0.0
	var index uint8

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < 2*cicles*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), index)
		} // loop

		phase += 0.1
		freq += 0.01
		index++
		if index == 7 {
			index = 0
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	} // loop

	gif.EncodeAll(out, &anim)
}
