package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/buger/jsonparser"
	"github.com/gosuri/uiprogress"
)

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
    err := os.MkdirAll(dirPath, 0755)

    if err == nil {
      return dirPath, nil
    } else { 
      return "", err
    }
  }
}

func createFile(basePath, filename string, imageBytes io.ReadCloser) error {
  //Create a empty file
	file, err := os.Create(filepath.Join(basePath, filename))
	if err != nil {
		return err
	}
	defer file.Close()

  defer imageBytes.Close()

	//Write the bytes to the file
	_, err = io.Copy(file, imageBytes)
	if err != nil {
		return err
	}

  return nil
}

func makeProgressBar() {
  uiprogress.Start()            // start rendering
  bar := uiprogress.AddBar(100) // Add a new bar


  for bar.Incr() {
    time.Sleep(time.Millisecond * 20)
  }
}

func DownloadManga(mangaId, lang string, 
  startChapter, endChapter uint16, dataSaver bool) {
   if lang == "" {
     lang = "en"
   }

   imageQuality := "dataSaver"
   downloadQuality := "data-saver"
   if !dataSaver {
     imageQuality = "data" // original quality
     downloadQuality = "data"
   }

   // Fetch Manga Feed (ie, chapters and volumes)
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

   fmt.Println("Downloading ", title)
   fmt.Println("Start Chapter: ", startChapter, "End Chapter: ", endChapter)

   jsonparser.ArrayEach(feed, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    
      chapterId, err := jsonparser.GetString(value, "id")

      if err != nil {
        log.Panic("No available chapter id", err)
      }
    
    chapterNumber, _ := jsonparser.GetString(value, "attributes", "chapter")
    volume, _ := jsonparser.GetString(value, "attributes", "volume")
    // pages, _ :=  jsonparser.GetInt(value, "attributes", "pages")

    fmt.Println("Chapter: ", chapterNumber)

    chapterImages := GetChapterImage(chapterId)
    baseUrl, _ := jsonparser.GetString(chapterImages, "baseUrl")
    chapterHash, _ := jsonparser.GetString(chapterImages, "chapter", "hash")

    targetDir, err := makeDir(saveDir, filepath.Join("volume-"+volume, "chapter-"+chapterNumber))
    if err != nil {log.Fatal(err)}

    // Downloads chapter images
    jsonparser.ArrayEach(chapterImages, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
      imageName := string(value)
      imageUrl := baseUrl +"/"+ downloadQuality +"/"+ chapterHash +"/"+ imageName
      image := FetchImage(imageUrl)

      e := createFile(targetDir, imageName, image)
      fmt.Println(imageUrl)

      if e != nil {
        log.Fatal("unable to download image", e)
      }
      

    }, "chapter", imageQuality)
  }, "data" )
 }
