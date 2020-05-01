package yourls

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
)

// Service ...
type Service struct {
	cO Options
}

func (c *Service) setOptions(yourlsOptions Options) BuildYourls {
	c.cO = yourlsOptions
	return c
}

// SendAction ...
func (c *Service) SendAction(request interface{}) (returnData interface{}, errorResponse interface{}, resp *http.Response, err error) {
	errorResponse = &GeneralErrorResponse{}
	switch r := request.(type) {
	case ActionShortURLRequest:
		returnData = &ActionShortURLResponse{}
		r.Action = ShortURL
		r.Format = Json
		r.Signature = c.cO.Token
		r.UserName = c.cO.Username
		r.Password = c.cO.Password
		resp, err = c.sendActionUnknown(r, returnData, errorResponse)
	case ActionExpandRequest:
		returnData = &ActionExpandResponse{}
		r.Action = Expand
		r.Format = Json
		r.Signature = c.cO.Token
		r.UserName = c.cO.Username
		r.Password = c.cO.Password
		resp, err = c.sendActionUnknown(r, returnData, errorResponse)
	case ActionURLStatsRequest:
		returnData = &ActionURLStatsResponse{}
		r.Action = URLStats
		r.Format = Json
		r.Signature = c.cO.Token
		r.UserName = c.cO.Username
		r.Password = c.cO.Password
		resp, err = c.sendActionUnknown(r, returnData, errorResponse)
	case ActionDBStatsRequest:
		returnData = &ActionDBStatsSimpleResponse{}
		r.Action = DBStats
		r.Format = Json
		r.Signature = c.cO.Token
		r.UserName = c.cO.Username
		r.Password = c.cO.Password
		resp, err = c.sendActionUnknown(r, returnData, errorResponse)
	case ActionStatsRequest:
		if r.Limit == 0 {
			returnData = &ActionStatsSimpleResponse{}
		} else {
			returnData = &ActionStatsFullResponse{}
		}
		r.Action = Stats
		r.Format = Json
		r.Signature = c.cO.Token
		r.UserName = c.cO.Username
		r.Password = c.cO.Password
		resp, err = c.sendActionUnknown(r, returnData, errorResponse)
	default:
		returnData = &ActionUnknownResponse{}
		resp, err = c.sendActionUnknown(request, returnData, errorResponse)
	}

	return
}

func (c *Service) sendActionUnknown(request interface{}, response interface{}, errorResponse interface{}) (resp *http.Response, err error) {
	req, err := c.newRequest(http.MethodPost, YourlsAPIEndpoint, nil)

	if err != nil {
		return
	}

	q := req.URL.Query()
	for k, v := range c.structToMap(request) {
		q.Add(k, v)
	}

	req.URL.RawQuery = q.Encode()

	return c.do(req, response, errorResponse)
}

func (c *Service) newRequest(method, path string, form url.Values) (req *http.Request, err error) {
	rel := &url.URL{Path: path}
	u := c.cO.baseURL.ResolveReference(rel)
	buf := bytes.NewBufferString(form.Encode())
	return http.NewRequest(method, u.String(), buf)
}

func (c *Service) structToMap(request interface{}) map[string]string {
	val := reflect.ValueOf(request)
	m := make(map[string]string, val.NumField())

	for i := 0; i < val.NumField(); i++ {
		tag := val.Type().Field(i).Tag.Get(FormKeyTagName)
		switch val.Field(i).Kind() {
		case reflect.String:
			{
				d := fmt.Sprintf("%v", val.Field(i).Interface())
				if d != "" {
					m[tag] = d
				}
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			{
				d := fmt.Sprintf("%v", val.Field(i).Interface())
				if d != "0" {
					m[tag] = d
				}
			}
		}
	}

	return m
}

func (c *Service) do(req *http.Request, normalResponse interface{}, errorResponse interface{}) (resp *http.Response, err error) {
	resp, err = c.cO.httpClient.Do(req)

	if err != nil {
		return
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	bodyByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return
	}

	_ = json.Unmarshal(bodyByte, &normalResponse)
	_ = json.Unmarshal(bodyByte, &errorResponse)

	return
}
