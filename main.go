package main

import (
	"bytes"
	"image"
	"image/draw"
	"image/png"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"golang.org/x/exp/shiny/iconvg"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

func main() {
	ui := app.New()
	win := ui.NewWindow("main")
	win.Resize(fyne.NewSize(400, 400))

	win.SetContent(container.NewVBox(getIcon()))
	win.ShowAndRun()
}

func getIcon() fyne.CanvasObject {
	var rast iconvg.Rasterizer
	m := image.NewRGBA(image.Rect(0, 0, 50, 50))
	rast.SetDstImage(m, m.Bounds(), draw.Src)
	err := iconvg.Decode(&rast, icons.AVAirplay, nil)
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(nil)
	err = png.Encode(buf, m)
	if err != nil {
		panic(err)
	}
	return widget.NewIcon(fyne.NewStaticResource("icon.png", buf.Bytes()))
}

func NewImg() *canvas.Image {
	dst := image.NewAlpha(image.Rect(0, 0, 50, 50))
	var r iconvg.Rasterizer
	r.SetDstImage(dst, dst.Bounds(), draw.Src)

	opt := iconvg.DecodeOptions{
		Palette: &iconvg.DefaultPalette,
	}
	if err := iconvg.Decode(&r, icons.AVAirplay, &opt); err != nil {
		panic(err)
	}
	return canvas.NewImageFromImage(dst)
}
