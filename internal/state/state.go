package state

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func _() {
	// _ = desktop.Hoverable
}

var _ fyne.Widget = (*BinState)(nil)
var _ fyne.WidgetRenderer = (*binStateRenderer)(nil)

type BinState struct {
	widget.BaseWidget

	active   *widget.Icon
	inactive *widget.Icon
	state    binding.Bool
}

func NewBinState(act, inact *widget.Icon) *BinState {
	bs := &BinState{active: act, inactive: inact, state: binding.NewBool()}
	bs.BaseWidget.ExtendBaseWidget(bs)
	return bs
}

func NewBinStateWithData(act, inact *widget.Icon, data binding.Bool) *BinState {
	bs := &BinState{active: act, inactive: inact, state: data}
	bs.BaseWidget.ExtendBaseWidget(bs)
	return bs
}

func (bs *BinState) CreateRenderer() fyne.WidgetRenderer {
	return &binStateRenderer{ref: bs}
}

func (bs *BinState) SetState(active bool) {
	bs.state.Set(active)
	bs.Refresh()
}

func (bs *BinState) State() bool {
	return bs.getState()
}

func (bs *BinState) getState() bool {
	s, _ := bs.state.Get()
	return s
}

func (bs *BinState) MinSize() fyne.Size {
	return fyne.Size{Width: 50, Height: 50}
}

type binStateRenderer struct {
	ref *BinState
}

func (r *binStateRenderer) Destroy() {}

func (r *binStateRenderer) Layout(size fyne.Size) {
	if r.ref.getState() {
		r.ref.active.Resize(size)
	} else {
		r.ref.inactive.Resize(size)
	}
}

func (r *binStateRenderer) MinSize() fyne.Size {
	return r.ref.MinSize()
}

func (r *binStateRenderer) Objects() []fyne.CanvasObject {
	if r.ref.getState() {
		return []fyne.CanvasObject{r.ref.active}
	}
	return []fyne.CanvasObject{r.ref.inactive}
}

func (r *binStateRenderer) Refresh() {
	canvas.Refresh(r.ref)
}
