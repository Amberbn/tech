package category

import (
	"fmt"
	"net/http"
	"tech/lib/flight"
	"tech/middleware/acl"
	"tech/model/category"

	"github.com/blue-jay/core/router"
)

var (
	uri = "/cms/category"
)

func Load() {
	c := router.Chain(acl.DisallowAnon)
	router.Get(uri, Index, c...)
	router.Get(uri+"/create", Create, c...)
	router.Post(uri+"/create", Store, c...)
	router.Get(uri+"/view/:id", Show, c...)
}

func Index(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	items, _, err := category.All(c.DB)
	if err != nil {
		c.FlashErrorGeneric(err)
		items = []category.Item{}
	}

	v := c.View.New("cms/category/index")
	v.Vars["items"] = items
	v.Render(w, r)
}

func Create(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)
	v := c.View.New("cms/category/create")
	c.Repopulate(v.Vars, "name_category")
	v.Render(w, r)
}

func Store(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	h := c.FormValid("name_category")
	fmt.Println("name_category", h)
	if !c.FormValid("name_category") {
		Create(w, r)
		return
	}

	_, err := category.Create(c.DB, r.FormValue("name_category"))
	if err != nil {
		c.FlashErrorGeneric(err)
		Create(w, r)
		return
	}

	c.FlashSuccess("success add category")
	c.Redirect(uri)
}

func Show(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := category.ById(c.DB, c.Param("id"))
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("cms/category/detail")
	v.Vars["item"] = item
	v.Render(w, r)
}
