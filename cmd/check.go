/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: command,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func command(cmd *cobra.Command, args []string) {
	oss := getEnv()
	if oss == nil {
		log.Fatal("OS not supported")
	}

	// Check if there are any arguments
	if len(args) == 0 {
		fmt.Println("No arguments provided")
		return
	}
	gitURL := fmt.Sprintf("https://github.com/%s", args[0])

	// Create a temp directory
	dir, err := os.MkdirTemp(".tmp", "sidewalk-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	// Clone the repository
	_, err = exec.Command("git", "clone", gitURL, dir).Output()
	if err != nil {
		log.Fatal(err)
	}

	// Get the list of items in the directory
	listDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("List items: ", listDir)
	for _, item := range listDir {
		reg, err := regexp.Compile(`[^\w]`)
		if err != nil {
			log.Fatal(err)
		}
		name := reg.ReplaceAllString(item.Name(), "")
		if item.IsDir() {
			_, err := oss.check(name)
			if err != nil {
				fmt.Println(name, "Not installed ")
			} else {
				fmt.Println(name, "Already installed ")
			}
		}
	}
}
