package main

import (
	"fmt"
	"github.com/VeronicaAlexia/BoluobaoAPI/Template"
	"github.com/VeronicaAlexia/BoluobaoAPI/boluobao/book"
	"github.com/VeronicaAlexia/BoluobaoAPI/pkg/config"
	"os"
	"testing"
)

var BookInfo Template.BookInfo

func GetContent(ChapID string) {
	contents := book.Content(ChapID)
	if contents != nil {
		content_text := []byte("\n\n\n" + contents.Data.Expand.Content)
		path := fmt.Sprintf("%v.txt", BookInfo.Data.NovelName)
		fl, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return
		}
		defer fl.Close()
		if _, err = fl.Write(content_text); err != nil {
			fmt.Println("Error:", err)
		}
	}
}

func TestDownload(t *testing.T) {
	book_id := "512854"
	config.AppConfig.App = false
	BookInfo = book.GET_BOOK_INFORMATION(book_id)
	if BookInfo.Status.HTTPCode == 200 {
		fmt.Println("bookName:", BookInfo.Data.NovelName)
		fmt.Println("AuthorName:", BookInfo.Data.AuthorName)
		fmt.Println("BookID:", BookInfo.Data.NovelID)
		fmt.Println("bookCover:", BookInfo.Data.NovelCover)

		if err := os.WriteFile(
			fmt.Sprintf("%v.txt", BookInfo.Data.NovelName),
			[]byte(BookInfo.Data.NovelName+"\n\n"), 0777); err != nil {
			fmt.Println(err)
		}
		for _, ChapID := range book.Catalogue(book_id) {
			GetContent(ChapID)
		}
	} else {
		fmt.Println("BookInfo Error:", BookInfo.Status.Msg)
	}
}
