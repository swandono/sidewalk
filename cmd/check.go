/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

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

	// List the items in the directory
	output, err := exec.Command("ls", dir).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Output list items: ", string(output))

	// Convert []byte to string then make it list with delimiter \n
	items := strings.Split(string(output), "\n")

	// Check if list items already installed
	for _, item := range items {
		fmt.Println("Checking ", item)
		_, err := exec.Command("brew", "list", string(item)).Output()
		if err != nil {
			fmt.Println("Not installed ", item)
			// fmt.Println("Installing ", item)
			// _, err := exec.Command("brew", "install", string(item)).Output()
			// if err != nil {
			// 	log.Fatal(err)
			// }
		} else {
			fmt.Println("Already installed ", item)
		}
	}

}
