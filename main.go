package main

import (
	"image/color"
	"log"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	firePalette = []color.RGBA{
		{R: 7, G: 7, B: 7, A: 255},       //  0
		{R: 31, G: 7, B: 7, A: 255},      //  1
		{R: 47, G: 15, B: 7, A: 255},     //  2
		{R: 71, G: 15, B: 7, A: 255},     //  3
		{R: 87, G: 23, B: 7, A: 255},     //  4
		{R: 103, G: 31, B: 7, A: 255},    //  5
		{R: 119, G: 31, B: 7, A: 255},    //  6
		{R: 143, G: 39, B: 7, A: 255},    //  7
		{R: 159, G: 47, B: 7, A: 255},    //  8
		{R: 175, G: 63, B: 7, A: 255},    //  9
		{R: 191, G: 71, B: 7, A: 255},    // 10
		{R: 199, G: 71, B: 7, A: 255},    // 11
		{R: 223, G: 79, B: 7, A: 255},    // 12
		{R: 223, G: 87, B: 7, A: 255},    // 13
		{R: 223, G: 87, B: 7, A: 255},    // 14
		{R: 215, G: 95, B: 7, A: 255},    // 15
		{R: 215, G: 95, B: 7, A: 255},    // 16
		{R: 215, G: 103, B: 15, A: 255},  // 17
		{R: 207, G: 111, B: 15, A: 255},  // 18
		{R: 207, G: 119, B: 15, A: 255},  // 19
		{R: 207, G: 127, B: 15, A: 255},  // 20
		{R: 207, G: 135, B: 23, A: 255},  // 21
		{R: 199, G: 135, B: 23, A: 255},  // 22
		{R: 199, G: 143, B: 23, A: 255},  // 23
		{R: 199, G: 151, B: 31, A: 255},  // 24
		{R: 191, G: 159, B: 31, A: 255},  // 25
		{R: 191, G: 159, B: 31, A: 255},  // 26
		{R: 191, G: 167, B: 39, A: 255},  // 27
		{R: 191, G: 167, B: 39, A: 255},  // 28
		{R: 191, G: 175, B: 47, A: 255},  // 29
		{R: 183, G: 175, B: 47, A: 255},  // 30
		{R: 183, G: 183, B: 47, A: 255},  // 31
		{R: 183, G: 183, B: 55, A: 255},  // 32
		{R: 207, G: 207, B: 111, A: 255}, // 33
		{R: 223, G: 223, B: 159, A: 255}, // 34
		{R: 239, G: 239, B: 199, A: 255}, // 35
		{R: 255, G: 255, B: 255, A: 255}, // 36
	}

	screenWidth, screenHeight = ebiten.Monitor().Size()
	screenSize                = screenWidth * screenHeight
)

type fire struct {
	pixels  []byte
	indices []byte
}

func (f *fire) renderFire() {
	for i, v := range f.indices {
		p := firePalette[v]
		f.pixels[i*4] = p.R
		f.pixels[i*4+1] = p.G
		f.pixels[i*4+2] = p.B

		if p.R <= 7 && p.G <= 7 && p.B <= 7 {
			f.pixels[i*4+3] = 0
			continue
		}

		// semi-transparent
		f.pixels[i*4+3] = p.A / 2
	}
}

func (f *fire) updateFireIntensityPerPixel(currentPixelIndex int) {
	below := currentPixelIndex + screenWidth
	if below >= screenSize {
		return
	}

	d := rand.IntN(2)
	newI := int(f.indices[below]) - d
	if newI < 0 {
		newI = 0
	}

	if currentPixelIndex-d < 0 {
		return
	}
	f.indices[currentPixelIndex-d] = byte(newI)
}

func (f *fire) updateFirePixels() {
	for i := 0; i < screenWidth; i++ {
		for j := 0; j < screenHeight; j++ {
			idx := i + (screenWidth * j)
			f.updateFireIntensityPerPixel(idx)
		}
	}
}

func New() *fire {
	indices := make([]byte, screenSize)

	for i := screenSize - screenWidth; i < screenSize; i++ {
		indices[i] = 36
	}

	return &fire{pixels: make([]byte, screenSize*4),

		indices: indices,
	}
}

func (f *fire) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (f *fire) Update() error {
	f.updateFirePixels()
	return nil
}

func (f *fire) Draw(screen *ebiten.Image) {
	f.renderFire()
	screen.WritePixels(f.pixels)
}

func (f *fire) Run() {
	const name = "doom-fire"

	screenWidth, screenHeight = ebiten.Monitor().Size()
	screenSize = (screenWidth * screenHeight)

	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetRunnableOnUnfocused(true)
	ebiten.SetScreenClearedEveryFrame(false)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowDecorated(false)
	ebiten.SetWindowFloating(true)
	ebiten.SetWindowMousePassthrough(true)
	ebiten.SetWindowPosition(0, 0)
	ebiten.SetWindowSize(screenWidth, screenHeight)

	ebiten.SetWindowTitle(name)

	err := ebiten.RunGameWithOptions(f, &ebiten.RunGameOptions{
		InitUnfocused:     true,
		ScreenTransparent: true,
		SkipTaskbar:       true,
		X11ClassName:      name,
		X11InstanceName:   name,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	f := New()
	f.Run()
}
