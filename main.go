package main

import (
	"fmt"

	"github.com/buger/jsonparser"
)


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

   jsonparser.ArrayEach(feed, processChapters, "data" )

   fmt.Println(string(title), lastChapter)
 }


func main() {
  testId := "3d0f2aab-2c57-4284-a2f0-ca4520130d4e"
  DownloadManga(testId, "pt-br",  0, 0)
  // fmt.Println(GetFeed(testId, "pt-br"))
}
