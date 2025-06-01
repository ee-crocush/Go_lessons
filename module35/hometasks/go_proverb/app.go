package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"
)

// parseURL страница с Go proverbs.
const parseURL = "https://go-proverbs.github.io/"

// tickerDuration период обновления цитаты.
const tickerDuration = 3 * time.Second

// addr - сетевой адрес.
const addr = "0.0.0.0:12345"

// protocol - протокол сетевой службы.
const protocol = "tcp4"

func main() {
	// Парсим поговорки один раз при старте
	proverbs, err := parseGoProverbsPage()
	if err != nil {
		log.Fatalf("ошибка при загрузке поговорок: %v", err)
	}

	// Слушаем TCP-порт
	listener, err := net.Listen(protocol, addr)
	if err != nil {
		log.Fatalf("ошибка при прослушивании порта: %v", err)
	}
	defer listener.Close()

	fmt.Printf("Сервер запущен на %s %s\n", protocol, addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("ошибка соединения: %v", err)
			continue
		}
		fmt.Printf("Новое подключение: %s\n", conn.RemoteAddr())

		// Обрабатываем каждого клиента в отдельной горутине
		go handleConn(conn, proverbs)
	}
}

// handleConn Обработчик соединения
func handleConn(conn net.Conn, proverbs []string) {
	defer func() {
		fmt.Printf("Отключение клиента: %s\n", conn.RemoteAddr())
		conn.Close()
	}()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ticker := time.NewTicker(tickerDuration)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			proverb := proverbs[r.Intn(len(proverbs))]
			_, err := fmt.Fprintf(conn, "%s\r\n", proverb)
			if err != nil {
				log.Printf("ошибка записи клиенту %s: %v", conn.RemoteAddr(), err)
				return
			}
		}
	}
}

func parseGoProverbsPage() ([]string, error) {
	res, err := http.Get(parseURL)
	if err != nil {
		return nil, fmt.Errorf("ошибка HTTP-запроса: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("неожиданный код ответа: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка парсинга HTML: %w", err)
	}

	var proverbs []string
	doc.Find("h3 a").Each(
		func(i int, s *goquery.Selection) {
			content := s.Text()
			proverbs = append(proverbs, content)
		},
	)

	if len(proverbs) == 0 {
		return nil, fmt.Errorf("поговорки не найдены")
	}
	return proverbs, nil
}
