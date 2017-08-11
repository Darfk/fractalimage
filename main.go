package main

import (
	"./fractalimage"
	"log"
	"fmt"
	"math/cmplx"
)

func main() {

	// Change these variables 

	// width of image
	w := 1920
	h := 1080

	// inverse zoom factor
	scale := float64(0.001)
	
	// point to zoom in on
	point := complex(0.34,0.068)

	// OK probably don't change these

	aspect := float64(w) / float64(h)

	ds := complex(scale * aspect, scale)

	tl := point - ds / 2
	br := point + ds / 2

	img := fractalimage.NewFractalImage(tl, br, w, h)
	dx, dy := img.Dx(), img.Dy()

	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			var c complex128 = img.ImagCoordsFromPixelCoords(x, y)

			var p complex128
			var d int
			for {
				m := cmplx.Abs(p)
				if m > 1e50 {
					break
				}
				d++
				if d >= 100 {
					break
				}
				p = cmplx.Pow(p, 2) + c
			}
			img.Set(x, y, d)
		}
	}

	err := img.ToFile(fmt.Sprintf("mandlebrot-%G-%G-%dx%d.png", tl, br, img.Dx(), img.Dy()))
	if err != nil {
		log.Fatal(err)
	}

}
