package server

import "bytes"

const (
	NOT_FOUND = "HTTP/1.1 404 Not Found\r\n"
	OK        = "HTTP/1.1 200 OK \r\n"
	CREATED   = "HTTP/1.1 201 Created\r\n"
)

func CreateHeaders(headers map[string]string) string {
	header_str := ""
	for key, value := range headers {
		header_str += key + ": " + value + "\r\n"
	}
	return header_str
}

func CreateResponse(headers map[string]string, body string, status string) string {
	response := status + CreateHeaders(headers) + "\r\n" + body
	return response
}

func ParseRequest(header *[]byte) (method []byte, path []byte, version []byte) {
	requests_part := bytes.Split(*header, []byte(" "))
	method = requests_part[0]
	path = requests_part[1]
	version = requests_part[2]

	return method, path, version
}
