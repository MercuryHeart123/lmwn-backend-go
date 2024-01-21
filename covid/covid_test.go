package covid_test

import (
	"encoding/json"
	"testing"

	"github.com/mercuryheart123/lmwn-backend-go/covid"
	"github.com/mercuryheart123/lmwn-backend-go/model"
)

func TestGetSummarize(t *testing.T) {
	mockJsonData := []byte(`[
		{
			"ConfirmDate": "2021-05-01",
			"No": null,
			"Age": null,
			"Gender": "หญิง",
			"GenderEn": "Female",
			"Nation": null,
			"NationEn": "India",
			"Province": "Chai Nat",
			"ProvinceId": 6,
			"District": null,
			"ProvinceEn": "Chai Nat",
			"StatQuarantine": 1
		  }, {
			"ConfirmDate": "2021-05-01",
			"No": null,
			"Age": 35,
			"Gender": null,
			"GenderEn": null,
			"Nation": null,
			"NationEn": "USA",
			"Province": null,
			"ProvinceId": null,
			"District": null,
			"ProvinceEn": null,
			"StatQuarantine": 0
		  }, {
			"ConfirmDate": "2021-05-02",
			"No": null,
			"Age": 26,
			"Gender": null,
			"GenderEn": null,
			"Nation": null,
			"NationEn": "Thailand",
			"Province": "Kamphaeng Phet",
			"ProvinceId": 14,
			"District": null,
			"ProvinceEn": "Kamphaeng Phet",
			"StatQuarantine": 17
		  }
	]`)

	var jsonMap []model.Case
	json.Unmarshal([]byte(mockJsonData), &jsonMap)
	summarize, err := covid.GetSummarize(jsonMap)
	if err != nil {
		t.Error("Error should be nil")
	}
	if summarize.Province["Chai Nat"] != 1 {
		t.Error("Province Chai Nat should have 1 case")
	}
	if summarize.Province["Kamphaeng Phet"] != 1 {
		t.Error("Province Kamphaeng Phet should have 1 case")
	}
	if summarize.Province["N/A"] != 1 {
		t.Error("Province N/A should have 1 case")
	}
	if summarize.AgeGroup["0-30"] != 1 {
		t.Error("AgeGroup 0-30 should have 2 cases")
	}
	if summarize.AgeGroup["31-60"] != 1 {
		t.Error("AgeGroup 31-60 should have 0 cases")
	}
	if summarize.AgeGroup["61+"] != 0 {
		t.Error("AgeGroup 61+ should have 0 cases")
	}
	if summarize.AgeGroup["N/A"] != 1 {
		t.Error("AgeGroup N/A should have 1 case")
	}

}
