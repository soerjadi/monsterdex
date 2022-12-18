package main

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/soerjadi/monsterdex/internal/config"
	"github.com/soerjadi/monsterdex/internal/delivery/rest"
	monsterHandler "github.com/soerjadi/monsterdex/internal/delivery/rest/monster"
	"github.com/soerjadi/monsterdex/internal/log"
	"github.com/soerjadi/monsterdex/internal/repository/access_token"
	"github.com/soerjadi/monsterdex/internal/repository/monster"
	"github.com/soerjadi/monsterdex/internal/repository/user"
	tokenUsecase "github.com/soerjadi/monsterdex/internal/usecase/access_token"
	monsterUsecase "github.com/soerjadi/monsterdex/internal/usecase/monster"
	userUsecase "github.com/soerjadi/monsterdex/internal/usecase/user"
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
	dataSource := fmt.Sprintf("user=%s password=%s	host=%s port=%s dbname=%s sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.Name,
	)

	db, err := sqlx.Open(cfg.Database.Driver, dataSource)
	if err != nil {
		log.ErrorWithFields("cannot connect to db", log.KV{"error": err})
		return
	}

	handlers, err := initiateHandler(cfg, db)
	if err != nil {
		log.ErrorWithFields("unable to initiate handler.", log.KV{
			"err": err,
		})
		return
	}

	r := mux.NewRouter()
	rest.RegisterHandlers(r, handlers...)

	srv := &http.Server{
		Addr:         fmt.Sprintf("0.0.0.0:%s", cfg.Server.Port),
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r, // Pass our instance of gorilla/mux in.
	}

	log.Info(fmt.Sprintf("Server running in port : %s", cfg.Server.Port))

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

func initiateHandler(cfg *config.Config, db *sqlx.DB) ([]rest.API, error) {
	monsterRepository, err := monster.GetRepository(db)
	if err != nil {
		return nil, fmt.Errorf("unable to initiate monsterRepository. ERR : %v", err)
	}

	tokenRepository, err := access_token.GetRepository(db)
	if err != nil {
		return nil, fmt.Errorf("unable to initiate tokenRepository. ERR : %v", err)
	}

	userRepository, err := user.GetRepository(db)
	if err != nil {
		return nil, fmt.Errorf("unable to initiate userRepository. ERR : %v", err)
	}

	usecase := monsterUsecase.GetUsecase(monsterRepository)
	tokenUsecase := tokenUsecase.GetUsecase(tokenRepository)
	userUsecase := userUsecase.GetUsecase(userRepository)

	handler := monsterHandler.NewHandler(usecase, tokenUsecase, userUsecase)

	return []rest.API{
		handler,
	}, nil
}

// func GenerateToken() string {
// 	randU64 := m.Uint64() + m.Uint64()
// 	hasher := sha512.New()
// 	byte := make([]byte, 8)

// 	binary.LittleEndian.PutUint64(byte, randU64)

// 	hasher.Write(byte)

// 	return hex.EncodeToString(hasher.Sum(nil))
// }
