package controllers

import "net/http"

func CreteShortUrl(w http.ResponseWriter, r *http.Request) {
	// TODO: Implement CreateShortUrl cotroller
	// Verify the request and body
	// Check if the url is already in the database
	// If not, create a new short url
	// Return the short url

	w.Write([]byte("Hello World!"))
}
