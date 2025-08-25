package main

import "os"

func GetEnvWithFallback(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

var Version = GetEnvWithFallback("GOSU_VERSION", "PRE_RELEASE")
