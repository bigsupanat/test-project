package obj

type CovidDataRequest struct {
	Data []struct {
		ConfirmDate    string      `json:"ConfirmDate"`
		No             interface{} `json:"No"`
		Age            int         `json:"Age"`
		Gender         string      `json:"Gender"`
		GenderEn       string      `json:"GenderEn"`
		Nation         interface{} `json:"Nation"`
		NationEn       string      `json:"NationEn"`
		Province       string      `json:"Province"`
		ProvinceID     int         `json:"ProvinceId"`
		District       interface{} `json:"District"`
		ProvinceEn     string      `json:"ProvinceEn"`
		StatQuarantine int         `json:"StatQuarantine"`
	} `json:"Data"`
}

type CovidDataSummaryResponse struct {
	Province map[string]int `json:"Province"`
	AgeGroup map[string]int `json:"AgeGroup"`
}
