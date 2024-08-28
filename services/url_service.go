package services

import (
	"strings"
	"travelgo/models"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-resty/resty/v2"
)

type UrlService struct{}

func NewUrlService() UrlService {
	return UrlService{}
}

func (us *UrlService) FetchMetadata(url string) (*models.Metadata, error) {
	client := resty.New()
	resp, err := client.R().Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(resp.String()))
	if err != nil {
		return nil, err
	}

	metadata := models.Metadata{
		Title:       getMetaContent(doc, "title", "meta[property='og:title']", ""),
		Description: getMetaContent(doc, "", "meta[name='description']", "meta[property='og:description']"),
	}

	return &metadata, nil
}

func getMetaContent(doc *goquery.Document, titleSelector string, metaSelectors ...string) string {
	if titleSelector != "" {
		title := doc.Find(titleSelector).Text()
		if title != "" {
			return title
		}
	}

	for _, selector := range metaSelectors {
		content, exists := doc.Find(selector).Attr("content")
		if exists {
			return content
		}
	}

	return ""
}
