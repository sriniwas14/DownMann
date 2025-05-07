package categories

import (
	"gioui.org/layout"
	"gioui.org/widget/material"
)

type Categories struct {
	categories []string
	onSelect   func(selected string)
}

func New(theme *material.Theme, ctx *layout.Context, categories []string) Categories {
	return Categories{
		categories: categories,
	}
}

func (c *Categories) SetOnSelect(cb func(selected string)) {
	c.onSelect = cb
}
