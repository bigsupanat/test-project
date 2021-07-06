package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bigsupanat/test-project/obj"
)

type RequestService interface {
	Request(url string) (resp *http.Response, err error)
}

const ENDPOINT = "http://static.wongnai.com/devinterview/covid-cases.json"

func CovidSummary(svc RequestService) obj.CovidDataSummaryResponse {
	req := obj.CovidDataRequest{}
	err := getJsonData(svc, ENDPOINT, &req)
	if err != nil {
		log.Println("Error Get Json data", err)
	}
	sumProvince := make(map[string]int)
	sumAge := map[string]int{"0-30": 0, "31-60": 0, "61+": 0, "N/A": 0}
	for _, data := range req.Data {
		if _, ok := sumProvince[data.ProvinceEn]; ok {
			sumProvince[data.ProvinceEn]++
		} else {
			sumProvince[data.ProvinceEn] = 1
		}
		if data.Age <= 30 {
			sumAge["0-30"]++
		} else if data.Age <= 60 {
			sumAge["31-60"]++
		} else if data.Age == 0 {
			sumAge["N/A"]++
		} else {
			sumAge["61+"]++
		}
	}
	sumProvince["N/A"] = sumProvince[""]
	delete(sumProvince, "")
	res := obj.CovidDataSummaryResponse{
		Province: sumProvince,
		AgeGroup: sumAge,
	}
	return res
}

func getJsonData(svc RequestService, url string, req *obj.CovidDataRequest) error {
	r, err := svc.Request(url)
	if err != nil {
		log.Println("Error GetUrl")
		return err
	}
	defer r.Body.Close()
	byt, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Error Read Body")
		return err
	}
	err = ioutil.WriteFile("../covid-cases.json", byt, 0644)
	if err != nil {
		log.Println("Error Write File")
		return err
	}
	return json.Unmarshal(byt, req)
}
