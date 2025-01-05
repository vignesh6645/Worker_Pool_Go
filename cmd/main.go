package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/vignesh/file_handler_service/internal/config"
	"github.com/vignesh/file_handler_service/internal/handler"
	"github.com/vignesh/file_handler_service/internal/pool"
)

func main() {
	cfg := config.Load()
	workerPool := pool.NewPool(cfg.WorkerCount)

	mux := http.NewServeMux()
	mux.HandleFunc("/job", handler.JobHandler(workerPool))
	mux.HandleFunc("/batch-jobs", handler.BatchJobHandler(workerPool))

	server := &http.Server{
		Addr:    cfg.ServerAddress,
		Handler: mux,
	}

	go func() {
		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
		<-sigChan
		log.Println("Shutting down server...")
		workerPool.Shutdown()
		server.Close()
	}()

	log.Printf("Server started at %s\n", cfg.ServerAddress)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
