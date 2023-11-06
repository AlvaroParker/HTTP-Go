# HTTP-Go
Basic HTTP server builded in Go. Run `go ./app` to start the server. 

## Features
- Get the current `user-agent`
- Get a file stored on the device when the server is initialized with `--directory <directory>`
- Upload a file to that directory
- HTTP responses to all paths (`404` to the ones that are not valid)
