package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func makeSimpleProgressBar() {
  uiprogress.Start()

  bar := uiprogress.AddBar(100).AppendCompleted().PrependElapsed()

  for bar.Incr() {
    time.Sleep(time.Millisecond * 30)
  }

}

func makeProgressBar(lenght int, label string) {
  dummyStep := make([]string, lenght)

  for i := 0; i < lenght; i++ {
    dummyStep[i] = fmt.Sprintf("image %v", i)
  }

  uiprogress.Start()
  bar := uiprogress.AddBar(lenght).AppendCompleted().PrependElapsed()

  // prepend the current step to the bar
  bar.PrependFunc(func(b *uiprogress.Bar) string {
    return label+": " 
  })

  for i:= 0; i < lenght; i++ {
    time.Sleep(time.Millisecond * 10)
    bar.Incr()
  }
}

func parseMangaFromLink(url string) (title, id string) {
  urlParts := strings.Split(url, "/")
  return urlParts[len(urlParts)-1], urlParts[len(urlParts)-2]
}

func SingleDownload(mangaUrl string, dataSaver bool)  {
  // chapter is the last index, the other can be ignored
  chapterId, _ := parseMangaFromLink(mangaUrl)

  imageQuality := "dataSaver"
  downloadQuality := "data-saver"
  if !dataSaver {
    imageQuality = "data" // original quality 
    downloadQuality = "data"
  }
    
  imagesLink := GetChapterImage(chapterId)
  baseUrl, _ := jsonparser.GetString(imagesLink, "baseUrl")
  chapterHash, _ := jsonparser.GetString(imagesLink, "chapter", "hash")
  
  targetDir, err := makeDir(".", "singleDownload-"+chapterHash) 
  
  if err != nil {
    log.Fatal("Error making directories")
  }

  fmt.Println("Downloaing to: ", targetDir)


  jsonparser.ArrayEach(imagesLink, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
    imageName := string(value)
    imageUrl := baseUrl +"/"+ downloadQuality +"/"+ chapterHash +"/"+ imageName
    image := FetchImage(imageUrl)

    e := createFile(targetDir, imageName, image)

    if e != nil {
      log.Fatal("unable to download image", e)
    }

  }, "chapter", imageQuality)

  fmt.Println("Done!")
}

func DownloadManga(mangaUrl, lang string, 
  startChapter, endChapter int, dataSaver bool) {

    title, mangaId := parseMangaFromLink(mangaUrl) 
    imageQuality := "dataSaver"
    downloadQuality := "data-saver"
    if !dataSaver {
      imageQuality = "data" // original quality
      downloadQuality = "data"
    }
    
    // Fetch Manga Feed (ie, chapters and volumes)
    feed := GetFeed(mangaId, lang)
    lastChapter, _ := jsonparser.GetInt(feed, "total")

    if endChapter == 0 { 
      endChapter = int(lastChapter) 
    }

    // make target dir
    saveDir, err := makeDir(".", title)
    if err != nil {
      panic(err)
    }

    fmt.Println("Downloading ", title, "Start Chapter: ", startChapter, "End Chapter: ", endChapter)

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

    // makeProgressBar(int(pages), chapterNumber)

    // Downloads chapter images
    jsonparser.ArrayEach(chapterImages, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
      imageName := string(value)
      imageUrl := baseUrl +"/"+ downloadQuality +"/"+ chapterHash +"/"+ imageName
      image := FetchImage(imageUrl)

      e := createFile(targetDir, imageName, image)

      if e != nil {
        log.Fatal("unable to download image", e)
      }
      

    }, "chapter", imageQuality)
  }, "data" )
 }
