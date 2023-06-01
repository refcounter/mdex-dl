package main

import (
	"log"

	rhttp "github.com/hashicorp/go-retryablehttp"
)

const (
  HOST = "https://api.mangadex.org/"
  MANGA = HOST + "manga"
  CHAPTER = HOST + "chapter"
  IMAGE = HOST + "at-home/server/"
)
func GetChapters(mangaId, lang string) string  {
  res, err := rhttp.Get(MANGA + "/"+ mangaId + "/aggregate" + "?translatedLanguage[]=" + lang)
  
  if err != nil {
    log.Fatal(err)
  }

  return ParseBody(res.Body)
}

func GetChapterImage(chapterId string, forceHttps string) string {
  res, err := rhttp.Get(IMAGE + chapterId + "?forcePort443=" + forceHttps)

  if err != nil {
    log.Fatal(err)
  }

  return ParseBody(res.Body)
}

func GetManga(id string) string {
  res, err := rhttp.Get(MANGA +"/"+ id)

  if err != nil {
    log.Fatal(err)
  }

  return ParseBody(res.Body)
}

func GetChapter(id string) string {
  res, err := rhttp.Get(CHAPTER + "/"+ id)

  if err != nil {
    log.Fatal(err)
  }

  return ParseBody(res.Body)
}

func SearchTitle(title string) string {
  res, err := rhttp.Get(MANGA + "?title=" + title)

  if err != nil {
    log.Fatal(err)
  }

  return ParseBody(res.Body)
}
