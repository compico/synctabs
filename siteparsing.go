package main

import (
	"encoding/json"

	"github.com/badoux/goscraper"
)

type SiteData struct {
	Icon        string   `json:"icon"`
	Name        string   `json:"name"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Images      []string `json:"images"`
	Url         string   `json:"url"`
}

func getSiteData(url string) (*SiteData, error) {
	data := new(SiteData)

	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		return nil, err
	}

	data = &SiteData{
		Icon:        s.Preview.Icon,
		Name:        s.Preview.Name,
		Title:       s.Preview.Title,
		Description: s.Preview.Description,
		Url:         s.Preview.Link,
	}

	if len(s.Preview.Images) > 0 {
		data.Images = s.Preview.Images
	}
	return data, nil
}

func (sd *SiteData) toJSON() (string, error) {
	data, err := json.MarshalIndent(sd, "", " ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}
