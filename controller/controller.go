// Package controller loads the routes for each of the controllers.
package controller

import (
	"tech/controller/about"
	"tech/controller/cms/category"
	"tech/controller/debug"
	"tech/controller/home"
	"tech/controller/login"
	"tech/controller/notepad"
	"tech/controller/register"
	"tech/controller/static"
	"tech/controller/status"
)

// LoadRoutes loads the routes for each of the controllers.
func LoadRoutes() {
	about.Load()
	debug.Load()
	register.Load()
	login.Load()
	home.Load()
	static.Load()
	status.Load()
	notepad.Load()
	category.Load()
}
