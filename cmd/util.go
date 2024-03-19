/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"runtime"
)

func getEnv() env {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MAC operating system")
		return &macos{}
	case "linux":
		fmt.Println("Linux")
		return &linux{}
	default:
		log.Println("OS not supported")
		return nil
	}
}
