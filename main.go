package main

import "github.com/danyukod/go-client-server-api/src/server"

func main() {

	err := server.Serve()
	if err != nil {
		return
	}

}
