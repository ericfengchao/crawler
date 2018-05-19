package parser

import (
	"io/ioutil"
	"testing"
)

func ParseCityList_test(t *testing.T) {
	contents, err := ioutil.ReadFile("cityList_test_data.html")

	if err != nil {
		t.Error(err)
	}
	result := ParseCityList(contents)
	expectedResult := 470

	if len(result.Requests) != expectedResult {
		t.Errorf("Expected %d, Got %d\n", expectedResult, len(result.Requests))
	}

	if len(result.Items) != expectedResult {
		t.Errorf("Expected %d, Got %d\n", expectedResult, len(result.Items))
	}
}
