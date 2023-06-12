package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "1")
	})

	srv := http.Server{
		Addr:    ":80",
		Handler: handler,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()

		if err != nil && err != http.ErrServerClosed {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-quit

	if err := srv.Shutdown(nil); err != nil {
		log.Println("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
