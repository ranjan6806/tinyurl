package main

import (
	"fmt"
	"testing"
	"tinyurl/URLRepository"
	"tinyurl/converter"
	"tinyurl/service"
)

func Test_service(t *testing.T) {
	converterObj := converter.NewHashConverter()

	urlRepo := URLRepository.NewInMemoryURLRepository()

	serviceObj := service.NewTinyURLServiceImpl(urlRepo, converterObj)

	url := "www.google.com"

	tinyURL, err := serviceObj.SaveURL(url)
	if err != nil {
		t.Errorf("error saving url - %v\n", err)
		return
	}

	fmt.Printf("short url - %s\n", tinyURL)

	actualURL, err := serviceObj.GetURL(tinyURL)
	if err != nil {
		t.Errorf("error getting url - %v\n", err)
		return
	}

	fmt.Printf("actualURL - %s", actualURL)

	if actualURL != url {
		t.Errorf("actualURL - %s != %s", actualURL, url)
	}
}
