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
	Id      int64
}

func AddEntry(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()

	newTitle := request.FormValue("title")
	newContent := request.FormValue("content")
	timeStamp := time.Now()
	id := timeStamp.Unix()

	entry := DiaryEntry{Content: newContent, Title: newTitle, Date: timeStamp, Id: id}
	card := layout.DiaryCard(entry.Content, entry.Title, entry.Date, entry.Id)
	card.Render(context.Background(), writer)
}

func GetForm(writer http.ResponseWriter, request *http.Request) {
	form := layout.EntryForm()
	form.Render(context.Background(), writer)
}
