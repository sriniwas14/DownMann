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

func New(theme *material.Theme, items []*download.Download) ListView {
	return ListView{
		items: items,
	}
}

func (lv *ListView) SetOnSelect(cb func(selected string)) {
	lv.onSelect = cb
}

func (lv *ListView) Layout(gtx layout.Context) layout.Dimensions {
	return layout.Flex{
		Axis: layout.Vertical,
	}.Layout(gtx, layout.Rigid(layout.Spacer{Width: 10}.Layout))
}
