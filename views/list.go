package views

import "github.com/xxxmailk/cera/view"

type ListView struct {
	view.View
}

func (l *ListView) Get() {
	l.Tpl = "list"
	l.Data["a"] = "b"
}
