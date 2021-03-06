package main

import (
	"context"
	"database/sql"
	"fmt"
	_ "gokomodo/docs"
	"gokomodo/internal/config"
	docs_handler "gokomodo/internal/delivery/http/docs"
	order_handler "gokomodo/internal/delivery/http/order"
	product_handler "gokomodo/internal/delivery/http/product"
	user_handler "gokomodo/internal/delivery/http/user"
	order_repository "gokomodo/internal/repository/psql/order"
	product_repository "gokomodo/internal/repository/psql/product"
	user_repository "gokomodo/internal/repository/psql/user"
	"gokomodo/pkg/logger"
	"gokomodo/pkg/service/jwt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

var (
	cfg         = config.Server()
	appLogger   = logger.NewApiLogger()
	db          = config.InitDatabase()
	jwtService  = jwt.NewJWTService()
	userRepo    = user_repository.NewUserRepository(db)
	productRepo = product_repository.NewProductRepository(db)
	orderRepo   = order_repository.NewOrderRepository(db)
)

func main() {
	psqlConn := config.InitDatabase()
	defer func(db *sql.DB) { _ = db.Close() }(psqlConn)

	router := mux.NewRouter()

	initHandler(router)
	http.Handle("/", router)

	appLogger.Info("gokomodo Service Run on " + cfg.Port)

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

func initHandler(router *mux.Router) {
	user_handler.UserHandler(router, jwtService, userRepo)
	product_handler.ProductHandler(router, jwtService, productRepo)
	order_handler.OrderHandler(router, jwtService, orderRepo, productRepo)
	docs_handler.DocsHandler(router)
}
