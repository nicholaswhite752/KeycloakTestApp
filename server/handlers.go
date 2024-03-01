package main

import (
	"net/http"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.Cookies())
	//cookie, _ := r.Cookie("next-auth.csrf-token")
	//fmt.Println(cookie.Value)
	w.Write([]byte("hi"))
}

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hi"))
}