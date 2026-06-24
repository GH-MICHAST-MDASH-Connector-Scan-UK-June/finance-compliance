package main

import (
"encoding/json"
"fmt"
"net/http"
)

type AppStatus struct {
Name   string +"`"+json:"name"+"`"+
Status string +"`"+json:"status"+"`"+
}

func main() {
fmt.Println("finance-compliance - Go Application")

http.HandleFunc("/", handler)
http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
status := AppStatus{Name: "finance-compliance", Status: "running"}
json.NewEncoder(w).Encode(status)
}
