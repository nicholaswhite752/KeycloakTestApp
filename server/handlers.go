package main

import "net/http"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}