package scraper

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

const (
	url = "https://www.wordreference.com"
)

// WrTranslate visit wordReferece´s page and scrap the content about the translation
func WrTranslate(word, from, to string) {
	c := colly.NewCollector()

	results := make(map[string]struct{}, 0)

	c.OnHTML("tr.even > td.ToWrd, tr.odd > td.ToWrd", func(e *colly.HTMLElement) {
		if len(results) < 3 {
			results[sanitize(e.Text)] = struct{}{}
		}
	})

	c.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished!")
		keys := make([]string, 0, len(results))
		for k := range results {
			keys = append(keys, k)
		}
		fmt.Println(strings.Join(keys, ", "))
	})
	completeURL := fmt.Sprintf("%s/%s%s/%s", url, from, to, word)
	c.Visit(completeURL)
}

func sanitize(text string) string {
	var words = []string{"nm", "+", "adj", "prnl", "loc", "adv", "⇒", "vtr", "verb", "prep", "mf", "nf", "propio", " n", " f", " vi", " v", " expr"}
	for _, v := range words {
		text = strings.ReplaceAll(text, v, "")
	}
	text = strings.TrimSpace(text)
	return text
}
