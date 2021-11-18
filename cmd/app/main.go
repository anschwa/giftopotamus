package main

import (
	"context"
	"encoding/json"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/anschwa/giftopotamus/handlers"
	"github.com/anschwa/giftopotamus/logger"
	"github.com/anschwa/giftopotamus/middleware"
)

const Version = "0.0.1"

var Build = "" // Set at build-time with: go run -ldflags "-X main.Build=abc123"

var (
	port       = ""  // Heroku port
	listenAddr = ""  // Host and port
	healthy    int32 // 0 is dead, nonzero is healthy
	appEnv     string

	authDB = &middleware.AuthDB{}
)

const (
	ioTimeout = 60 * time.Second

	// Limit form submissions to 5 requests per minute (allow 1 request every N seconds)
	rateLimitInterval = (1 * time.Minute) / 5
)

func init() {
	logger.Init(os.Stdout)

	switch os.Getenv("APP_ENV") {
	case "production":
		appEnv = "production"
	default:
		appEnv = "development"
	}

	if port := os.Getenv("PORT"); port != "" {
		listenAddr = ":" + port
	} else {
		flag.StringVar(&listenAddr, "http", ":8080", "Server listen address")
		flag.Parse()
	}

	if db := os.Getenv("AUTH_DB_JSON"); db != "" {
		var entries map[string]string
		if err := json.Unmarshal([]byte(db), &entries); err != nil {
			logger.Error("Error unmarshaling AUTH_DB_JSON:", err)
			os.Exit(1)
		}

		authDB = middleware.NewAuthDB(entries)
	}
}

func main() {
	sm := middleware.NewSessionManager("giftexsession", 0)

	// Router
	r := http.NewServeMux()
	r.Handle("/healthz", healthz()) // Healthcheck
	r.Handle("/public/", public())  // Static files

	r.Handle("/", handlers.Index(sm))
	r.Handle("/login", handlers.Login(sm, authDB))
	r.Handle("/logout", handlers.Logout(sm))

	r.Handle("/import", handlers.ImportGiftExchange(sm))
	r.Handle("/edit", handlers.EditGiftExchange(sm))
	r.Handle("/create", handlers.CreateGiftExchange(sm))
	r.Handle("/download", handlers.DownloadGiftExchange(sm))
	r.Handle("/sendmail", handlers.SendGiftExchange(sm))

	// Set request middleware
	handler :=
		middleware.RequestID(
			middleware.Logger(
				middleware.Recoverer(r)))

	srv := &http.Server{
		Addr:         listenAddr,
		Handler:      handler,
		ReadTimeout:  ioTimeout,
		WriteTimeout: ioTimeout,
	}

	waitForShutdown := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		logger.Info("Shutting down...")
		atomic.StoreInt32(&healthy, 0)

		srv.SetKeepAlivesEnabled(false)
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			logger.Fatalf("Error shutting down server: %v", err)
		}

		close(waitForShutdown) // Successful shutdown
	}()

	logger.Infof("Version: %q; Build: %q; Env: %q", Version, Build, appEnv)
	logger.Infof("Starting server on: %s", srv.Addr)
	atomic.StoreInt32(&healthy, 42)

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Error starting server: %v", err)
	}

	<-waitForShutdown
	logger.Info("Goodbye.")
}

var fileServer = http.FileServer(http.Dir("public"))

func public() http.Handler {
	return http.StripPrefix("/public/", fileServer)
}

// healthz is a naming convention from google/k8s
func healthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 0 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
