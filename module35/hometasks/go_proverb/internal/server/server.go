// Package server содержит пакет для работы с сервером.
package server

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"sync"
	"time"
)

const (
	// tickerDuration период обновления цитаты.
	tickerDuration = 3 * time.Second
	// addr - сетевой адрес.
	addr = "0.0.0.0:12345"
	// protocol - протокол сетевой службы.
	protocol = "tcp4"
)

// TCPServer представляет TCP сервер
type TCPServer struct {
	proverbs []string
	listener net.Listener
	wg       sync.WaitGroup
	ctx      context.Context
	cancel   context.CancelFunc
}

// NewTCPServer создает новый TCPServer.
func NewTCPServer(proverbs []string) (*TCPServer, error) {
	lis, err := net.Listen(protocol, addr)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())

	return &TCPServer{
		proverbs: proverbs,
		listener: lis,
		ctx:      ctx,
		cancel:   cancel,
	}, nil
}

// Start запускает TCP-сервер.
func (s *TCPServer) Start() {
	log.Printf("Server start on %s %s", protocol, addr)

	go func() {
		for {
			conn, err := s.listener.Accept()
			if err != nil {
				select {
				case <-s.ctx.Done():
					log.Println("Server shutting down")
					return
				default:
					log.Printf("Failed to accept connection: %v", err)
					continue
				}
			}

			s.wg.Add(1)

			go s.handleConn(conn)
		}
	}()
}

// handleConn обрабатывает отдельные соединения с клиентами.
func (s *TCPServer) handleConn(conn net.Conn) {
	defer func() {
		log.Printf("Closing connection for client: %v", conn.RemoteAddr())
		conn.Close()
		s.wg.Done()
	}()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-s.ctx.Done():
			return
		case <-ticker.C:
			proverb := s.proverbs[r.Intn(len(s.proverbs))]

			if _, err := fmt.Fprintf(conn, "%s\r\n", proverb); err != nil {
				log.Printf("error write to client %s: %v", conn.RemoteAddr(), err)
				return
			}
		}
	}
}

// Shutdown останавливает TCP-сервер.
func (s *TCPServer) Shutdown() {
	s.cancel()
	s.listener.Close()
	s.wg.Wait()
	log.Println("Server shutdown complete")
}
