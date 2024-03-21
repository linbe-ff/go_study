package colly

// 爬取
import (
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"time"
)

func testCollyWoyaogexing() {
	url := "https://www.pexels.com/zh-cn/search/wallpaper/"
	cookieStr := "__cf_bm=A0m14BGt0xNbV3AMfFR2T1q6uKH7NkyiAPKfEXhpKaE-1711003121-1.0.1.1-oDoaKfDQLcyxXee2QVXU7U_Vc2iCBbx0vfxB2ZSSuO0cDZf2SVLwDPjUq3wtk6iVyVF400DyOChscNs8AQ9IXw; _sp_ses.9ec1=*; country-code-v2=US; _ga=GA1.1.1908980835.1711003130; cf_clearance=DaBd.niwavn_ntdv22xxhnkVOyTGOHV.O24gOns.kNQ-1711003128-1.0.1.1-hRD7r6A5_4tDl1pMRfxrbst97UOtexUckU9tfit5eo6FWALNeIGrJ8dBkLBq9lLbs0eRwIjWOrv1UzD0pJI_iQ; _fbp=fb.1.1711003136439.433319942; OptanonConsent=isGpcEnabled=0&datestamp=Thu+Mar+21+2024+14%3A41%3A46+GMT%2B0800+(%E4%B8%AD%E5%9B%BD%E6%A0%87%E5%87%86%E6%97%B6%E9%97%B4)&version=202301.1.0&isIABGlobal=false&hosts=&landingPath=NotLandingPage&groups=C0001%3A1%2CC0002%3A0%2CC0003%3A0%2CC0004%3A0&AwaitingReconsent=false; _sp_id.9ec1=eee1ba33-5d87-4600-b563-3057fa36f7bd.1711003129.1.1711003307..9d381904-e4cd-4e0d-ab27-2fba4217bbb7..67559036-6ecb-46d7-9194-0a6b9b1a38bf.1711003128869.6; _ga_8JE65Q40S6=GS1.1.1711003129.1.1.1711003307.0.0.0"
	c := colly.NewCollector(
		colly.Async(false),
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0"),
	)

	// 添加一个具体的cookie，例如：
	cookie := &http.Cookie{
		Name:    "cookie",
		Value:   cookieStr,
		Domain:  url,                            // 需要设置合适的域名
		Path:    "/",                            // 设置路径，默认为"/"
		Expires: time.Now().Add(24 * time.Hour), // 设置过期时间
	}
	c.SetCookies(url, []*http.Cookie{cookie})

	c.Limit(&colly.LimitRule{DomainGlob: "*.woyaogexing.*", Parallelism: 1})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting11111111111", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong2222222222:", err)
	})

	c.OnHTML(".Link_link__mTUkz spacing_noMargin__Q_PsJ", func(e *colly.HTMLElement) {
		log.Println(e.ChildAttr("a", "href"))
	})

	//c.OnHTML(".paginator a", func(e *colly.HTMLElement) {
	//	e.Request.Visit(e.Attr("href"))
	//})

	c.Visit(url)
	c.Wait()
}
