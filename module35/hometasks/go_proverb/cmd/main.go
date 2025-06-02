package main

import (
	httpserver "example.com/go_proverb/internal/server"
	pvb "example.com/go_proverb/pkg/proverbs"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// parseURL страница с Go proverbs.
const parseURL = "https://go-proverbs.github.io/"

func main() {
	// Парсим поговорки один раз при старте
	proverbs, err := pvb.ParseGoProverbsPage(parseURL)
	if err != nil {
		log.Fatalf("ошибка при загрузке поговорок: %v", err)
	}

	server, err := httpserver.NewTCPServer(proverbs)
	if err != nil {
		log.Fatalf("не удалось запустить сервер: %v", err)
	}

	// Graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	server.Start()

	<-sigCh
	server.Shutdown()
}
