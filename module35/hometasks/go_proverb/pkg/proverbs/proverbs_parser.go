// Package proverbs содержит пакет получения поговорок с сайта.
package proverbs

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
)

// ParseGoProverbsPage считывает поговорки Go со страницы и возвращает их в виде слайса строк.
func ParseGoProverbsPage(url string) ([]string, error) {
	res, err := http.Get(url)
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
