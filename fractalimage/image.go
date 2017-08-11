package fractalimage

import (
	"image"
	"image/color"
	"image/png"
	"math"
	"os"
)

type FractalImage struct {
	img    *image.NRGBA
	w, h   int
	tl, br complex128
}

func (q *FractalImage) ImagCoordsFromPixelCoords(x, y int) complex128 {
	return complex(
		(float64(x)/float64(q.w))*(real(q.br)-real(q.tl))+real(q.tl),
		(float64(y)/float64(q.h))*(imag(q.br)-imag(q.tl))+imag(q.tl),
	)
}

func NewFractalImage(tl, br complex128, w, h int) (q *FractalImage) {
	q = &FractalImage{
		img: image.NewNRGBA(image.Rect(0, 0, w, h)),
		tl:  tl,
		br:  br,
		w:   w,
		h:   h,
	}
	return
}

func (q *FractalImage) Dx() int {
	return q.img.Rect.Dx()
}

func (q *FractalImage) Dy() int {
	return q.img.Rect.Dy()
}

func (q *FractalImage) ToFile(path string) error {
	fd, err := os.Create(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	err = png.Encode(fd, q.img)
	if err != nil {
		return err
	}
	return nil
}

func (q *FractalImage) Set(x, y, d int) {
	if(d==100) {
		q.img.Set(x, y, color.NRGBA{uint8(0),uint8(0),uint8(0), 255})
		return
	}
	r, g, b := HSL2RGB(d*3, 100, 50)
	q.img.Set(x, y, color.NRGBA{uint8(r), uint8(g), uint8(b), 255})
}

// function stolen from https://github.com/SimonWaldherr/ColorConverterGo/blob/master/ColorConverter.go
func HSL2RGB(h, s, l int) (int, int, int) {
	var hf, sf, lf, vf, minf, svf, sixf, fractf, vsfractf, rf, gf, bf float64

	hf = math.Max(math.Min(float64(int(h)), 360), 0) / 360
	sf = math.Max(math.Min(float64(int(s)), 100), 0) / 100
	lf = math.Max(math.Min(float64(int(l)), 100), 0) / 100

	if lf <= 0.5 {
		vf = lf * (1 + sf)
	} else {
		vf = lf + sf - sf - 1*sf
	}
	if vf == 0 {
		return int(0), int(0), int(0)
	}
	minf = 2*lf - vf
	svf = (vf - minf) / vf
	hf = 6 * hf
	sixf = float64(int(hf))
	fractf = hf - sixf
	vsfractf = vf * svf * fractf
	switch sixf {
	case 1:
		rf = vf - vsfractf
		gf = vf
		bf = minf
	case 2:
		rf = minf
		gf = vf
		bf = minf + vsfractf
	case 3:
		rf = minf
		gf = vf - vsfractf
		bf = vf
	case 4:
		rf = minf + vsfractf
		gf = minf
		bf = vf
	case 5:
		rf = vf
		gf = minf
		bf = vf - vsfractf
	default:
		rf = vf
		gf = minf + vsfractf
		bf = minf
	}
	return int(rf * 255), int(gf * 255), int(bf * 255)
}
