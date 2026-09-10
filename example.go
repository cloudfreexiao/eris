package main

import (
	"log/slog"
	"os"
	"os/signal"

	"github.com/cloudfreexiao/eris/cluster"
	"github.com/cloudfreexiao/eris/service"
)

func main() {
	// initialize services
	httpService := service.NewHttpService()
	pingService := service.NewPingService()

	// initialize cluster
	clusterd := cluster.GetClusterd()
	clusterd.Reload(cluster.DefaultConfig{
		"moon": "0.0.0.0:3345",
	})

	// register services
	clusterd.Register("http", httpService)
	clusterd.Register("ping", pingService)

	// start cluster
	clusterd.Open("moon")

	slog.Info("moon started")

	term := make(chan os.Signal, 1)

	signal.Notify(term, os.Interrupt)

	<-term
}
