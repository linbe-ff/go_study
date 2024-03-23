package colly

// 爬取
import (
	"bytes"
	"fmt"
	"github.com/WooRho/rhtool/rhtool_core/rexcel"
	snowflack "github.com/WooRho/rhtool/rhtool_snowflack"
	"github.com/gocolly/colly"
	. "go_study/mysql"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func AllPageInfo() {
	var (
		pageInfoList = make(PageInfoList, 0)
		buffer       = &bytes.Buffer{}
	)
	baseInfo := Init("acw_tc=3ccdc14317110130576548934e6a617df6034d445d08e68d0ca9770882fcce; PHPSESSID=klct3h61l6sn29qlknqqr7i4t6; saw_terminal=default; UM_distinctid=18e60535a5127c-0992569951295-4c657b58-1fa400-18e60535a52a57; CNZZDATA1261553859=862871121-1711013059-https%253A%252F%252Flink.zhihu.com%252F%7C1711013059; _abfpc=7ca9a41165e72934d0f58094a6439fa41a64c1c4_2.0; cna=dcd62a1926c3e8901691942d3b8ae25d",
		"https://www..org.cn/lights/royal/p/",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	)
	preUrl := baseInfo.Url
	for i := 1; i <= 141; i++ {

		baseInfo.Url = preUrl + strconv.Itoa(i) + ".html"
		baseInfo.GetPageOne(&pageInfoList)
		randNumber := time.Duration(rand.Intn(1000) + 1200)
		time.Sleep(randNumber * time.Millisecond)
	}

	err := rexcel.BufferToExcel(pageInfoList, buffer, "", "")
	if err != nil {
		log.Println(err)
	}
}

func (l *BaseInfo) GetPageOne(pageInfoList *PageInfoList) {
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

	c.OnHTML(".", func(e *colly.HTMLElement) {
		ALabel := e.DOM.Find("a")
		labelHrefAddr, _ := ALabel.Attr("href")
		Label := ALabel.Find("img")
		Titile, _ := Label.Attr("title")
		Src, _ := Label.Attr("src")
		*pageInfoList = append(*pageInfoList, &PageInfo{
			Title: Titile,
			Href:  labelHrefAddr,
			Src:   Src,
		})
	})

	c.Visit(l.Url)
	c.Wait()
}

func Download() {
	var (
		pageInfoList = make(PageInfoList, 0)
		PreAddrConst = "https://www..org.cn/light/"
	)
	// 自定义文件
	_data, err := rexcel.LoadFromExcelFile(
		".xlsx",
		PageInfo{},
		"",
	)
	if err != nil {
		return
	}

	for _, item := range _data {
		info := item.(PageInfo)
		pageInfoList = append(pageInfoList, &info)
	}

	baseInfo := Init("acw_tc=3ccdc14317110130576548934e6a617df6034d445d08e68d0ca9770882fcce; PHPSESSID=klct3h61l6sn29qlknqqr7i4t6; saw_terminal=default; UM_distinctid=18e60535a5127c-0992569951295-4c657b58-1fa400-18e60535a52a57; CNZZDATA1261553859=862871121-1711013059-https%253A%252F%252Flink.zhihu.com%252F%7C1711013059; _abfpc=7ca9a41165e72934d0f58094a6439fa41a64c1c4_2.0; cna=dcd62a1926c3e8901691942d3b8ae25d",
		"https://www..org.cn/light/",
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0",
	)
	// 下载操作
	for _, info := range pageInfoList {
		baseInfo.Url = PreAddrConst + strings.SplitAfter(info.Href, "/")[2]
		log.Println(baseInfo.Url)
		baseInfo.Save(info)
		randNumber := time.Duration(rand.Intn(1000))
		time.Sleep(randNumber * time.Millisecond)
	}
}

func (l *BaseInfo) Save(info *PageInfo) {
	preAddr := "https://img..org.cn/"
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

	c.OnHTML(".tureshow", func(e *colly.HTMLElement) {
		imgLabel := e.DOM.Find("img")
		//log.Println(imgLabel)
		srcAddr, _ := imgLabel.Attr("src")
		//log.Println(srcAddr)

		resp, err := http.Get(preAddr + srcAddr)
		if err != nil {
			log.Println("Error fetching image:", err)
			return
		}
		defer resp.Body.Close()
		// 保存图片到本地
		fileName := info.Title + strconv.Itoa(int(snowflack.NewCustomNode().GenerateID().Int64())) + ".jpg"
		log.Printf(fileName)
		outFile, err := os.Create(fmt.Sprintf("F:/gugong_/%s", fileName))
		if err != nil {
			log.Println("Error creating file:", err)
			return
		}
		defer outFile.Close()
		_, err = io.Copy(outFile, resp.Body)
		if err != nil {
			log.Println("Error writing to file:", err)
			return
		}

	})

	c.Visit(l.Url)
	c.Wait()
}
