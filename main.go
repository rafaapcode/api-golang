package main

import (
	"net/http"
	"time"

	"github.com/rafaapcode/api-golang/internal/api"
	"gitlab.com/0x4149/logz"
)

// import (
// 	"net/http"
// 	"time"

// 	"gitlab.com/0x4149/logz"
// )

const local bool = true

var srv *http.Server

func init() {
	if local {
		logz.VerbosMode()
	}

	logz.Run()
}

func main() {
	router := api.NewRouter()
	versionRouter := router.Version(1)
	versionRouter.Prefix("/auth").AuthEndpoints()
	versionRouter.Prefix("/user").UserEndpoints()

	if local {
		srv = &http.Server{
			Addr:         ":8080",
			Handler:      router,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		logz.Info("Server is running")
		logz.Fatal(srv.ListenAndServe())
	} else {
		srv = &http.Server{
			Addr:         ":443",
			Handler:      router,
			ReadTimeout:  time.Second * 10,
			WriteTimeout: time.Second * 10,
		}

		var cert string = ""
		var certKey string = ""

		logz.Info("Server is running")
		logz.Fatal(srv.ListenAndServeTLS(cert, certKey))
	}

}
