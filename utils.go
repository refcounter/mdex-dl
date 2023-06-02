package main

import (
  "io/ioutil"
  "io"

)

func parseBody(responseBody io.ReadCloser) []byte {
  defer responseBody.Close()
  bodyBytes, _ := ioutil.ReadAll(responseBody)
  
  return bodyBytes
}

