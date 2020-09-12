package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty"
	"github.com/stretchr/testify/assert"
	"testing"
)

type HtmlResponse struct {
	HtmlTitle    string `json:"htmltitle"`
	HtmlVersion  string `json:"htmlversion"`
	HeadingCount int    `json:"headingcount"`
	ExternalLink int    `json:"externallink"`
	InternalLink int    `json:"internalink"`
	Inaccessible int    `json:"inaccessible"`
	IsLogin      bool   `json:"islogin"`
}

func Test_StatusCodeShouldEqual200(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("https://www.facebook.com/")

	if resp.StatusCode() != 200 {
		t.Errorf("Unexpected status code, expected %d, got %d instead", 200, resp.StatusCode())
	}
}

func Test_ContentTypeShouldEqualApplicationJson(t *testing.T) {

	client := resty.New()

	resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

	assert.Equal(t, "application/json", resp.Header().Get("Content-Type"))
}

func Test_GetResponseShouldEqualToMockResponse(t *testing.T) {
	//here is the mock json data for testing purpose, once api is build so we can hit the api and get the json response
	resp := HtmlResponse{
		HtmlTitle:    "Go - The Programming Language",
		HtmlVersion:  "html5",
		HeadingCount: 2,
		ExternalLink: 5,
		InternalLink: 10,
		Inaccessible: 15,
		IsLogin:      false,
	}
	/*
		Here we call our api and get the json response
		resp, _ := client.R().Get("http://api.zippopotam.us/us/90210")

		myResponse := HtmlResponse{}

		err := json.Unmarshal(resp.Body(), &myResponse)

		if err != nil {
			fmt.Println(err)
			return
		}*/

	assert.Equal(t, "Go - The Programming Language", resp.HtmlTitle)
	assert.Equal(t, "html5", resp.HtmlVersion)
	assert.Equal(t, 5, resp.ExternalLink)
}
