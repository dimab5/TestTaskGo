package main

import (
	"TestTaskGo/internal/config"
	"TestTaskGo/internal/http-server/handlers/walletActions"
	mw "TestTaskGo/internal/http-server/middleware"
	"TestTaskGo/internal/storage"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	_ "github.com/lib/pq"
	"log"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	cfg := config.MustLoad()
	myLog := SetUpLogger()

	myLog.Info("starting app")

	db, err := storage.ConnectToDB()
	if err != nil {
		log.Panic(err)
	}
	s := storage.Storage{Db: db}
	myLog.Info("Connect to db")

	router := chi.NewRouter()
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(mw.LoggingMiddleware)

	router.Post("/api/v1/wallet", walletActions.NewWalletCreator(&s))
	router.Post("/api/v1/wallet/{walletId}/send", walletActions.NewTransferMoney(&s))
	router.Get("/api/v1/wallet/{walletId}/history", walletActions.NewTransactionHistory(&s))
	router.Get("/api/v1/wallet/{walletId}", walletActions.NewGetWallet(&s))

	myLog.Info("Start server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HttpServerConfig.Timeout,
		WriteTimeout: cfg.HttpServerConfig.Timeout,
		IdleTimeout:  cfg.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		myLog.Error("Failed to start server")
	}
}

func SetUpLogger() *slog.Logger {
	return slog.New(
		slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
