//
// Copyright (c) 2020 Mediana SMS
// All rights reserved.
//
// Author: Asghar Dadashzadeh <a.dadashzadeh@mediana.ir>

// Package medianasms is an official library for working with mediana sms api.
// brief documentation for mediana sms api provided at http://docs.medianasms.com
package medianasms

import (
	"net/http"
	"net/url"
	"time"
)

const (
	// ClientVersion is used in User-Agent request header to provide server with API level.
	ClientVersion = "1.0.1"

	// Endpoint points you to MedianaSMS REST API.
	Endpoint = "http://rest.medianasms.com/v1"

	// httpClientTimeout is used to limit http.Client waiting time.
	httpClientTimeout = 30 * time.Second
)

// MedianaSMS ...
type MedianaSMS struct {
	AccessKey string
	Client    *http.Client
	BaseURL   *url.URL
}

// New create new mediana sms instance
func New(accesskey string) *MedianaSMS {
	u, _ := url.Parse(Endpoint)
	client := &http.Client{
		Transport: http.DefaultTransport,
		Timeout:   httpClientTimeout,
	}

	return &MedianaSMS{
		AccessKey: accesskey,
		Client:    client,
		BaseURL:   u,
	}
}
