package api

import "github.com/go-resty/resty/v2"

var client = resty.New().
	SetBaseURL("https://dummyjson.com").
	SetHeader("Accept", "application/json")
