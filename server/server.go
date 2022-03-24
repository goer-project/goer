package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func Run(router *gin.Engine, port uint) {

	// http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: router,
	}

	// Listen and Server
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// log
	fmt.Printf("\r\n")
	fmt.Println("Server run at:")
	fmt.Printf("- Local: http://localhost:%d/ \r\n", port)
	fmt.Printf("\r\n")
	log.Printf("Enter Control + C Shutdown Server \r\n")

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// log
	fmt.Printf("\r\n")
	log.Println("Shutdown Server ...")

	// Timeout of 5 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Shutdown
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// log
	log.Println("Server exiting")
}
