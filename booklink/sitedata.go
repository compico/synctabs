package booklink

import "encoding/json"

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

func (sd *SiteData) ToJson() ([]byte, error) {
	data, err := json.MarshalIndent(sd, "", " ")
	if err != nil {
		return nil, err
	}
	return data, nil
}
