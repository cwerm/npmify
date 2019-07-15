package fetch

import (
	"io/ioutil"
	"log"
	"net/http"
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