package yourls

import (
	"context"
	"crypto/tls"
	"fmt"
	"math/rand"
	"net"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"gotest.tools/assert"
)

const (
	yoursURL   = "http://kb.in"
	username   = "default"
	password   = "default"
	signature  = "f6d89851aa"
	longURL    = "https://url.com/?a7"
	keyword    = "a7"
	shortURL   = yoursURL + "/" + keyword
	title      = "Link URL"
	dateString = "2020-04-30 21:08:44"
	ipString   = "172.17.0.1"
	success    = "success"
	s200       = "200"
)

type testData struct {
	requestStruct         interface{}
	responseString        string
	expectedResponse      interface{}
	expectedErrorResponse interface{}
	expectedError         error
}

var tests = make(map[string]testData)

func init() {
	tests["shorturl"] = testData{
		requestStruct: ActionShortURLRequest{
			URL:     longURL,
			Keyword: keyword,
			Title:   title,
		},
		responseString: fmt.Sprintf(
			`{
				"url": {
					"keyword": "%s",
					"url": "%s",
					"title": "%s",
					"date": "%s",
					"ip": "%s"
				},
				"status": "%s",
				"message": "%s added to database",
				"title": "%s",
				"shorturl": "%s",
				"statusCode": %s
			}`,
			keyword,
			longURL,
			title,
			dateString,
			ipString,
			success,
			longURL,
			title,
			shortURL,
			s200,
		),
		expectedResponse: &ActionShortURLResponse{
			URL: ActionShortURL{
				Keyword: keyword,
				URL:     longURL,
				Title:   title,
				Date:    dateString,
				IP:      ipString,
			},
			Status:     success,
			Message:    longURL + " added to database",
			Title:      title,
			ShortURL:   shortURL,
			StatusCode: 200,
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: longURL + " added to database",
		},
		expectedError: nil,
	}

	tests["expand"] = testData{
		requestStruct: ActionExpandRequest{
			ShortURL: longURL,
		},
		responseString: fmt.Sprintf(
			`{
				"keyword": "%s",
				"shorturl": "%s",
				"longurl": "%s",
				"title": "%s",
				"message": "%s",
				"statusCode": %s
			}`,
			keyword,
			shortURL,
			longURL,
			title,
			success,
			s200,
		),
		expectedResponse: &ActionExpandResponse{
			Keyword:    keyword,
			ShortURL:   shortURL,
			LongURL:    longURL,
			Title:      title,
			Message:    success,
			StatusCode: 200,
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: success,
		},
		expectedError: nil,
	}

	tests["url-stats"] = testData{
		requestStruct: ActionURLStatsRequest{
			ShortURL: shortURL,
		},
		responseString: fmt.Sprintf(
			`{
				"statusCode": %s,
				"message": "%s",
				"link": {
					"shorturl": "%s",
					"url": "%s",
					"title": "%s",
					"timestamp": "%s",
					"ip": "%s",
					"clicks": "0"
				}
			}`,
			s200,
			success,
			shortURL,
			longURL,
			title,
			dateString,
			ipString,
		),
		expectedResponse: &ActionURLStatsResponse{
			Message:    success,
			StatusCode: 200,
			Link: ActionShortURL{
				ShortURL:  shortURL,
				URL:       longURL,
				Title:     title,
				TimeStamp: dateString,
				IP:        ipString,
				Clicks:    "0",
			},
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: success,
		},
		expectedError: nil,
	}

	tests["db-stats"] = testData{
		requestStruct: ActionDBStatsRequest{},
		responseString: fmt.Sprintf(
			`{
				"db-stats": {
					"total_links": "19",
					"total_clicks": "2"
				},
				"statusCode": %s,
				"message": "%s"
			}`,
			s200,
			success,
		),
		expectedResponse: &ActionDBStatsSimpleResponse{
			Message:    success,
			StatusCode: 200,
			Stats: ActionStats{
				TotalLinks:  "19",
				TotalClicks: "2",
			},
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: success,
		},
		expectedError: nil,
	}

	tests["stats-simple"] = testData{
		requestStruct: ActionStatsRequest{},
		responseString: fmt.Sprintf(
			`{
				"stats": {
					"total_links": "19",
					"total_clicks": "2"
				},
				"statusCode": %s,
				"message": "%s"
			}`,
			s200,
			success,
		),
		expectedResponse: &ActionStatsSimpleResponse{
			Message:    success,
			StatusCode: 200,
			Stats: ActionStats{
				TotalLinks:  "19",
				TotalClicks: "2",
			},
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: success,
		},
		expectedError: nil,
	}

	rand.Seed(time.Now().Unix())
	filter := [...]Filter{Top, Bottom, Rand, Last}
	tests["stats-full"] = testData{
		requestStruct: ActionStatsRequest{
			Filter: filter[rand.Intn(len(filter))],
			Limit:  2,
		},
		responseString: fmt.Sprintf(
			`{
				"links": {
					"link_1": {
						"shorturl": "%s",
						"url": "%s",
						"title": "%s",
						"timestamp": "%s",
						"ip": "%s",
						"clicks": "0"
					}
				},
				"stats": {
					"total_links": "19",
					"total_clicks": "2"
				},
				"statusCode": %s,
				"message": "%s"
			}`,
			shortURL,
			longURL,
			title,
			dateString,
			ipString,
			s200,
			success,
		),
		expectedResponse: &ActionStatsFullResponse{
			Message:    success,
			StatusCode: 200,
			Links: map[string]ActionShortURL{
				"link_1": {
					ShortURL:  shortURL,
					URL:       longURL,
					Title:     title,
					TimeStamp: dateString,
					IP:        ipString,
					Clicks:    "0",
				},
			},
			Stats: ActionStats{
				TotalLinks:  "19",
				TotalClicks: "2",
			},
		},
		expectedErrorResponse: &GeneralErrorResponse{
			Message: success,
		},
		expectedError: nil,
	}
}

func setYourlsBuild(httpClient *http.Client) BuildYourls {
	yourlsBuilder := &Builder{}
	service := &Service{}
	baseURL, _ := url.Parse(yoursURL)

	yourlsBuilder.SetBuilder(service)
	yourlsBuilder.SetHTTPClient(httpClient)
	yourlsBuilder.SetBaseURL(baseURL)
	yourlsBuilder.SetToken(signature)
	yourlsBuilder.SetUsername(username)
	yourlsBuilder.SetPassword(password)

	return yourlsBuilder.Build()
}

func testingHTTPClient(handler http.Handler, isHTTPS bool) (*http.Client, func()) {
	if isHTTPS {
		s := httptest.NewTLSServer(handler)
		cli := &http.Client{
			Transport: &http.Transport{
				DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
					return net.Dial(network, s.Listener.Addr().String())
				},
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
			Timeout: 1000 * time.Millisecond,
		}
		return cli, s.Close
	}

	s := httptest.NewServer(handler)
	cli := &http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, network, _ string) (net.Conn, error) {
				return net.Dial(network, s.Listener.Addr().String())
			},
		},
		Timeout: 1000 * time.Millisecond,
	}
	return cli, s.Close
}

func anyHandler(bodyResponse string, httpStatus int) (handler http.Handler) {
	rand.Seed(time.Now().Unix())
	body := bodyResponse
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		_, _ = w.Write([]byte(body))
	})
}

func TestSendAction(t *testing.T) {
	for name, test := range tests {
		t.Run(name, func(tt *testing.T) {
			handler := anyHandler(test.responseString, 200)
			httpClient, teardown := testingHTTPClient(handler, false)
			defer teardown()
			yourls := setYourlsBuild(httpClient)
			returnData, errorResponse, resp, err := yourls.SendAction(test.requestStruct)

			if resp != nil {
				assert.Equal(
					tt,
					resp.StatusCode,
					http.StatusOK,
				)
			}

			assert.DeepEqual(
				tt,
				test.expectedResponse,
				returnData,
			)

			assert.DeepEqual(
				tt,
				test.expectedErrorResponse,
				errorResponse,
			)

			if test.expectedError == nil {
				assert.NilError(tt, err)
			} else {
				assert.Error(
					tt,
					err,
					test.expectedError.Error(),
				)
			}
		})
	}
}

func BenchmarkSendAction(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for name, test := range tests {
			b.Run(name, func(bb *testing.B) {
				handler := anyHandler(test.responseString, 200)
				httpClient, teardown := testingHTTPClient(handler, false)
				defer teardown()
				yourls := setYourlsBuild(httpClient)
				returnData, errorResponse, resp, err := yourls.SendAction(test.requestStruct)

				if resp != nil {
					assert.Equal(
						bb,
						resp.StatusCode,
						http.StatusOK,
					)
				}

				assert.DeepEqual(
					bb,
					test.expectedResponse,
					returnData,
				)

				assert.DeepEqual(
					bb,
					test.expectedErrorResponse,
					errorResponse,
				)

				if test.expectedError == nil {
					assert.NilError(bb, err)
				} else {
					assert.Error(
						bb,
						err,
						test.expectedError.Error(),
					)
				}
			})
		}
	}
}
