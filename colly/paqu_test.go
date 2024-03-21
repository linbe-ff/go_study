package colly

import (
	"testing"
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
	testCollyWoyaogexing()
}
