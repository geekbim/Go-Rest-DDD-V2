package main

import (
	"context"
	"database/sql"
	"fmt"
	"majoo/internal/config"
	user_handler "majoo/internal/delivery/http/user"
	user_repository "majoo/internal/repository/psql/user"
	"majoo/pkg/logger"
	"majoo/pkg/service/jwt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	cfg        = config.Server()
	appLogger  = logger.NewApiLogger()
	db         = config.InitDatabase()
	jwtService = jwt.NewJWTService()
	userRepo   = user_repository.NewUserRepository(db)
)

func main() {
	psqlConn := config.InitDatabase()
	defer func(db *sql.DB) { _ = db.Close() }(psqlConn)

	router := mux.NewRouter()

	initHandler(router, cfg)
	http.Handle("/", router)
	appLogger.Info("Majoo Service Run on " + cfg.Port)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		err := http.ListenAndServe(cfg.Port, router)
		if err != nil {
			appLogger.Error(err)
			cancel()
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	select {
	case v := <-quit:
		appLogger.Error(fmt.Sprintf("signal.Notify: %v", v))
	case done := <-ctx.Done():
		appLogger.Error(fmt.Sprintf("ctx.Done: %v", done))
	}
}

func initHandler(router *mux.Router, cfg config.ServerConfig) {
	user_handler.UserHandler(router, cfg, jwtService, userRepo)
}
