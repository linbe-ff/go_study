package colly

// 爬取
import (
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func testCollyDouban() {

	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*.douban.*", Parallelism: 1})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting11111111111", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong2222222222:", err)
	})

	c.OnHTML(".hd", func(e *colly.HTMLElement) {
		log.Println(strings.Split(e.ChildAttr("a", "href"), "/")[4],
			strings.TrimSpace(e.DOM.Find("span.title").Eq(2).Text()))
	})

	c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
	})

	c.Visit("https://movie.douban.com/top250?start=0&filter=")
	c.Wait()
}
