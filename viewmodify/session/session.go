// Package uri adds URI shortcuts to the view template.
package session

import (
	"net/http"
	"tech/lib/flight"

	"github.com/blue-jay/core/view"
)

// Modify sets BaseURI, CurrentURI, ParentURI, and the GrandparentURI
// variables for use in the templates.
func Modify(w http.ResponseWriter, r *http.Request, v *view.Info) {
	c := flight.Context(w, r)
	v.Vars["ArchivingToken"] = 0
	v.Vars["ArchivingId"] = 0
	v.Vars["CaseId"] = 0

	if val, ok := c.Sess.Values["tech_token"]; ok {
		v.Vars["TechToken"] = val
	}
	if val, ok := c.Sess.Values["tech_id"]; ok {
		v.Vars["TechId"] = val
	}
	// if val, ok := c.Sess.Values["case_id"]; ok {
	// 	v.Vars["CaseId"] = val
	// }
}
