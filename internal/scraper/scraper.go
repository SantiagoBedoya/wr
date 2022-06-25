package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

const url = "https://www.wordreference.com"

// WrTranslate visit wordReferece´s page and scrap the content about the translation
func WrTranslate(word, from, to string) {
	c := colly.NewCollector(
		colly.MaxDepth(2),
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})
	results := make([]Result, 0)
	c.OnHTML("tr.even", func(e *colly.HTMLElement) {
		from := e.DOM.ChildrenFiltered("td.FrWrd")
		to := e.DOM.ChildrenFiltered("td.ToWrd")
		result := Result{
			FromWord: from.Text(),
		}
		// fmt.Println(from.Text())
		if from.Text() == "" {
			result.ToWord = to.Text()
			// fmt.Println("\t", to.Text())
		} else {

		}
		results = append(results, result)
		// if from.Text() == "" {
		// 	fmt.Println("\t" + to.Text())
		// } else {
		// 	fmt.Println(from.Text())
		// }
		// e.DOM.Map(func(i int, s *goquery.Selection) string {
		// 	fmt.Println(s.Text())

		// 	fmt.Println(strings.HasPrefix(s.Text(), "  "))
		// 	fmt.Println("........")
		// 	return s.Text()
		// })
		// e.ForEach("td.FrWrd > strong, td.ToWrd", func(i int, h *colly.HTMLElement) {

		// })
	})
	// c.OnHTML("tr.even > td.FrWrd", func(e *colly.HTMLElement) {
	// 	fmt.Println(fmt.Sprintf("FromWorkd [index=%d, text=%s]", e.Index, e.Text))
	// })
	// c.OnHTML("tr.even > td.ToWrd", func(e *colly.HTMLElement) {
	// 	fmt.Println(fmt.Sprintf("ToWorkd [index=%d, text=%s]", e.Index, e.Text))
	// })
	c.OnScraped(func(r *colly.Response) {
		for _, v := range results {
			fmt.Printf("%s -> %s\n", v.FromWord, v.ToWord)
		}
	})
	completeURL := fmt.Sprintf("%s/%s%s/%s", url, from, to, word)
	c.Visit(completeURL)
	c.Wait()
}
