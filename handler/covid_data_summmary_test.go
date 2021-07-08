package handler

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bigsupanat/test-project/obj"
	"github.com/go-playground/assert/v2"
)

type mockService1 struct{}

func (s mockService1) Request(url string) (resp *http.Response, err error) {
	mockResponse := `{
		"Data":[
		   {
			  "ConfirmDate":"2021-05-04",
			  "No":null,
			  "Age":51,
			  "Gender":"หญิง",
			  "GenderEn":"Female",
			  "Nation":null,
			  "NationEn":"China",
			  "Province":"Phrae",
			  "ProvinceId":46,
			  "District":null,
			  "ProvinceEn":"Phrae",
			  "StatQuarantine":5
		   },
		   {
			"ConfirmDate":"2021-05-01",
			"No":null,
			"Age":51,
			"Gender":"ชาย",
			"GenderEn":"Male",
			"Nation":null,
			"NationEn":"India",
			"Province":"Suphan Buri",
			"ProvinceId":65,
			"District":null,
			"ProvinceEn":"Suphan Buri",
			"StatQuarantine":8
		 	}
		]
	}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockResponse)
	}))
	defer ts.Close()
	return http.Get(ts.URL)
}

func TestGetJsonData(t *testing.T) {
	mocksvc := mockService1{}
	url := ""
	actual := obj.CovidDataRequest{}

	expected := obj.CovidDataRequest{
		Datas: []obj.Data{
			{ConfirmDate: "2021-05-04", No: "", Age: 51, Gender: "หญิง", GenderEn: "Female", Nation: "", NationEn: "China", Province: "Phrae", ProvinceID: 46, District: "", ProvinceEn: "Phrae", StatQuarantine: 5},
			{ConfirmDate: "2021-05-01", No: "", Age: 51, Gender: "ชาย", GenderEn: "Male", Nation: "", NationEn: "India", Province: "Suphan Buri", ProvinceID: 65, District: "", ProvinceEn: "Suphan Buri", StatQuarantine: 8},
		},
	}
	err := getJsonData(mocksvc, url, &actual)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, expected, actual)
}

type mockService2 struct{}

func (s mockService2) Request(url string) (resp *http.Response, err error) {
	mockResponse := `{
		"Data":[
			{
				"ConfirmDate":null,
				"No":null,
				"Age":7,
				"Gender":null,
				"GenderEn":null,
				"Nation":null,
				"NationEn":"China",
				"Province":"Ratchaburi",
				"ProvinceId":51,
				"District":null,
				"ProvinceEn":"Ratchaburi",
				"StatQuarantine":20
			},
			{
				"ConfirmDate":"2021-05-03",
				"No":null,
				"Age":25,
				"Gender":"หญิง",
				"GenderEn":"Female",
				"Nation":null,
				"NationEn":"India",
				"Province":"Ratchaburi",
				"ProvinceId":51,
				"District":null,
				"ProvinceEn":"Ratchaburi",
				"StatQuarantine":0
			},
			{
				"ConfirmDate":"2021-05-02",
				"No":null,
				"Age":40,
				"Gender":"หญิง",
				"GenderEn":"Female",
				"Nation":null,
				"NationEn":"China",
				"Province":"Ratchaburi",
				"ProvinceId":51,
				"District":null,
				"ProvinceEn":"Ratchaburi",
				"StatQuarantine":10
			},
			{
				"ConfirmDate":"2021-05-03",
				"No":null,
				"Age":98,
				"Gender":"หญิง",
				"GenderEn":"Female",
				"Nation":null,
				"NationEn":null,
				"Province":"Nonthaburi",
				"ProvinceId":35,
				"District":null,
				"ProvinceEn":"Nonthaburi",
				"StatQuarantine":7
			},
			{
				"ConfirmDate":"2021-05-02",
				"No":null,
				"Age":87,
				"Gender":"ชาย",
				"GenderEn":"Male",
				"Nation":null,
				"NationEn":"India",
				"Province":"Nonthaburi",
				"ProvinceId":35,
				"District":null,
				"ProvinceEn":"Nonthaburi",
				"StatQuarantine":2
			}
		]
	}`
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, mockResponse)
	}))
	defer ts.Close()
	return http.Get(ts.URL)
}

func TestCovidSummary(t *testing.T) {
	svc := mockService2{}
	expected := obj.CovidDataSummaryResponse{Province: map[string]int{"N/A": 0, "Nonthaburi": 2, "Ratchaburi": 3}, AgeGroup: map[string]int{"0-30": 2, "31-60": 1, "61+": 2, "N/A": 0}}
	actual := CovidSummary(svc)
	assert.Equal(t, expected, actual)
}
