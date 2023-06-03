package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/buger/jsonparser"
)

// exported, so i can reuse later...
func DownloadChapter(value []byte, dataType jsonparser.ValueType, 
  offset int, err error) {

}

// Yeah i stole it from some random guy on stackoverflow lmfao
func exists(path string) bool {
  _, err := os.Stat(path)
  if err == nil { return true }
  if os.IsNotExist(err) { return false }

  return false
}

func makeDir(baseDir, dirName string) (string, error) {
  dirPath := filepath.Join(baseDir, dirName)

  if exists(dirPath) { 
    return dirPath, nil
  } else {
    // Make dir if it doesn't exist
    err := os.Mkdir(dirPath, 777)

    if err == nil {
      return dirPath, nil
    } else { 
      return "", err
    }
  }
}

func DownloadManga(mangaId, lang string, startChapter, endChapter uint16) {
   if lang == "" {
     lang = "en"
   }

   feed := GetFeed(mangaId, lang)
   manga := GetManga(mangaId)

   title, _ := jsonparser.GetString(manga, "data", "attributes", "title", "en")
   lastChapter, _ := jsonparser.GetInt(feed, "total")

   if startChapter == 0 {
     startChapter = 1
   }

   if endChapter == 0 {
     endChapter = uint16(lastChapter)
   }

   // make target dir
   saveDir, err := makeDir(".", title)

   if err != nil {
     panic(err)
   }

   jsonparser.ArrayEach(feed, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    
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

    chapterImages := GetChapterImage(chapterId)
    // baseUrl, _ := jsonparser.GetString(chapterImages, "baseUrl")
    // hash, _ := jsonparser.GetString(chapterImages, "chapter")

    jsonparser.ArrayEach(chapterImages, DownloadChapter, "chapter", "dataSaver")
  }, "data" )

   fmt.Println(saveDir, lastChapter)
 }


