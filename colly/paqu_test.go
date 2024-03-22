package colly

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"testing"
	"time"
)

func TestPaqu(t *testing.T) {
	go GenerateUrlProducer("http://example.com")
	go GenerateWorkers()
	ResultsConsumer()
}

func TestPaqu2(t *testing.T) {
	testCollyDouban()
}

func TestCollyWoyaogexing(t *testing.T) {
	AllPageInfo()
}

func TestDownloadPic(t *testing.T) {
	DownloadPic()
}

func TestTime(t *testing.T) {
	for i := 0; i < 10; i++ {
		randNumber := time.Duration(rand.Intn(1000) + 1200)
		fmt.Println(randNumber)
		time.Sleep(randNumber * time.Millisecond)
	}
}

func TestCreateFile(t *testing.T) {
	resp, err := http.Get("https://img.dpm.org.cn/Uploads/Picture/2024/02/28/s65de7fdbd64e0.jpg")
	if err != nil {
		log.Println("Error fetching image:", err)
		return
	}
	defer resp.Body.Close()
	outFile, err := os.Create(fmt.Sprintf("F:/gugong_pic/%s", "dm"+".jpg"))
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
