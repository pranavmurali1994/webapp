package main

import (
    "encoding/json"
    "net/http"
    "os"
    "log"
)

func versionHandler(w http.ResponseWriter, r *http.Request) {
    msg := os.Getenv("APP_MESSAGE")
    if msg == "" {
        msg = "(empty)"
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": msg})
}

func main() {
    http.HandleFunc("/version", versionHandler)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }
    log.Printf("starting server on :%s\n", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
