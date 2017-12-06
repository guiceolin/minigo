package handlers

import (
	"html/template"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/guiceolin/minigo/interactors"
)

func (e Env) NewURLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/new_url.html"))
	tmpl.Execute(w, nil)
}

func UnshortURLHandler(i interactors.UrlInteractor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := i.FindShortURL(chi.URLParam(r, "short"))
		i.IncrementAccess(url)

		http.Redirect(w, r, url.Original, 302)
	}
}

func GetShortURLInfo(i interactors.UrlInteractor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		url := i.FindShortURL(chi.URLParam(r, "short"))

		tmpl := template.Must(template.ParseFiles("templates/url_info.html"))
		tmpl.Execute(w, url)
	}
}

func CreateUrlHandler(i interactors.UrlInteractor) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		encoded := i.CreateUrl(r.FormValue("url"))
		http.Redirect(w, r, "/"+encoded+"/info", 302)
	}
}
