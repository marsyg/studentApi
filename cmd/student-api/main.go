package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/marsyg/studentApi/internal/config"
	"github.com/marsyg/studentApi/internal/http/handlers/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//setup router
	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New())
	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}
	fmt.Printf("server is started at :%s", cfg.HTTP_SERVER.Addr)
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("Filed to start server ")
		}
	}()
	<-done
	slog.Info("Shutting down the server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("failed tyo shutdown the server ", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown sucessfully")
}
