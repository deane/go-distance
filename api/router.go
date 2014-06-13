package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func newRouter() *mux.Router {
  router := mux.NewRouter()

  router.HandleFunc(
    "/distance/{unit:(km|miles)}", calculateDistance,
  ).Methods("POST")

  return router
}

// Runserver starts the service
func RunServer(host string) {
  http.Handle("/", newRouter())

  log.Printf("Starting server at %s", host)

  log.Fatalln(http.ListenAndServe(host, nil))
}
