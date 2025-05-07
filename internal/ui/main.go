package ui

import (
	"downmann/internal/ui/categories"
	"downmann/internal/ui/listview"
	"downmann/internal/ui/toolbar"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Window struct {
	ops        *op.Ops
	toolbar    *toolbar.Toolbar
	categories *categories.Categories
	listview   *listview.ListView
}

func New() *Window {
	w := new(app.Window)
	w.Option(app.Title("Downmann"))
	w.Option(app.Size(unit.Dp(800), unit.Dp(600)))

	theme := material.NewTheme()
	var ops op.Ops

	// listen for events in the window
	for {
		evt := w.Event()

		switch evtType := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, evtType)

			layout.Flex{
				Axis: layout.Vertical,
			}.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				tb := toolbar.New(theme, &gtx)
				return tb.Layout(gtx)
			}))

			// cg := categories.New(theme, &gtx, configloader.ItemCategories)
			// lv := listview.New(theme, &gtx, nil)

			evtType.Frame(gtx.Ops)

		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}
