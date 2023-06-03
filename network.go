package main

import (
	"log"

	"github.com/sethgrid/pester"
)

const (
  HOST = "https://api.mangadex.org/"
  MANGA = HOST + "manga"
  CHAPTER = HOST + "chapter"
  IMAGE = HOST + "at-home/server/"
)

var phttp *pester.Client;

func init() {
  phttp = pester.New()
  phttp.Concurrency = 3
  phttp.MaxRetries = 7
  phttp.Backoff = pester.ExponentialBackoff
}

func GetFeed(mangaId, lang string) []byte {
   res, err := phttp.Get(MANGA + "/"+ mangaId + "/feed" + "?translatedLanguage[]=" + lang)
  
  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}

func GetMangaAggregate(mangaId, lang string) []byte  {
  res, err := phttp.Get(MANGA + "/"+ mangaId + "/aggregate" + "?translatedLanguage[]=" + lang)
  
  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}

func GetChapterImage(chapterId string) []byte {
  res, err := phttp.Get(IMAGE + chapterId)

  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}

func GetManga(id string ) []byte {
  res, err := phttp.Get(MANGA +"/"+ id)

  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}

func GetChapter(id string ) []byte {
  res, err := phttp.Get(CHAPTER + "/"+ id)

  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}

func SearchTitle(title string ) []byte {
  res, err := phttp.Get(MANGA + "?title=" + title)

  if err != nil {
    log.Fatal(err)
  }

  return parseBody(res.Body)
}
