package colly

import (
	"context"
	"fmt"
	. "go_study/mysql"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

var (
	Cookie    = "__uvt=a%3A1%3A%7Bi%3A0%3Bs%3A33%3A%2266599f72461fb8.934030602872256642%22%3B%7D"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/125.0.0.0 Safari/537.36 Edg/125.0.0.0"
)

func TestCollyWoyaogexing(t *testing.T) {
	AllPageInfo()
}

func TestCollySt(t *testing.T) {
	proxyAllPageInfo()
}

func TestCollyStP18(t *testing.T) {
	proxyAllPageInfoP18()
}

func TestProxyDownload(t *testing.T) {
	ProxyDownload(Cookie, UserAgent)
}

func TestProxyDownloadForMysql(t *testing.T) {
	ProxyDownloadForMysql(Cookie, UserAgent)
}

func TestProxyDownloadForMysqlP18(t *testing.T) {
	ProxyDownloadForMysqlP18(Cookie, UserAgent)
}

func TestDownload(t *testing.T) {
	Download()
}

func TestTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		randNumber := time.Duration(rand.Intn(1000) + 1200)
		fmt.Println(randNumber)
		time.Sleep(randNumber * time.Millisecond)
	}
}

func TestCreateFile(t *testing.T) {
	resp, err := http.Get("https://img..org.cn/Uploads/ture/2024/02/28/.jpg")
	if err != nil {
		log.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()
	outFile, err := os.Create(fmt.Sprintf("F:/gugong_/%s", "dm"+".jpg"))
	if err != nil {
		log.Println("Error creating file:", err)
		return
	}
	_, err = io.Copy(outFile, resp.Body)
	if err != nil {
		log.Println("Error writing to file:", err)
		return
	}

	log.Printf("Successfully downloaded image: ")
}

func TestReadExcel(t *testing.T) {
	Read(".xlsx", "")
}

func TestCreateExcel(t *testing.T) {
	list := make(PageInfoList, 0)
	list = append(list, &PageInfo{Page: 1, Title: "1", Href: "1", Src: "1"})
	list = append(list, &PageInfo{Page: 2, Title: "2", Href: "2", Src: "2"})
	list = append(list, &PageInfo{Page: 3, Title: "3", Href: "3", Src: "3"})
	CreateExcel(list, ".xlsx", "")
}

func TestExcelAppend(t *testing.T) {
	list := make(PageInfoList, 0)
	list = append(list, &PageInfo{Page: 4, Title: "4", Href: "4", Src: "4"})
	list = append(list, &PageInfo{Page: 5, Title: "5", Href: "5", Src: "5"})
	list = append(list, &PageInfo{Page: 6, Title: "6", Href: "6", Src: "6"})
	ExcelAppend(list, ".xlsx", "")
}

func TestGo(t *testing.T) {
	for i := 0; i < 10; i++ {
		go printI(i)
	}

}

func TestCreate(t *testing.T) {
	var (
		listLen int
		count   int
		isBreak bool
	)
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}
	dao := NewPageInfoDao(db)
	list, err := Read(".xlsx", "")
	for _, info := range list {
		//info.ID = uint(snowflack.NewCustomNode().GenerateID().Int64())
		info.InfoType = "wz"
	}
	// 分批更新(1000条一次)
	listLen = len(list)
	for {
		if count+1000 > listLen {
			dao.CreatePageInfoList(context.TODO(), list[count:])
			isBreak = true
		} else {
			dao.CreatePageInfoList(context.TODO(), list[count:count+1000])
			count += 1000
		}
		if isBreak {
			break
		}
	}

}
func printI(i int) {
	time.Sleep(time.Second)
	print(i)
}
