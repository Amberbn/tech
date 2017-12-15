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
	router.Get(uri+"/edit/:id", Edit, c...)
	router.Patch(uri+"/edit/:id", Update, c...)
	router.Get(uri+"/delete/:id", Delete, c...)
	router.Delete(uri+"/:id", Destroy, c...)
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

func Edit(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := category.ById(c.DB, c.Param("id"))
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("cms/category/edit")
	c.Repopulate(v.Vars, "name_category")
	v.Vars["item"] = item
	v.Render(w, r)
}

func Update(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	if !c.FormValid("name_category") {
		Edit(w, r)
		return
	}

	_, err := category.Update(c.DB, r.FormValue("name_category"), c.Param("id"))
	if err != nil {
		c.FlashErrorGeneric(err)
		Edit(w, r)
		return
	}

	c.FlashSuccess("Item updated.")
	c.Redirect(uri)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	item, _, err := category.ById(c.DB, c.Param("id"))
	if err != nil {
		c.FlashErrorGeneric(err)
		c.Redirect(uri)
		return
	}

	v := c.View.New("cms/category/delete")
	v.Vars["item"] = item
	v.Render(w, r)
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	c := flight.Context(w, r)

	b := c.Param("id")
	fmt.Println("parammmm=", b)
	_, err := category.DeleteHard(c.DB, c.Param("id"))
	if err != nil {
		c.FlashErrorGeneric(err)
	} else {
		c.FlashNotice("Item deleted.")
	}

	c.Redirect(uri)
}
