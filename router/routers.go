package router

import (
	"github.com/xxxmailk/cera/router"
	"y-blog/views"
)

func Routers() *router.Router {
	r := router.New()
	r.ANY("/", &views.IndexView{})
	r.ANY("/list", &views.ListView{})
	r.ServeFiles("/static/{filepath:*}", "/root/go/src/y-blog/static")
	return r
}
