package main

import (
	"fmt"
	"tinyurl/URLRepository"
	"tinyurl/converter"
	"tinyurl/service"
)

func main() {
	converterObj := converter.NewHashConverter()

	urlRepo := URLRepository.NewInMemoryURLRepository()

	service := service.NewTinyURLServiceImpl(urlRepo, converterObj)

	url := "www.google.com"

	tinyURL, err := service.SaveURL(url)
	if err != nil {
		fmt.Printf("error saving url - %v\n", err)
		return
	}

	fmt.Printf("short url - %s\n", tinyURL)

	actualURL, err := service.GetURL(tinyURL)
	if err != nil {
		fmt.Printf("error getting url - %v\n", err)
		return
	}

	fmt.Printf("actualURL - %s", actualURL)
}
