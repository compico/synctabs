package booklink

import (
	"encoding/json"
	"io/ioutil"
)

type (
	SiteData struct {
		Icon        string   `json:"icon"`
		Name        string   `json:"name"`
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Images      []string `json:"images"`
		Url         string   `json:"url"`
	}
)

func newSiteData() *SiteData {
	return new(SiteData)
}

func (sd *SiteData) ToJson() ([]byte, error) {
	data, err := json.Marshal(sd)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (sd *SiteData) FromJson(path string) error {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	err = json.Unmarshal(bytes, sd)
	if err != nil {
		return err
	}
	return nil
}
