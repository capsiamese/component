package state

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
	"image/png"

	"fyne.io/fyne/v2"
	"golang.org/x/exp/shiny/iconvg"
)

/*
golang.org/x/exp/shiny/iconvg
golang.org/x/exp/shiny/materialdesign/colornames
golang.org/x/exp/shiny/materialdesign/icons
https://fonts.google.com/icons
*/

func SingleColor(c color.RGBA) *iconvg.Palette {
	return &iconvg.Palette{c}
}

func MustMaterialIcon(src []byte, sz fyne.Size, palette *iconvg.Palette) fyne.Resource {
	icon, err := NewMaterialIcon(src, sz, palette)
	if err != nil {
		panic(err)
	}
	return icon
}

func NewMaterialIcon(src []byte, sz fyne.Size, palette *iconvg.Palette) (fyne.Resource, error) {
	target := image.NewRGBA(image.Rect(0, 0, int(sz.Width), int(sz.Height)))
	dst := iconvg.Rasterizer{}
	dst.SetDstImage(target, target.Bounds(), draw.Src)
	var opt *iconvg.DecodeOptions
	if palette != nil {
		opt = &iconvg.DecodeOptions{Palette: palette}
	}
	if err := iconvg.Decode(&dst, src, opt); err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(nil)
	if err := png.Encode(buf, target); err != nil {
		return nil, err
	}
	return fyne.NewStaticResource("icon.png", buf.Bytes()), nil
}

func NewSVGIcon(src []byte, sz fyne.Size) fyne.Resource {
	return fyne.NewStaticResource("icon.svg", src)
}
