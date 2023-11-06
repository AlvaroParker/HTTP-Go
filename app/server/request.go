package server

import (
	"bytes"
	"errors"
	"fmt"
)

// HTTP request methods
const (
	CONNECT = "CONNECT"
	DELETE  = "DELETE"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	TRACE   = "TRACE"
)

type RequestData struct {
	method       string
	path         string
	http_version string
	headers      map[string]string
	body         []byte
}

func ParseRequestData(buf []byte) (RequestData, error) {
	parts := bytes.Split(buf, []byte("\r\n\r\n"))
	if len(parts) < 2 {
		return RequestData{}, errors.New("Expected \\r\\n\\r\\n")
	}
	// Request and headers part
	requests_n_headers := bytes.Split(parts[0], []byte("\r\n"))

	// The request (1st line of an HTTP request)
	request_full := bytes.Split(requests_n_headers[0], []byte(" "))
	if len(request_full) < 3 {
		return RequestData{}, errors.New(
			fmt.Sprintf("The first line of the requst is incomplete, got:\n\"%s\"",
				string(requests_n_headers[0])))
	}
	method := request_full[0]
	path := request_full[1]
	http_version := request_full[2]

	// The headers (between the HTTP request info and the body)
	headers := requests_n_headers[1:]
	headers_map := ParseRequestHeaders(headers)

	// Body part
	body_arr := parts[1]

	return RequestData{
		method:       string(method),
		path:         string(path),
		http_version: string(http_version),
		headers:      headers_map,
		body:         body_arr,
	}, nil
}

func ParseRequestHeaders(raw_headers [][]byte) map[string]string {
	headers_map := make(map[string]string)
	for _, raw_header := range raw_headers {
		header := bytes.SplitN(raw_header, []byte(": "), 2)
		key := header[0]
		value := header[1]
		headers_map[string(key)] = string(value)
	}
	return headers_map
}
