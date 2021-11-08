package booklink

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/url"
	"strings"
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
	Bookmark struct {
		Sitelinks []SiteData
		Name      string
		Icon      string
		ItsFolder bool
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

func (sd *SiteData) FromUrl(_url string) error {
	var err error
	sd.Url, err = vurl(_url)
	if err != nil {
		return err
	}
	return nil
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

func vurl(_url string) (string, error) {
	u, err := url.Parse(_url)
	if err != nil {
		return "", err
	}
	u.Scheme = strings.ToLower(u.Scheme)
	var qurl string
	if u.Scheme == "" {
		return "", errors.New("Scheme empty!")
	}
	if u.Host == "" {
		return "", errors.New("Empty host!")
	}
	qurl = u.Scheme + "://" + u.Host
	if u.Path != "" || u.Path == "/" {
		qurl += "/"
	}
	if u.RawQuery != "" {
		qurl += "?" + u.RawQuery
	}
	if u.Fragment != "" {
		qurl += "#" + u.Fragment
	}
	return qurl, nil
}
