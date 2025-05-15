package categories

import (
	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

type Categories struct {
	theme      *material.Theme
	categories []string
	onSelect   func(selected string)
	buttons    []*widget.Clickable
	Selected   string
}

func New(theme *material.Theme, categories []string, onSelect func(string)) Categories {
	return Categories{
		theme:      theme,
		categories: categories,
		buttons:    []*widget.Clickable{},
		Selected:   categories[0],
		onSelect:   onSelect,
	}
}

func (c *Categories) SetOnSelect(cb func(selected string)) {
	c.onSelect = cb
}

func (c *Categories) Layout(gtx layout.Context) layout.Dimensions {
	options := []layout.FlexChild{}
	in := layout.Inset{Top: 3}

	if len(c.buttons) == 0 {
		for len(c.buttons) < len(c.categories) {
			c.buttons = append(c.buttons, &widget.Clickable{})
		}
	}

	for i, cat := range c.categories {
		b := c.buttons[i]
		label := material.Button(c.theme, b, cat)

		if b.Clicked(gtx) {
			c.onSelect(cat)
			c.Selected = cat
		}

		if c.Selected == cat {
			label.Background = c.theme.ContrastBg
			label.Color = c.theme.ContrastFg
		} else {
			label.Background = c.theme.Bg
			label.Color = c.theme.Fg
		}

		options = append(options, layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return in.Layout(gtx, func(c layout.Context) layout.Dimensions {
				return label.Layout(gtx)
			})
		}))
	}

	return layout.Flex{
		Axis: layout.Horizontal,
	}.Layout(gtx, options...)
}
