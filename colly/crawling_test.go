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
	Cookie    = "a=51krykFTJ7qI2tOlUmVbedSVcZw7h04c; _popprepop=1; zone-cap-3576187=1%3B1711166608; token_QpUJAAAAAAAAGu98Hdz1l_lcSZ2rY60Ajjk9U1c=BQPyAAAAAAAACZUAApG6kSTWc7nWCL1mRgpNrrpwkRwqeKbvFUYUxMJzzG6MticixdK-mNKN8oI3e4qbBcpXk25D5hX3TUsfV7zgEfcC3bVFvsJw6GvzhX-J8tY7GKL5lT4TQEUUNGzj6f0pAHmySNM3HV9qxfWv4uM6Uu8fZ4k4G2C9_Uy50gKWc-_R6X1_9pHd0P-QW9UbqtA4pGGzoav7iSfkQs7ChcGJ9GMJzZeCBduWzyl7hzLB0q0N0zTesiJLGU6wni3ZaVuatQB_HgBsOyo6ERmr42tdyDXs5rgclW85j1NxcvSxrlctJFpUUZx4ZnUZcQw3xVeNuAZntINdnt1b8pWbw1M6A9F3ruEqpNimATWfopLWDCp5RdaCmva6McAh22wBjBiQ6GjWoBvKJhYg5KXXqpb2ojBkSDo3Y37buZZzm6LFlmQmSLRK4enMuymD-0Z17xRQYeMJdFBJ9XnGlyq2SXuURLa2KgxrVyW165SrYByEuVgPRoYH-RNkNWqFnZ7PGUyW2bYUWOU8L4BVITCq1ADBqMGslh2v-yHx2LKXN58R7sahUqpRae6LaMH39bY8tQabsOCwnwiiYl1cXjWwtVSRdaYKTS7a1GcDonqLAn2DjpdqATNSJYTvoXlFR_H1S0XyPR86S14Hen3pRvPFgb_0oG2sSlogeALEYH94wMYIqWgwAUEDw5lUL4XU6O94NVwckkUZ_6Q0oXWJkhs4uFofJRm23PKLcBa0k6i05y2VuFKNQljg7hzJe28bQ9FqyWKdP8ZTeFHqlpv_-uwSYwyCstIJxWspJRzKpuYViOeFVPYD6naF3mz2_eeFSF-lfHkfeJKckh62yK4pdhFs72OBiPkw4YgKSZKnXI8jAP-5Z4UMI1qDc0BPgghRcUEd9cPRQV1WrJ0gkMHfE3epGozgt3iDMot33pSHYf4rGpeK-soxI-rZ90znNyj2nOg7L74QMj-FrFCSad7V0x2aScg013jSQH7Z3ymWTCpK0Haam5N3S2A1qe8LlPqK4x5xyyoR6LXycszhBjxb2vad5aFzeRnxuxvBy5EC-AuNTUSlR-7QIqFQytHUg9nHIyp_MI5xsvPimrrgyYqi1VJANpuBk3me5JFSg_D4nMo9axhnOewU073de1SL_EU3KFDjkhXmQAjAjdCI5T0ayeC9TVMkOJ9IEAtbC7IIJhVR3BThe55y_9vrPaX3Lo5pHQkMGik5YfUqCPEV-Tk_m2jCa-tF-k_yaK_3fWBGchLFzGDotuYOML8YAq-AQ6gFQMP-9H7VAyPG6u4QHaXiTZX2Y56DBhC7vYvZfT_q6-cwf6zQrB-9U1oLiUDfuHxrwLdcb8yFcw"
	UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36 Edg/122.0.0.0"
)

func TestCollyWoyaogexing(t *testing.T) {
	AllPageInfo()
}

func TestCollySt(t *testing.T) {
	proxyAllPageInfo()
}

func TestProxyDownload(t *testing.T) {
	ProxyDownload(Cookie, UserAgent)
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

func TestConnect(t *testing.T) {
	//m := make(map[string]interface{})
	//list := []interface{}{m}
	//
	db, err := Connect()
	if err != nil {
		fmt.Println(err)
	}

	var results []map[string]interface{}

	err = db.Table("table").Where("id = ?", 0).Find(&results).Error
	if err != nil {
		log.Println(err)
	}

}
