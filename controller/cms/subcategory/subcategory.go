package subcategory

import (
	"net/http"
	"tech/lib/flight"
	"tech/middleware/acl"

	"github.com/blue-jay/core/router"
)

var (
	uri = "/cms/subcategory"
)

func Load() {
	c := router.Chain(acl.DisallowAnon)
	// router.Get(uri, Index, c...)
	router.Get(uri+"/create", Create, c...)
	// router.Post(uri+"/create", Store, c...)
	// router.Get(uri+"/view/:id", Show, c...)
}

func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("cms/subcategory/create")
	c.Repopulate(v.Vars, "name_sub_category")
	v.Render(w, r)
}
