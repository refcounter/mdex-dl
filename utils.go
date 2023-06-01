package main

import (
	"log"
  "io/ioutil"
  "io"

	"github.com/valyala/fastjson"
)

var (
  parser fastjson.Parser
)

func ToJSON(str string) *fastjson.Value {
  v, e := parser.Parse(str)


  if e != nil {
    log.Fatal(e)
  }

  return v
}

func ParseBody(responseBody io.ReadCloser) string {
  defer responseBody.Close()
  bodyBytes, _ := ioutil.ReadAll(responseBody)
  
  return string(bodyBytes)
}

