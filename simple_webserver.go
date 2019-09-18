package main

import (
  "fmt"
  "net/http"
  "os"
)

func index_handler(w http.ResponseWriter, r *http.Request){
  name,err := os.Hostname()
  if err != nil {
	  fmt.Fprintln(w, err)
  }
  fmt.Fprintf(w, "This is v10. Served from:" + name + "\n")
}

func contact_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "contacts page")
}


func main() {
  http.HandleFunc("/", index_handler)
  http.HandleFunc("/contact/", contact_handler)
  err := http.ListenAndServe(":8000",nil)
  if err != nil    {
    fmt.Println(err)
    return
  }
}
