package booklink_test

import (
	"testing"

	bl "github.com/compico/synctabs/booklink"
)
var data = &bl.SiteData{
	Icon:        "Test",
	Name:        "Test",
	Title:       "Test",
	Description: "Test",
	Images:      []string{"Test","Test"},
	Url:         "Test",
}

func TestToJson(t *testing.T) {
	want := "{\"icon\":\"Test\",\"name\":\"Test\",\"title\":\"Test\",\"description\":\"Test\",\"images\":[\"Test\",\"Test\"],\"url\":\"Test\"}"
	result, err := bl.ToJson(data)
	if err != nil {
		t.Errorf("Marshall error: %v\n", err.Error())
	}
	if string(result) == want {
		t.Logf("\nresult\t=\t%v\nwant\t=\t%v\nresult and want equal\n",string(result), want)
	}
}