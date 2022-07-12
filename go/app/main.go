package main

import (
	"log"
	"net/http"

	"go-http-server/infrastructure"
	"go-http-server/interface/controller"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	hh := infrastructure.NewHttpHandler()
	ic := controller.NewIndexController(hh)
	sc := controller.NewShowbannerController(hh)

	http.HandleFunc("/", ic.View)
	http.HandleFunc("/showbanner", sc.View)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Failed Listen and Serve. due to an error: ", err)
	}
}
