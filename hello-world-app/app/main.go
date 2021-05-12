package main

import (
    "fmt"
    "os"
    "net/http"
    "io/ioutil"
    "log"
    "encoding/json"
)

func main() {

  dataFileLocation := "../data/data.json"
  if len(os.Args) == 2 {
    dataFileLocation = os.Args[1]
  }

  jsonFile, err := os.Open(dataFileLocation)
  // if we os.Open returns an error then handle it
  if err != nil {
      log.Fatal(err)
  }
  defer jsonFile.Close()

  byteValue, _ := ioutil.ReadAll(jsonFile)

  var result map[string]interface{}
  json.Unmarshal([]byte(byteValue), &result)

  fmt.Println(result["value"])

  text := result["value"]

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello, World!") 
    fmt.Fprintln(w, text) 
  })
  http.ListenAndServe(":3001", nil)
}