package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PoliteCupcake/TemplApp/layout"
)

type DiaryEntry struct {
	Title   string
	Content string
	Date    time.Time
}

func ShowSample(content string) DiaryEntry {
	timeStamp := time.Now()
	entry := DiaryEntry{Content: content, Date: timeStamp}
	return entry
}

func AddSampleDay(writer http.ResponseWriter, request *http.Request) {
	timeStamp := time.Now()
	entry := DiaryEntry{Content: "Content woohoo", Date: timeStamp}
	fmt.Fprintf(writer, entry.Content)
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
