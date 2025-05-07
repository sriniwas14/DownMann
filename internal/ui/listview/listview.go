package listview

import (
	"downmann/internal/download"

	"gioui.org/layout"
	"gioui.org/widget/material"
)

type ListView struct {
	items    []*download.Download
	onSelect func(selected string)
}

func New(theme *material.Theme, ctx *layout.Context, items []*download.Download) ListView {
	return ListView{
		items: items,
	}
}

func (lv *ListView) SetOnSelect(cb func(selected string)) {
	lv.onSelect = cb
}
