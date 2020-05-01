# Go SDK YOURLS
Golang SDK for YOURLS

[![Go Report Card](https://goreportcard.com/badge/github.com/kitabisa/go_sdk_yourls?style=flat-square)](https://goreportcard.com/report/github.com/kitabisa/go_sdk_yourls)
[![Build Status](http://img.shields.io/travis/kitabisa/go_sdk_yourls.svg?style=flat-square)](https://travis-ci.org/kitabisa/go_sdk_yourls)
[![Codecov](https://img.shields.io/codecov/c/github/kitabisa/go_sdk_yourls.svg?style=flat-square)](https://codecov.io/gh/kitabisa/go_sdk_yourls)
[![Maintainability](https://api.codeclimate.com/v1/badges/9d35afa5f60a03fdee63/maintainability)](https://codeclimate.com/github/kitabisa/go_sdk_yourls/maintainability)
[![License](http://img.shields.io/badge/license-mit-blue.svg?style=flat-square)](https://raw.githubusercontent.com/kitabisa/go_sdk_yourls/master/LICENSE)
[![Release](https://img.shields.io/github/v/release/kitabisa/go_sdk_yourls.svg?style=flat&color=green)](https://github.com/kitabisa/go_sdk_yourls/releases)


## Features
### Your Own URL Shortener API 
Based on `https://yourls.org/#API`

#### Example
```go
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"

	yourls "github.com/kitabisa/go_sdk_yourls"
)

func setYourlsBuild(httpClient *http.Client, config *Config) yourls.BuildYourls {
	yourlsBuilder := &yourls.Builder{}
	service := &yourls.Service{}
	baseURL, _ := url.Parse(config.fYoursURL)

	yourlsBuilder.SetBuilder(service)
	yourlsBuilder.SetHTTPClient(httpClient)
	yourlsBuilder.SetBaseURL(baseURL)
	yourlsBuilder.SetToken(config.fSignature)
	yourlsBuilder.SetUsername(config.fUsername)
	yourlsBuilder.SetPassword(config.fPassword)

	return yourlsBuilder.Build()
}

func main() {
	config := parseConfigFromArgs(os.Args)

	if !isValidUrl(config.fYoursURL) {
		_, _ = fmt.Fprintf(os.Stderr, "Use -yours_url to specify YoursURL host : %s\n", config.fYoursURL)
		os.Exit(1)
	}

	if config.fSignature == "" && config.fUsername == "" {
		_, _ = fmt.Fprintf(os.Stderr, "Use -signature of (-username with -password) of Yourls credentials\n")
		os.Exit(1)
	}

	httpClient := &http.Client{
		Timeout: 5 * time.Second,
	}

	yourlsBuild := setYourlsBuild(httpClient, &config)

	var requestStruct interface{}

	switch config.fAction {
	case string(yourls.ShortURL):
		requestStruct = yourls.ActionShortURLRequest{
			URL:     config.fURL,
			Keyword: config.fKeyword,
			Title:   config.fTitle,
		}

	case string(yourls.Expand):
		requestStruct = yourls.ActionExpandRequest{
			ShortURL: config.fShortURL,
		}

	case string(yourls.URLStats):
		requestStruct = yourls.ActionURLStatsRequest{
			ShortURL: config.fShortURL,
		}

	case string(yourls.DBStats):
		requestStruct = yourls.ActionDBStatsRequest{}

	case string(yourls.Stats):
		requestStruct = yourls.ActionStatsRequest{
			Filter: yourls.Filter(config.fFilter),
			Limit:  config.fLimit,
		}

	default:
		_, _ = fmt.Fprintf(os.Stderr, "Use -action to specify action (shorturl, expand, url-stats, stats, db-stats) : %s\n", config.fAction)
		os.Exit(1)
	}

	returnData, errorResponse, resp, err := yourlsBuild.SendAction(requestStruct)

	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error : %s\n", err)
		os.Exit(1)
	}

	returnDataString, _ := json.Marshal(returnData)
	errorResponseString, _ := json.Marshal(errorResponse)

	fmt.Printf("HTTP response Status \t: %v\n", resp.StatusCode)
	fmt.Printf("returnData \t\t: %s\n", returnDataString)
	fmt.Printf("errorResponse \t\t: %s\n", errorResponseString)
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}

type Config struct {
	fYoursURL  string
	fSignature string
	fUsername  string
	fPassword  string
	fAction    string
	fURL       string
	fKeyword   string
	fTitle     string
	fShortURL  string
	fFilter    string
	fLimit     int
}

func parseConfigFromArgs(args []string) Config {
	config := Config{}

	flagSet := flag.NewFlagSet(args[0], flag.ExitOnError)

	flagSet.StringVar(&config.fYoursURL, "yours_url", "", "Yours URL address")
	flagSet.StringVar(&config.fSignature, "signature", "", "Signature for authorization")
	flagSet.StringVar(&config.fUsername, "username", "", "Username for authorization")
	flagSet.StringVar(&config.fPassword, "password", "", "Password for authorization")
	flagSet.StringVar(&config.fAction, "action", "", "Type of action")
	flagSet.StringVar(&config.fURL, "url", "", "URL to shorten")
	flagSet.StringVar(&config.fKeyword, "keyword", "", "Keyword to custom short URLs")
	flagSet.StringVar(&config.fTitle, "title", "", "Title to custom short URLs")
	flagSet.StringVar(&config.fShortURL, "shorturl", "", "ShortURL is shorted URL address")
	flagSet.StringVar(&config.fFilter, "filter", "", "Type of filter of stats action (top, bottom, rand, last)")
	flagSet.IntVar(&config.fLimit, "limit", 0, "Limit of displayed links of stats action")

	_ = flagSet.Parse(args[1:])

	return config
}

```

## Test, Code Coverage & Benchmark
```bash
go test -v -cover ./... -bench=.
```

## Installation
```bash
go get -u github.com/kitabisa/go_sdk_yourls
```


## License
MIT License