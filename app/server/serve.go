package server

import (
	"fmt"
	"net"
	"os"
	"path/filepath"
	"strings"
)

type Serve struct {
	listener net.Listener
}

func Connect(ip string) (Serve, error) {
	if listener, err := net.Listen("tcp", ip); err == nil {
		return Serve{listener}, nil
	} else {
		return Serve{}, err
	}
}

func (serve Serve) ServeConn() error {
	for {
		if conn, err := serve.listener.Accept(); err == nil {
			go handleConn(conn)
		} else {
			return err
		}
	}
}

func handleConn(conn net.Conn) {
	// Buffer for the request
	buf := make([]byte, 1024)
	// Read the request to the buffer, check there was no error
	if readed, err := conn.Read(buf); err == nil {
		request_data, err := ParseRequestData(buf[0:readed])
		if err != nil {
			conn.Close()
			return
		}
		var response string
		headers := make(map[string]string)

		if request_data.path == "/" {
			response = CreateResponse(nil, "", OK)
		} else if request_data.path == "/user-agent" {
			content := request_data.headers["User-Agent"]

			headers["Content-Type"] = "text/plain"
			headers["Content-Length"] = fmt.Sprint(len(content))

			response = CreateResponse(headers, content, OK)

		} else if strings.HasPrefix(request_data.path, "/echo") {
			content := strings.TrimPrefix(request_data.path, "/echo/")

			headers["Content-Type"] = "text/plain"
			headers["Content-Length"] = fmt.Sprint(len(content))

			response = CreateResponse(headers, content, OK)

		} else if strings.HasPrefix(request_data.path, "/files/") {
			// The full file path = argument of the program + the path of the file on the path request
			file_path := filepath.Join(os.Args[2], strings.TrimPrefix(request_data.path, "/files/"))
			switch string(request_data.method) {
			case GET:
				// We try to read the file and return 404 if we fail
				if file, err_file := os.ReadFile(file_path); err_file == nil {
					headers["Content-Type"] = "application/octet-stream"
					headers["Content-Length"] = fmt.Sprint(len(file))
					response = CreateResponse(headers, string(file), OK)
				} else {
					response = CreateResponse(headers, "", NOT_FOUND)
				}
			case POST:
				if err_file := os.WriteFile(file_path, request_data.body, 0644); err_file == nil {
					response = CreateResponse(headers, "", CREATED)
				}
			}
		} else {
			response = CreateResponse(headers, "", NOT_FOUND)
		}
		conn.Write([]byte(response))
		conn.Close()
		return
	}
}
