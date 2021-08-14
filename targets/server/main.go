package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Contra-Culture/gorr"
)

func main() {
	rr, err := gorr.New(func(r *gorr.RouterProxy) {
		r.OnError(gorr.NotFoundError, respondWithNotFoundError)
		r.OnError(gorr.MethodNotAllowedError, respondWithMethodNotAllowed)
		r.OnError(gorr.InternalServerError, respondWithInternalServerError)
		r.Before(func(w http.ResponseWriter, r *http.Request) { fmt.Printf("\n\t-> %s %s", r.Method, r.URL.String()) })
		r.BeforeMethod(func(w http.ResponseWriter, r *http.Request, ps map[string]string) { w.Write([]byte("beforehook\n")) })
		r.AfterMethod(func(w http.ResponseWriter, r *http.Request, ps map[string]string) { fmt.Print("\nafterhook\n") })
		r.After(func(w http.ResponseWriter, r *http.Request) {})
		r.ShowRoutes("routes", func() bool { return true })
		r.Root("welcome", "welcome page", func(n *gorr.NodeProxy) {
			// n.Method()
		})
	})
	if err != nil {
		panic(err)
	}
	err = http.ListenAndServe(":8080", rr)
	if err != nil {
		log.Fatal(err)
	}
}
func respondWithURL(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	w.Write([]byte(fmt.Sprintf("%s", ps)))
	w.WriteHeader(http.StatusOK)
}
func respondWithInternalServerError(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	w.Write([]byte("{\"error\":\"internal server error\"}"))
	w.WriteHeader(http.StatusInternalServerError)
}
func respondWithMethodNotAllowed(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	w.Write([]byte("{\"error\":\"method not allowed\"}"))
	w.WriteHeader(http.StatusMethodNotAllowed)
}
func respondWithNotFoundError(w http.ResponseWriter, r *http.Request, ps map[string]string) {
	w.Write([]byte("{\"error\":\"not found\"}"))
	w.WriteHeader(http.StatusNotFound)
}
