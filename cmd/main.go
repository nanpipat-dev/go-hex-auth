package main

import (
	"fmt"
	"go-hex-auth/package/logger"
	"go-hex-auth/protocol"
)

func main() {
	logger.Info("Starting ....")
	err := protocol.Start()
	if err != nil {
		fmt.Println(err)
	}
}
