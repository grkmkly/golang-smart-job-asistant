package main

import (
	"smartjob/internal/server"
)

func main() {

	server := server.NewServer()
	server.Run()

}
