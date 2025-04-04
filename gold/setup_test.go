package main

import (
	"net/http"
	"os"
	"testing"
)

var testApp Config

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

var jsonToReturn = `
{
  "ts": 1743759521232,
  "tsj": 1743759516121,
  "date": "Apr 4th 2025, 05:38:36 am NY",
  "items": [
    {
      "curr": "USD",
      "xauPrice": 3091.325,
      "xagPrice": 31.3295,
      "chgXau": -11.695,
      "chgXag": -0.5657,
      "pcXau": -0.3769,
      "pcXag": -1.7736,
      "xauClose": 3103.02,
      "xagClose": 31.89517
    }
  ]
}
  `

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: fn,
	}
}
