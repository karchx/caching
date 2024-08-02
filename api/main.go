// https://github.com/Allah-The-Dev/caching-service/blob/master/caching-service-api/main.go
package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/karchx/api/config"
)

const (
	httpServerPort    = ":8000"
	readHeaderTimeout = 1 * time.Second
	writeTimeout      = 10 * time.Second
	idleTimeout       = 90 * time.Second
	maxHeaderBytes    = http.DefaultMaxHeaderBytes
)

func main() {
	c := config.NewConfig()
	c.InitializeAppConfig()
	router := initializeHTTPRouter()
	// HTTP server configuration
	httpServer := &http.Server{
		Addr:              httpServerPort,
		Handler:           router,
		ReadHeaderTimeout: readHeaderTimeout,
		WriteTimeout:      writeTimeout,
		IdleTimeout:       idleTimeout,
		MaxHeaderBytes:    maxHeaderBytes,
	}

	go func() {
		c.Logger().Infof("http server listening on port %s ", httpServerPort)

		err := httpServer.ListenAndServe()
		if err != nil {
			c.Logger().Infof("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	signal.Notify(ch, os.Kill)

	sig := <-ch
	c.Logger().Debugf("Got os signal: %v", sig)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	httpServer.Shutdown(ctx)

	log.Println("Sutting down server")
	os.Exit(0)
}

func initializeHTTPRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", articlesCategoryHandler)
	return r
}

func articlesCategoryHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Category %s\n", "jfdlkj")
}
