package covid

import (
	"github.com/mercuryheart123/lmwn-backend-go/model"
)

type AllSummarize struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}

func GetSummarize(data []model.Case) (*AllSummarize, error) {
	var allSummarize AllSummarize
	provinces := make(map[string]int)
	ageGroup := make(map[string]int)

	for _, eachCase := range data {
		provinceName := eachCase.Province
		if provinceName == "" {
			provinceName = "N/A"
		}
		provinces[provinceName]++

		if age := eachCase.Age; age != nil {
			if *age < 30 {
				ageGroup["0-30"]++
			} else if *age < 60 {
				ageGroup["31-60"]++
			} else {
				ageGroup["61+"]++
			}
			continue
		}
		ageGroup["N/A"]++
	}
	allSummarize.Province = provinces
	allSummarize.AgeGroup = ageGroup
	return &allSummarize, nil
}
