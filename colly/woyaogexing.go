package colly

// 爬取
import (
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"time"
)

type (
	PageInfo struct {
		Title string
		Href  string
		Src   string
	}
	BaseInfo struct {
		Cookie    string
		Url       string
		userAgent string
	}
	PageInfoList []*PageInfo
)

func Init(cookie, url, userAgent string) *BaseInfo {
	return &BaseInfo{
		Cookie:    cookie,
		Url:       url,
		userAgent: userAgent,
	}
}

func (l *BaseInfo) InitCookie() *http.Cookie {
	return &http.Cookie{
		Name:    "cookie",
		Value:   l.Cookie,
		Domain:  l.Url,                          // 需要设置合适的域名
		Path:    "/",                            // 设置路径，默认为"/"
		Expires: time.Now().Add(24 * time.Hour), // 设置过期时间
	}
}

func AllPageInfo() {
	var (
		pageInfoList = make(PageInfoList, 0)
	)
	baseInfo := Init("acw_tc=3ccdc14317110130576548934e6a617df6034d445d08e68d0ca9770882fcce; PHPSESSID=klct3h61l6sn29qlknqqr7i4t6; saw_terminal=default; UM_distinctid=18e60535a5127c-0992569951295-4c657b58-1fa400-18e60535a52a57; CNZZDATA1261553859=862871121-1711013059-https%253A%252F%252Flink.zhihu.com%252F%7C1711013059; _abfpc=7ca9a41165e72934d0f58094a6439fa41a64c1c4_2.0; cna=dcd62a1926c3e8901691942d3b8ae25d",
		"https://www.dpm.org.cn/lights/royal.html",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	)
	baseInfo.GetPageOne(pageInfoList)
}

func (l *BaseInfo) GetPageOne(pageInfoList PageInfoList) {

	c := colly.NewCollector(
		colly.Async(false),
		colly.UserAgent(l.userAgent),
	)
	c.SetCookies(l.Url, []*http.Cookie{l.InitCookie()})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnHTML(".pic", func(e *colly.HTMLElement) {
		//log.Println(e.ChildAttr("a", "href"),
		//	strings.TrimSpace(e.DOM.Find("img").Eq(0).Text()),
		//)
		A标签 := e.DOM.Find("a")
		标签指向地址, _ := A标签.Attr("href")
		图片标签 := A标签.Find("img")
		// 输出biaoqianImg的src
		图片标题, _ := 图片标签.Attr("title")
		图片来源, _ := 图片标签.Attr("src")
		pageInfoList = append(pageInfoList, &PageInfo{
			Title: 图片标题,
			Href:  标签指向地址,
			Src:   图片来源,
		})
	})

	c.Visit(l.Url)
	c.Wait()
}
