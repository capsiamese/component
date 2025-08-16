package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/capsiamese/component/internal/state"
	"golang.org/x/exp/shiny/materialdesign/icons"
	"golang.org/x/image/colornames"
)

func main() {
	sz := fyne.NewSize(50, 50)
	inactIcon := state.MustMaterialIcon(icons.ToggleRadioButtonChecked, sz, state.SingleColor(colornames.Red))
	actIcon := state.MustMaterialIcon(icons.ToggleRadioButtonChecked, sz, state.SingleColor(colornames.Green))

	ui := app.New()
	win := ui.NewWindow("main")
	win.Resize(fyne.NewSize(400, 400))

	bs := state.NewBinState(widget.NewIcon(actIcon), widget.NewIcon(inactIcon))
	btn := widget.NewButton("Toggle", func() {
		bs.SetState(!bs.State())
	})

	win.SetContent(container.NewVBox(bs, btn))
	win.ShowAndRun()
}
