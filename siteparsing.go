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
	Error       string   `json:"error"`
}

func getSiteData(url string) (*SiteData, error) {

	data := new(SiteData)

	s, err := goscraper.Scrape(url, 5)
	if err != nil {
		data.Error = err.Error()
		return data, err
	}

	data.Icon = s.Preview.Icon
	data.Name = s.Preview.Name
	data.Title = s.Preview.Title
	data.Description = s.Preview.Description
	data.Url = s.Preview.Link

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
