package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/marsyg/studentApi/internal/config"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the server"))
	})
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("server is started at :%s", cfg.HTTP_SERVER.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Filed to start server ")
	}
}
