package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/guiceolin/minigo/models"
	"github.com/xor-gate/bjf"
)

func (e Env) findURL(shortID string) models.Url {
	url := models.Url{}
	e.DB.Find(&url, bjf.Decode(shortID))
	return url
}

func (e Env) CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	original := r.FormValue("url")

	url := models.Url{Original: original, Count: 0}

	e.DB.Create(&url)
	var idToEncode = strconv.FormatUint(uint64(url.ID), 10)

	http.Redirect(w, r, "/"+bjf.Encode(idToEncode)+"/info", 302)
}

func (e Env) NewURLHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/new_url.html"))
	tmpl.Execute(w, nil)
}

func (e Env) UnshortURLHandler(w http.ResponseWriter, r *http.Request) {
	url := e.findURL(chi.URLParam(r, "short"))
	url.Count++
	e.DB.Save(&url)

	http.Redirect(w, r, url.Original, 302)
}

func (e Env) ShortURLInfo(w http.ResponseWriter, r *http.Request) {
	url := e.findURL(chi.URLParam(r, "short"))

	tmpl := template.Must(template.ParseFiles("templates/url_info.html"))
	tmpl.Execute(w, url)
}
