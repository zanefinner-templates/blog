package accounts

import (
	"fmt"
	"net/http"

	"github.com/zanefinner-templates/blog/template"
)

//Signup resolves /accounts/signup
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		template.ExecuteTemplate(w, "index.tmpl", nil)

	} else if r.Method == http.MethodPost {

		//Take in data
		//Validate (prevent sql injection)
		//Match to see if exists
		//If all is good, start a session
	}
	fmt.Println("ANY: /accounts/signup")
}
