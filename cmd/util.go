/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"runtime"
)

func getOss() oss {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("MAC operating system")
		return &macos{
			name: "macos",
		}
	case "linux":
		fmt.Println("Linux")
		return &linux{
			name: "linux",
		}
	default:
		log.Println("OS not supported")
		return nil
	}
}
