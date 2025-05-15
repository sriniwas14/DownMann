package ui

import (
	"downmann/internal/configloader"
	"downmann/internal/download"
	"downmann/internal/ui/categories"
	"downmann/internal/ui/listview"
	"downmann/internal/ui/toolbar"
	"fmt"
	"image/color"
	"os"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget/material"
)

type Window struct {
	ops        *op.Ops
	toolbar    toolbar.Toolbar
	categories categories.Categories
	listview   listview.ListView
}

func New() *Window {
	w := new(app.Window)
	w.Option(app.Title("Downmann"))
	w.Option(app.Size(unit.Dp(800), unit.Dp(600)))

	theme := material.NewTheme()
	theme.Bg = color.NRGBA{R: 25, G: 23, B: 36, A: 255}
	theme.Fg = color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	theme.ContrastBg = color.NRGBA{R: 234, G: 157, B: 52, A: 255}
	var ops op.Ops

	win := Window{
		ops: &ops,
	}

	win.toolbar = toolbar.New(theme)
	win.categories = categories.New(theme, configloader.ItemCategories, func(selected string) {
		fmt.Println("Selected ", selected)
	})
	win.listview = listview.New(theme, []*download.Download{})

	// listen for events in the window
	for {
		evt := w.Event()

		switch evtType := evt.(type) {
		case app.FrameEvent:
			gtx := app.NewContext(&ops, evtType)
			paint.Fill(gtx.Ops, theme.Bg)

			layout.UniformInset(configloader.Margin).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
				return layout.Flex{
					Axis: layout.Vertical,
				}.Layout(gtx,
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return win.toolbar.Layout(gtx)
					}),
					layout.Rigid(layout.Spacer{Height: 10}.Layout),
					layout.Rigid(func(gtx layout.Context) layout.Dimensions {
						return win.categories.Layout(gtx)
					}),
				)
			})

			// lv := listview.New(theme, &gtx, nil)

			// fmt.Println(theme.Bg)
			evtType.Frame(gtx.Ops)

		case app.DestroyEvent:
			os.Exit(0)
		}
	}
}
