package api

import "net/http"

func (r router) AuthEndpoints() {
	r.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Register"))
	}) 

	r.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("login"))
	}) 

	r.HandleFunc("/logout", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("logout"))
	}) 
}
