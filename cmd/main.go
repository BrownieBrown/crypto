package main

import "crypto/pkg/logger"

func main() {
	defer logger.Log.Sync()
}
