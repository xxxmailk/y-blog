package views

import "github.com/xxxmailk/cera/view"

type IndexView struct {
	view.View
}

func (i *IndexView) Get() {
	i.Tpl = "index"
	i.Data["menu"] = "index"
}
