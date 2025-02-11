package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/ElizaDeani/archivio/config"
	"github.com/ElizaDeani/archivio/internal/builder"
	"github.com/ElizaDeani/archivio/pkg/database"
	"github.com/ElizaDeani/archivio/pkg/server"
)

func main() {
	cfg, err := config.NewConfig(".env")
	checkError(err)

	db, err := database.InitDatabase(cfg.MySQLConfig)
	checkError(err)
	fmt.Println("database connected")

	publicRoutes := builder.BuildPublicRoutes(cfg, db)
	privateRoutes := builder.BuildPrivateRoutes(cfg, db)

	//init server
	srv := server.NewServer(cfg, publicRoutes, privateRoutes)

	runServer(srv, cfg.PORT)
	waitForShutdown(srv)

}

func waitForShutdown(srv *server.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Gagal mematikan server: %v", err)
	}
	log.Println("Server berhasil dimatikan.")
}

func runServer(srv *server.Server, port string) {
	go func() {
		err := srv.Start(fmt.Sprintf(":%s", port))
		log.Fatal(err)
	}()
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
