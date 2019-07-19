package fetch

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Get(url string) []uint8 {
	res, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return body
}

func AsyncGet(url string, ch chan<-string) {
	start := time.Now()
	resp, _ := http.Get(url)

	secs := time.Since(start).Seconds()
	body, _ := ioutil.ReadAll(resp.Body)
	ch <- fmt.Sprintf("%.2f elapsed with response length: %d %s", secs, len(body), url)
}