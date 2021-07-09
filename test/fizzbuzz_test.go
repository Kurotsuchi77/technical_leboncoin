// +build e2e

package test

import (
	"testing"
	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

// TestGetFizzBuzzOk - Tests result of a fizzbuzz request in the nominal use case
func TestGetFizzBuzzOk(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/fizzbuzz/request?int1=3&int2=5&limit=100&str1=fizz&str2=buzz")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}

// TestGetFizzBuzzWrongParams - Tests result of a fizzbuzz request when one of the parameter is of the wrong type
func TestGetFizzBuzzWrongParams(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/fizzbuzz/request?int1=string&int2=2&limit=100&str1=fizz&str2=buzz")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 404, resp.StatusCode())
}

// TestGetFizzBuzzMissingParams - Tests result of a fizzbuzz request when one of the parameter is missing
func TestGetFizzBuzzMissingParams(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/fizzbuzz/request?int1=1")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 404, resp.StatusCode())
}

// TestGetMostRequested - Tests result of a statistics request
func TestGetMostRequested(t *testing.T) {
	client := resty.New()
	resp, err := client.R().Get(BASE_URL + "/api/fizzbuzz/request?int1=3&int2=5&limit=100&str1=fizz&str2=buzz")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())

	resp, err = client.R().Get(BASE_URL + "/api/fizzbuzz/statistics")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}