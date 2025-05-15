package toolbar

import (
	"downmann/internal/configloader"
	"downmann/internal/ui/components"
	"fmt"

	"gioui.org/layout"
	"gioui.org/text"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
	"golang.org/x/exp/shiny/materialdesign/icons"
)

type Toolbar struct {
	theme          *material.Theme
	onSearch       func(value string)
	onAction       func(value string)
	layout         *layout.Flex
	settingsButton widget.Clickable
	newButton      widget.Clickable
	startButton    *widget.Clickable
	searchBar      widget.Editor
}

func New(theme *material.Theme) Toolbar {

	tb := Toolbar{
		theme:       theme,
		startButton: new(widget.Clickable),
		searchBar: widget.Editor{
			SingleLine: true,
			LineHeight: 50,
		},
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
	return t.layout.Layout(gtx, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
		search := material.Editor(t.theme, &t.searchBar, "Search")

		search.Editor.Alignment = text.Middle
		in := layout.Inset{
			Right: configloader.Margin,
		}

		newIcon, _ := widget.NewIcon(icons.ContentAdd)
		settingsIcon, _ := widget.NewIcon(icons.ActionSettings)
		startIcon, _ := widget.NewIcon(icons.AVPlayArrow)

		return layout.Flex{
			Axis: layout.Horizontal,
		}.Layout(gtx,
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return in.Layout(gtx, (&components.IconButton{Theme: t.theme, Icon: newIcon, Label: "", Button: &t.newButton}).Layout)
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				btn := components.IconButton{Theme: t.theme, Icon: startIcon, Label: "", Button: t.startButton}
				if t.startButton.Clicked(gtx) {
					fmt.Println("Clicked!")
				}
				return in.Layout(gtx, (&btn).Layout)
			}),
			layout.Flexed(0.5, func(gtx layout.Context) layout.Dimensions {
				return layout.UniformInset(unit.Dp(10)).Layout(gtx, func(gtx layout.Context) layout.Dimensions {
					return search.Layout(gtx)
				})
			}),
			layout.Rigid(func(gtx layout.Context) layout.Dimensions {
				return in.Layout(gtx, (&components.IconButton{Theme: t.theme, Icon: settingsIcon, Label: "", Button: &t.settingsButton}).Layout)
			}),
		)
	},
	))
}
