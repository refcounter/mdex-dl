package main

import (
	"fmt"
	"log"

	"github.com/buger/jsonparser"
)

func processChapters(value []byte, dataType jsonparser.ValueType, 
  offset int, err error) {
    chapterId, err := jsonparser.GetString(value, "id")

    if err != nil {
      log.Panic("No available chapter id", err)
    }
    
    // chapter may (or may not have a title) 
    title, _ := jsonparser.GetString(value, "attributes", "title")

    chapterNumber, _ := jsonparser.GetString(value, "attributes", "chapter")
    volume, _ := jsonparser.GetString(value, "attributes", "volume")
    pages, _ :=  jsonparser.GetInt(value, "attributes", "pages")

    fmt.Println(title, chapterNumber, volume, pages, chapterId)
}


