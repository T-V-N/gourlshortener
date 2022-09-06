package main

import (
	"fmt"
	"net/http"

	server "github.com/T-V-N/gourlshortener/internal/app"
	"github.com/T-V-N/gourlshortener/internal/handler"
	"github.com/T-V-N/gourlshortener/internal/storage"
)

func main() {
	st := storage.NewStorage(map[string]string{})
	h := handler.InitHandler(st)
	server := server.InitServer(h)

	http.HandleFunc("/", server.HandleRequest)
	fmt.Println("go!")
	http.ListenAndServe(":8080", nil)
}