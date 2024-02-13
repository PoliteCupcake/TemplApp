package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/PoliteCupcake/TemplApp/layout"
)

type DiaryEntry struct {
	Title   string
	Content string
	Date    time.Time
}

func AddEntry(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	newTitle := request.FormValue("title")
	timeStamp := time.Now()
	newContent := request.FormValue("content")

	entry := DiaryEntry{Content: newContent, Title: newTitle, Date: timeStamp}
	card := layout.DiaryCard(entry.Content, entry.Title, entry.Date)
	card.Render(context.Background(), writer)
}

func GetForm(writer http.ResponseWriter, request *http.Request) {
	form := layout.EntryForm()
	form.Render(context.Background(), writer)
}
