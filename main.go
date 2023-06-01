package main

import (
	"fmt"
  "strconv"

)

func DownloadManga(mangaId, lang string, startChapter, endChapter int) {
  manga := ToJSON(GetManga(mangaId))
  title := manga.GetStringBytes("data", "attributes", "title", "en")
  lastVolume := manga.GetStringBytes("data", "attributes", "lastVolume")
  lastChapter := manga.GetStringBytes("data", "attributes", "lastChapter")

  if lang == "" {
    lang = "en"
  }

  if startChapter == 0 {
    startChapter = 1
  }
  if endChapter == 0 {
    endChapter, _ = strconv.Atoi(string(lastChapter))
  }

  fmt.Println(string(title), lastVolume, lastChapter)
}

func main() {
  testId := "ff861098-94a8-470e-ad24-f21f691d3a5d"
  DownloadManga(testId, "pt-br", 0, 0)
}
