package service

import (
	"fmt"
	"tinyurl/URLRepository"
	"tinyurl/converter"
)

type TinyURLService interface {
	SaveURL(url string) (string, error)
	GetURL(shortURL string) (string, error)
}

type TinyURLServiceImpl struct {
	urlRepo   URLRepository.URLRepository
	converter converter.Converter
}

func NewTinyURLServiceImpl(urlRepo URLRepository.URLRepository, converter converter.Converter) *TinyURLServiceImpl {
	return &TinyURLServiceImpl{
		urlRepo:   urlRepo,
		converter: converter,
	}
}

func (s *TinyURLServiceImpl) SaveURL(url string) (string, error) {
	encodedURL := s.converter.Encode(url)
	err := s.urlRepo.SaveURL(url, encodedURL)
	if err != nil {
		fmt.Println("error saving url - ", err)
	}
	return encodedURL, err
}

func (s *TinyURLServiceImpl) GetURL(shortURL string) (string, error) {
	url, err := s.urlRepo.GetURL(shortURL)
	if err != nil {
		fmt.Println("error getting url - ", err)
		return "", err
	}

	return url, nil
}
