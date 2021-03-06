package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mmnalaka/go-short-url/models/urls"
	"github.com/mmnalaka/go-short-url/pkg/utils"
)

type UrlCreatePayload struct {
	Url string `json:"url"`
}

func CreteShortUrl(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement CreateShortUrl cotroller
	data := &UrlCreatePayload{}

	if err := json.NewDecoder(r.Body).Decode(data); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	// check if the url is valid
	if !utils.IsValidUrl(data.Url) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid url"))
		return
	}

	sUrl := urls.GetUrl(data.Url)
	if sUrl != nil {
		w.WriteHeader(http.StatusOK)
		render.Respond(w, r, sUrl)
		return
	}

	url := urls.CreateUrl(data.Url)
	if url == nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating short url"))
		return
	}

	w.WriteHeader(http.StatusOK)
	render.Respond(w, r, url)
}
