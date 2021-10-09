package main

import (
    "net/http"
)


func indexHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<h1>Hello World!</h1>"))
}

func main() {
    port := "3000"
    mux := http.NewServeMux()
    mux.HandleFunc("/", indexHandler)
    http.ListenAndServe(":"+port, mux)

}