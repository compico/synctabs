package booklink_test

import (
	"testing"

	bl "github.com/compico/synctabs/booklink"
)

func TestToJson(t *testing.T) {
	data := bl.SiteData{
		Icon:        "Test",
		Name:        "Test",
		Title:       "Test",
		Description: "Test",
		Images:      []string{"Test","Test"},
		Url:         "Test",
	}
	r, err := bl.ToJson(data)
	if err != nil {
		t.Errorf("Marshall error: %v\n", err.Error())
	}
	t.Log(r)
}