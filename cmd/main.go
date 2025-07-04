package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	db "diffinlist/internal/db/generated"
	"diffinlist/internal/server"

	"github.com/jackc/pgx/v5/pgxpool"
)

func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")
	stop() // Allow Ctrl+C to force shutdown

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	// DB INIT START
	// pgx pool starts a pool thats concurrency safe
	dbpool, errDbConn := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	// connects to db via url
	if errDbConn != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", errDbConn)
		os.Exit(1)
	}
	defer dbpool.Close()
	queries := db.New(dbpool)
	// DB INIT END

	server := server.NewServer(queries)

	// Create a done channel to signal when the shutdown is complete
	done := make(chan bool, 1)

	// Run graceful shutdown in a separate goroutine
	go gracefulShutdown(server, done)

	fmt.Println("Starting server on port 8080")
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Wait for the graceful shutdown to complete
	<-done
	log.Println("Graceful shutdown complete.")
}
