package handler

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/PoliteCupcake/TemplApp/layout"
)

type DiaryEntry struct {
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
	timeStamp := time.Now()
	request.ParseForm()

	newContent := request.FormValue("content")
	entry := DiaryEntry{Content: newContent, Date: timeStamp}
	card := layout.DiaryCard(entry.Content, entry.Date)
	card.Render(context.Background(), writer)
}
