functionalities

SaveURL(url string) string
GetURL(shortURL string) string

type Converter interface {
    Encode(url string) string
}

type HashConverted struct {
    Encode(url string) string
}

type URLRepo interface{
    SaveURL(url string) string
    GetURL(shortURL string) string
}

type InMemoryURLRepo struct {
    urlMap map[string][string]
    SaveURL(url string) string
    GetURL(shortURL string) string
}

type Service interface {
    SaveURL(url string) string
    GetURL(shortURL string) string 
}

type ServiceImpl struct {
    converter Converter
    repo URLRepo
    SaveURL(url string) string
    GetURL(shortURL string) string
}