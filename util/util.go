package util

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func Api(apiUrl string) []byte {

	_, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		panic(err)
	}

	response, err := http.Get(apiUrl)

	if err != nil {
		log.Fatal(err)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return responseData
}
