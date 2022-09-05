package main

import (
	"fmt"
	"net/http"

	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

func main() {
	st := storage.NewStorage()
	handler := handler.InitHandler(st)
	http.HandleFunc("/", handler.HandleRequest)
	fmt.Println("go!")
	http.ListenAndServe(":8080", nil)
}