package main

import (
	"encoding/json"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
	"sync"
)

var mu sync.Mutex
var count int

func main() {

	http.HandleFunc("/", handlerRicher)
	http.HandleFunc("/hello/", handerForHelloWorld)
	http.HandleFunc("/count", handlerCounter)
	http.HandleFunc("/favicon.ico", handlerFavicon)
	http.HandleFunc("/animation", handlerLassajous)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, req *http.Request) {
	addCount()
	fmt.Fprintf(w, "URL Path = %q\n", req.URL.Path)
	fmt.Printf("URL Path = %q\n", req.URL.Path)
}

func handerForHelloWorld(w http.ResponseWriter, req *http.Request) {
	addCount()
	fmt.Fprintf(w, "hello world at %s", req.URL.Path)
	fmt.Printf("hello world at %s\n", req.URL.Path)
}

func handlerCounter(w http.ResponseWriter, rew *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "count is %d", count)
	mu.Unlock()
}

func handlerFavicon(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("adress:%s asks for icon\n", req.RemoteAddr)
}

func handlerRicher(w http.ResponseWriter, req *http.Request) {
	addCount()

	fmt.Fprintf(w, "%s %s %s\n", req.Method, req.URL, req.Proto)
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", req.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", req.RemoteAddr)
	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}

	for k, v := range req.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}

	decoder := json.NewDecoder(req.Body)
	var reqBody string
	decoder.Decode(&reqBody)
	fmt.Fprintf(w, "Body = %q\n", reqBody)
}

func handlerLassajous(w http.ResponseWriter, req *http.Request) {
	addCount()
	loopCount := 5

	if err := req.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range req.Form {
		if k == "circles" {
			log.Printf("%q = %q", k, v)
			loopCount, _ = strconv.Atoi(v[0])
		}
	}
	log.Printf("Loopcount = %d", loopCount)
	lissajous(w, loopCount)
}

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

// lissajous execise 1.12
func lissajous(out io.Writer, count int) {
	const (
		cicles  = 5    //	number of complete oscillator revolutons in x axis
		res     = 0.01 // resolution in angular
		size    = 100  // image size, [-size...size]
		nframes = 64   // number of frames in a animation
		delay   = 8    // delay between images in an animation, unit 10ms
	)

	freq := 0.0
	anim := gif.GIF{LoopCount: count}
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

func addCount() {
	mu.Lock()
	count++
	mu.Unlock()
}
