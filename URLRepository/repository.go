package URLRepository

import "fmt"

type URLRepository interface {
	SaveURL(url, shortURL string) error
	GetURL(shortURL string) (string, error)
}

type InMemoryURLRepository struct {
	repo map[string]string
}

func NewInMemoryURLRepository() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		repo: make(map[string]string),
	}
}

func (ur *InMemoryURLRepository) SaveURL(url, shortURL string) error {
	if _, exists := ur.repo[shortURL]; exists {
		fmt.Println("short url already exists, hash collision")
		return fmt.Errorf("url already exists")
	}

	ur.repo[shortURL] = url
	return nil
}

func (ur *InMemoryURLRepository) GetURL(shortURL string) (string, error) {
	if _, exists := ur.repo[shortURL]; !exists {
		return "", fmt.Errorf("short url not found")
	}

	return ur.repo[shortURL], nil
}
