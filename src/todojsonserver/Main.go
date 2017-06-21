package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter()

    log.Fatal(http.ListenAndServe(":8080", router))
}
// http://localhost:8080
// http://localhost:8080/todos - GET
// http://localhost:8080/todos - POST
// http://localhost:8080/todos/{todoId}