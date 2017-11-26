package handlers

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/guiceolin/minigo/models"
	"github.com/xor-gate/bjf"
)

func (e Env) CreateUrlHandler(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	original := queryValues.Get("url")

	url := models.Url{Original: original, Count: 0}

	e.DB.Create(&url)
	var idToEncode = strconv.FormatUint(uint64(url.ID), 10)

	http.Redirect(w, r, "/"+bjf.Encode(idToEncode), 302)
}

func (e Env) UnshortURLHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := chi.URLParam(r, "short")
	url := models.Url{}
	e.DB.Find(&url, bjf.Decode(shortURL))
	url.Count++
	e.DB.Save(&url)

	http.Redirect(w, r, url.Original, 302)
}
