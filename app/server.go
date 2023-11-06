package main

import (
	"github.com/codecrafters-io/http-server-starter-go/app/server"
)

func main() {
	// You can use print statements as follows for debugging, they'll be visible when running tests.
	if serve, err := server.Connect("127.0.0.1:4221"); err == nil {
		serve.ServeConn()
	}
}
