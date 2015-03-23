package main

import (
	"fmt"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	go exec.Command("xdg-open", "http://localhost:2007").Run()
	go exec.Command("xdg-open", "http://localhost:2007/echo/Hello%20World").Run()
	go exec.Command("xdg-open", "http://localhost:2007/f/").Run()

	httpDemo()
}

func httpDemo() {
	router := mux.NewRouter()

	router.PathPrefix("/f/").Handler(http.StripPrefix("/f/", http.FileServer(http.Dir("/tmp/"))))

	router.HandleFunc("/echo/{v:[a-zA-Z0-9 ]+}", func(w http.ResponseWriter, r *http.Request) {
		routeVars := mux.Vars(r)
		fmt.Fprintf(w, "echo %q", routeVars["v"])
	})

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`these aren't the droids you're looking for`))
	})

	http.ListenAndServe(":2007", router)
}
