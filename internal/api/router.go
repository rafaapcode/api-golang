package api

import (
	"fmt"
	"net/http"
	"strconv"
)

type router struct {
	*http.ServeMux
}

func NewRouter() router {
	return router{http.NewServeMux()}
}

func (r *router) Version(v int) *router {
	mux := NewRouter()
	var handler http.Handler = mux 
	versionPath := fmt.Sprintf("/api/v%s", strconv.Itoa(v))
	r.Handle(versionPath+"/", http.StripPrefix(versionPath, handler))

	return &mux
}

func (r *router) Prefix(path string) router {
	mux := NewRouter()
	var handler http.Handler = mux 
	r.Handle(path+"/", http.StripPrefix(path, handler))

	return mux
}