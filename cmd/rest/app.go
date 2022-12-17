package main

import (
	"context"
	"fmt"

	// "log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/soerjadi/monsterdex/internal/config"
	"github.com/soerjadi/monsterdex/internal/log"
)

func main() {
	cfg, err := config.Init()
	if err != nil {
		log.ErrorWithFields("[Config] error reading config from file", log.KV{
			"err": err,
		})
		return
	}

	// open database connection
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	_, err = sqlx.Open(cfg.Database.Driver, dataSource)
	if err != nil {
		log.ErrorWithFields("cannot connect to db", log.KV{"error": err})
		return
	}

	r := mux.NewRouter()
	// Add your routes as needed

	srv := &http.Server{
		Addr:         fmt.Sprintf("127.0.0.1:%s", cfg.Server.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Info("Server running in port : 8080")

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.ErrorWithFields("error running apps", log.KV{
				"err": err,
			})
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait  for.
	ctx, cancel := context.WithTimeout(context.Background(), cfg.WaitTimeout())
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	log.Info("shutting down")
	os.Exit(0)
}
