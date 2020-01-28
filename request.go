package medianasms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"runtime"
)

var (
	// ErrUnexpectedResponse is used when there was an internal server error and nothing can be done at this point.
	ErrUnexpectedResponse = errors.New("The MedianaSMS API is currently unavailable")
)

// ListParams ...
type ListParams struct {
	Limit int64 `json:"limit"`
	Page  int64 `json:"page"`
}

// PaginationInfo ...
type PaginationInfo struct {
	Total int64   `json:"total"`
	Limit int64   `json:"limit"`
	Page  int64   `json:"page"`
	Pages int64   `json:"pages"`
	Prev  *string `json:"prev"`
	Next  *string `json:"next"`
}

// BaseResponse base response model
type BaseResponse struct {
	Status string          `json:"status"`
	Code   ResponseCode    `json:"code"`
	Data   json.RawMessage `json:"data"`
	Meta   *PaginationInfo `json:"meta"`
}

// request preform http request
func (sms MedianaSMS) request(method string, uri string, params map[string]string, data interface{}) (*BaseResponse, error) {
	u := *sms.BaseURL

	// join base url with extra path
	u.Path = path.Join(sms.BaseURL.Path, uri)

	// set query params
	p := url.Values{}
	for key, param := range params {
		p.Add(key, param)
	}
	u.RawQuery = p.Encode()

	marshaledBody, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	requestBody := bytes.NewBuffer(marshaledBody)
	req, err := http.NewRequest(method, u.String(), requestBody)
	if err != nil {
		return nil, err
	}

	//req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "AccessKey "+sms.AccessKey)
	req.Header.Set("User-Agent", "MedianaSMS/ApiClient/"+ClientVersion+" Go/"+runtime.Version())

	res, err := sms.Client.Do(req)
	if err != nil || res == nil {
		return nil, err
	}

	defer res.Body.Close()

	responseBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusCreated:
		_res := &BaseResponse{}
		if err := json.Unmarshal(responseBody, _res); err != nil {
			return nil, fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
		}

		return _res, nil
	case http.StatusNoContent:
		// Status code 204 is returned for successful DELETE requests. Don't try to
		// unmarshal the body: that would return errors.
		return nil, nil
	case http.StatusInternalServerError:
		// Status code 500 is a server error and means nothing can be done at this
		// point.
		return nil, ErrUnexpectedResponse
	default:
		_res := &BaseResponse{}
		if err := json.Unmarshal(responseBody, _res); err != nil {
			return nil, fmt.Errorf("could not decode response JSON, %s: %v", string(responseBody), err)
		}

		return _res, ParseErrors(_res)
	}
}

// get do get request
func (sms MedianaSMS) get(uri string, params map[string]string) (*BaseResponse, error) {
	return sms.request("GET", uri, params, nil)
}

// post do post request
func (sms MedianaSMS) post(uri string, contentType string, data interface{}) (*BaseResponse, error) {
	return sms.request("POST", uri, nil, data)
}
