package controllers

import (
	"encoding/json"
	"net/http"

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

	// Check if url is valid
	// Check if url is already in database
	// If not, create a new short url
	// Return the short url

	w.Write([]byte(data.Url))
}
