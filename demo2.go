package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"

	"time"

	"github.com/gorilla/mux"
)

func main() {
	go exec.Command("xdg-open", "http://localhost:2007").Run()
	go exec.Command("xdg-open", "http://localhost:2007/echo/Hello%20World").Run()
	go exec.Command("xdg-open", "http://localhost:2007/f/").Run()

	httpDemo()
}

func httpDemo() {
	logger := func(callback http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			callback(w, r)
			log.Printf("%s\t%s\t%s\n", r.Method, r.RequestURI, time.Duration(time.Now().Sub(start)))
		}
	}

	router := mux.NewRouter()
	router.PathPrefix("/f/").
		HandlerFunc(logger(http.StripPrefix("/f/", http.FileServer(http.Dir("/tmp/"))).ServeHTTP))
	router.HandleFunc("/echo/{v:[a-zA-Z0-9 ]+}", logger(func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		fmt.Fprintf(w, "echo %q", routeVars["v"])
	}))
	router.HandleFunc("/", logger(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`these aren't the droids you're looking for`))
	}))
	http.ListenAndServe(":2007", router)
}
