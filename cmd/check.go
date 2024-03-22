/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"

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
	Run: check,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func check(cmd *cobra.Command, args []string) {
	oss := getOss()
	if oss == nil {
		log.Fatal("OS not supported")
	}

	data := getYaml()
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
		if v.Exe != "" && v.Config != nil {
			fmt.Println("Config:")
			for _, v := range v.Config {
				fmt.Printf(" - %v\n", v)
			}
		}
	}
}
