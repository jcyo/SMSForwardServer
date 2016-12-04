package main

import (
  "encoding/json"
  "log"
  "fmt"
  "net/http"
  "io/ioutil"
)

type Message struct {
    From string
    Message  string
}

func main() {

    http.HandleFunc("/message", func(w http.ResponseWriter, req *http.Request) {
      body, err := ioutil.ReadAll(req.Body)
      if err != nil {
          panic(err)
      }
      log.Println(string(body))
      var t Message
      err = json.Unmarshal(body, &t)
      if err != nil {
          panic(err)
      }
      fmt.Fprintf(w, t.From)
    })

    log.Fatal(http.ListenAndServe(":8081", nil))

}
