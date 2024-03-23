package main

import (
	"fmt"
	"net/http"

	"github.com/PoliteCupcake/TemplApp/handler"
	"github.com/PoliteCupcake/TemplApp/layout"
	"github.com/a-h/templ"
)

func main() {
	fmt.Println("Server started..")

	landing := layout.Base("Diary")

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static/"))))

	http.Handle("/", templ.Handler(landing))
	http.HandleFunc("/post_card", handler.AddEntry)

	fmt.Println("Starting server ...")
	fmt.Println("Running on localhost 8080")
	fmt.Println("Waiting for Requests ...")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server shutdown")
	}
}
