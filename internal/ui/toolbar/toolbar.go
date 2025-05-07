package toolbar

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Toolbar struct {
	theme    *material.Theme
	onSearch func(value string)
	onAction func(value string)
	layout   *layout.Flex
}

func New(theme *material.Theme, ctx *layout.Context) Toolbar {
	tb := Toolbar{
		theme: theme,
	}

	tb.layout = &layout.Flex{}

	return tb
}

func (t *Toolbar) SetOnSearch(cb func(value string)) {
	t.onSearch = cb
}

func (t *Toolbar) SetOnAction(cb func(value string)) {
	t.onAction = cb
}

func (t *Toolbar) Layout(gtx layout.Context) layout.Dimensions {
	var startButton widget.Clickable

	return t.layout.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		btn := material.Button(t.theme, &startButton, "Settings")
		return btn.Layout(gtx)
	},
	))
}
