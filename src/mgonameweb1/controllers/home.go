// home
package controllers

import (
	"fmt"
	"html/template"
	"mgonameweb1/models"
	"net/http"
	"strings"
)

/*
func main() {
	fmt.Println("Hello World!")
}
*/

func HomeIndexController(r http.ResponseWriter, rq *http.Request) {
	t, err := template.ParseFiles("src/mgonameweb1/views/home/index.html")
	if err != nil {
		r.WriteHeader(http.StatusNotFound)
		r.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	t.Execute(r, nil)
}

func HomeValidateController(r http.ResponseWriter, rq *http.Request) {
	errString := "no errors"
	firstName := strings.TrimSpace(rq.FormValue("first"))
	lastName := strings.TrimSpace(rq.FormValue("last"))
	if firstName == "" || lastName == "" {
		errString = "An entry is missing"
	} else {
		name := models.Name{FirstName: firstName, LastName: lastName}
		count, err := name.GetDuplicateCount()
		if err != nil {
			errString = fmt.Sprintf("Error on dupcheck: %v", err)
		} else {
			if count == 0 {
				if err := name.AddName(); err != nil {
					errString = fmt.Sprintf("Error on insert: %v", err)
				}
			} else {
				errString = "That name is already on file"
			}
		}
	}

	names, listErr := models.GetAllNames()
	if listErr != nil {
		errString = fmt.Sprintf("Error on read: %v", listErr)
	}

	t, err := template.ParseFiles("src/mgonameweb1/views/home/validate.html")
	if err != nil {
		r.WriteHeader(http.StatusNotFound)
		r.Write([]byte("File not found"))
		fmt.Println("Template parse error:", err)
		return
	}
	t.Execute(r, struct {
		Names   []models.Doc
		Message string
	}{Names: names, Message: errString})

}
