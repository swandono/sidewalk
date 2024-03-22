/*
Copyright © 2024 GUNAWAN SWANDONO <me@swandono.com>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: install,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
	oss := getOss()
	if oss == nil {
		log.Fatal("OS not supported")
	}

	data := getYaml()
	for k, v := range data {
		fmt.Printf("\n")
		fmt.Printf("Name: %v \n", k)
		if v.Exe != "" {
			err := oss.install(k)
			fmt.Printf("Executable: %v \n", v.Exe)
			if err != nil {
				fmt.Println(" - Installing Failed")
			} else {
				fmt.Println(" - Installing Successfull")
			}
		}
		if v.Exe != "" && v.Dependencies != nil {
			fmt.Println("Dependencies:")
			for _, dep := range v.Dependencies {
				err := oss.install(dep)
				if err != nil {
					fmt.Printf(" - %v: Installing Failed\n", dep)
				} else {
					fmt.Printf(" - %v: Installing Successfull\n", dep)
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
