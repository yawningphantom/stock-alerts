package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"stock-alerts/handlers"
	"time"
)

func main() {

	l := log.New(os.Stdout, "stock-alerts-api", log.LstdFlags)
	priceHandler := handlers.NewPriceHandler(l)

	sm := http.NewServeMux()
	sm.Handle("/price", priceHandler)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}

	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println(("Recieved terminate, graceful shutdown"), sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
