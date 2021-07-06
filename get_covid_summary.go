package main

// const ENDPOINT = "http://static.wongnai.com/devinterview/covid-cases.json"

// type CovidDataRequest struct {
// 	Data []struct {
// 		ConfirmDate    string      `json:"ConfirmDate"`
// 		No             interface{} `json:"No"`
// 		Age            int         `json:"Age"`
// 		Gender         string      `json:"Gender"`
// 		GenderEn       string      `json:"GenderEn"`
// 		Nation         interface{} `json:"Nation"`
// 		NationEn       string      `json:"NationEn"`
// 		Province       string      `json:"Province"`
// 		ProvinceID     int         `json:"ProvinceId"`
// 		District       interface{} `json:"District"`
// 		ProvinceEn     string      `json:"ProvinceEn"`
// 		StatQuarantine int         `json:"StatQuarantine"`
// 	} `json:"Data"`
// }

// type CovidDataSummaryResponse struct {
// 	Province map[string]int `json:"Province"`
// 	AgeGroup map[string]int `json:"AgeGroup"`
// }

// func getCovidSummary(c *gin.Context) {
// 	req := CovidDataRequest{}
// 	cli := InitClient()
// 	err := getJsonData(cli, ENDPOINT, &req)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res := handlerCovidSummary(req)
// 	c.JSON(http.StatusOK, res)
// }

// func handlerCovidSummary(req CovidDataRequest) CovidDataSummaryResponse {
// 	sumProvince := make(map[string]int)
// 	sumAge := map[string]int{"0-30": 0, "31-60": 0, "61+": 0, "N/A": 0}
// 	for _, data := range req.Data {
// 		if _, ok := sumProvince[data.ProvinceEn]; ok {
// 			sumProvince[data.ProvinceEn]++
// 		} else {
// 			sumProvince[data.ProvinceEn] = 1
// 		}
// 		if data.Age <= 30 {
// 			sumAge["0-30"]++
// 		} else if data.Age <= 60 {
// 			sumAge["31-60"]++
// 		} else if data.Age == 0 {
// 			sumAge["N/A"]++
// 		} else {
// 			sumAge["61+"]++
// 		}
// 	}
// 	sumProvince["N/A"] = sumProvince[""]
// 	delete(sumProvince, "")
// 	res := CovidDataSummaryResponse{
// 		Province: sumProvince,
// 		AgeGroup: sumAge,
// 	}
// 	return res
// }

// type RequestService interface {
// 	Request(url string) (resp *http.Response, err error)
// }

// func getJsonData(svc RequestService, url string, req *CovidDataRequest) error {
// 	// c := &http.Client{Timeout: 10 * time.Second}
// 	// r, err := c.Get(url)
// 	r, err := svc.Request(url)
// 	if err != nil {
// 		log.Println("Error GetUrl")
// 		return err
// 	}
// 	defer r.Body.Close()
// 	byt, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		log.Println("Error Read Body")
// 		return err
// 	}
// 	err = ioutil.WriteFile("covid-cases.json", byt, 0644)
// 	if err != nil {
// 		log.Println("Error Write File")
// 		return err
// 	}
// 	return json.Unmarshal(byt, req)
// }

// func getJsonFromFile(url string, req *CovidDataRequest) error {
// 	byt, err := ioutil.ReadFile("covid-cases.json")
// 	if err != nil {
// 		log.Println("Error Read File")
// 		return err
// 	}
// 	return json.Unmarshal(byt, &req)
// }
