package handlers

import (
	"github.com/profemzy/dapsProject/pkg/render"
	"net/http"
)

// Home is the home page handler
func Home(writer http.ResponseWriter, _ *http.Request) {
	render.RenderTemplate(writer, "home.page.tmpl")
}

// About is the about page handler
func About(writer http.ResponseWriter, _ *http.Request) {
	//_, _ = fmt.Fprintf(writer, "About Page " + currentDateTime())
	render.RenderTemplate(writer, "about.page.tmpl")
}

// currentDateTime return current date and time in string format
//func currentDateTime() string{
//	dt := time.Now()
//	return "Current date and time: " + dt.String()
//}