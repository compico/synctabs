package booklink_test

import (
	"testing"

	bl "github.com/compico/synctabs/booklink"
)

func TestFromJson(t *testing.T) {
	var sd bl.SiteData
	err := sd.FromJson("data_for_tests/sitedata.json")
	if err != nil {
		t.Errorf("Unmarshel error : %v", err.Error())
	}
}