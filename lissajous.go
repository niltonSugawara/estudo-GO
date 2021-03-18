// Lissajous gera animacoes GIF de figuras de Lissajous aleatorias
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)
var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // primeira cor da paleta
	blackIndex = 1 // proxiam cor da paleta
)
func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}
func lissajous(out io.Writer) {
	const (
		cycles = 5 // numero de revolucoes completas do oscilador x
		res = 0.0001 // resolucao angular
    size = 100 // canvas de imagem cobre de [-size..+size]
		nframes = 64 // numero de quadros da animacao
		delay = 8 // tempo entre quadros  em unidades de 10ms
	)
	freq := rand.Float64() * 3.0 // frequencua relativa do oscilador
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // diferenca de fase
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTA: ignorando erros de codificacao
}
