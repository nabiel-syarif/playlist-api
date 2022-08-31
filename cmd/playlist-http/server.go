package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/nabiel-syarif/playlist-api/internal/config"
)

type OnShutdownCallback func() error

func startServer(handler http.Handler, cfg *config.Config, callbacks []OnShutdownCallback) error {
	addr := fmt.Sprintf("%s:%s", cfg.HttpConfig.Addr, cfg.HttpConfig.Port)
	srv := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	errChan := make(chan error, 1)

	go func() {
		log.Printf("Server started on %s\n", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errChan <- err
		}
	}()

	go func() {
		// graceful shutdown
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

		<-signalChan

		ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*time.Duration(cfg.HttpConfig.GracefulTimeout))
		defer cancel()

		err := srv.Shutdown(ctx)
		if err != nil {
			errChan <- err
		}

		var wg sync.WaitGroup
		for _, v := range callbacks {
			wg.Add(1)
			go func(callback OnShutdownCallback) {
				defer wg.Done()
				err := callback()
				if err != nil {
					errChan <- err
					close(errChan)
					return
				}
			}(v)
		}

		wg.Wait()

		select {
		case <-ctx.Done():
			errChan <- ctx.Err()
			return
		default:
			log.Println("Server gracefully shutdown")
		}

		close(errChan)
	}()

	return <-errChan
}
