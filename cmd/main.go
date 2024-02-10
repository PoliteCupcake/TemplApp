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
	//entry := handler.ShowSample("Test")

	landing := layout.Base("Diary")
	//landing := Base("Title")

	http.Handle("/", templ.Handler(landing))
	http.HandleFunc("/post_card", handler.AddEntry)
	http.HandleFunc("/templ_test", handler.AddEntry)

	fmt.Println("Starting server ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server shutdown")
	}

	fmt.Println("Waiting for Requests ...")

}
