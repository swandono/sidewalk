/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check installed software",
	Long:  `Check installed software in your system. Compare with the software in the repository.`,
	Args:  cobra.ExactArgs(1),
	Run:   check,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check(cmd *cobra.Command, args []string) {
	oss := getOss()
	if oss == nil {
		log.Fatal("OS not supported")
	}

	repo := checkRepo(args[0])
	dir, err := cloneRepo(repo)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)

	fmt.Println("................")
	data, err := getYaml(dir)
	if err != nil {
		log.Fatal(err)
	}
	for k, v := range data {
		fmt.Printf("\n")
		fmt.Printf("Name: %v \n", k)
		if v.Exe != "" {
			_, err := oss.check(k)
			fmt.Printf("Executable: %v \n", v.Exe)
			if err != nil {
				fmt.Println(" - Not installed")
			} else {
				fmt.Println(" - Already installed")
			}
		}
		if v.Exe != "" && v.Dependencies != nil {
			fmt.Println("Dependencies:")
			for _, dep := range v.Dependencies {
				_, err := oss.check(dep)
				if err != nil {
					fmt.Printf(" - %v: Not installed\n", dep)
				} else {
					fmt.Printf(" - %v: Already installed\n", dep)
				}
			}
		}
	}
}
