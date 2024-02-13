package main

import (
	"fmt"
	"net/http"

	"github.com/PoliteCupcake/TemplApp/handler"
	"github.com/PoliteCupcake/TemplApp/layout"
	"github.com/a-h/templ"
)

func main() {
	fmt.Println("Main started..")

	landing := layout.Base("Diary")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", templ.Handler(landing))
	http.HandleFunc("/get_form", handler.GetForm)
	http.HandleFunc("/post_card", handler.AddEntry)

	fmt.Println("Starting server ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server shutdown")
	}

	fmt.Println("Waiting for Requests ...")

}
