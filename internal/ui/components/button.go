package components

import (
	"downmann/internal/configloader"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type IconButton struct {
	Theme  *material.Theme
	Button *widget.Clickable
	Icon   *widget.Icon
	Label  string
}

func (b *IconButton) Layout(gtx layout.Context) layout.Dimensions {
	return material.ButtonLayout(b.Theme, b.Button).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
		return layout.UniformInset(configloader.Margin).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
			container := layout.Flex{
				Axis:      layout.Horizontal,
				Alignment: layout.Middle,
			}

			icon := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return b.Icon.Layout(gtx, b.Theme.ContrastFg)
			})
			text := layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				l := material.Body1(b.Theme, b.Label)
				l.Color = b.Theme.ContrastFg

				return l.Layout(gtx)
			})

			return container.Layout(gtx, icon, text)
		})
	})
}
