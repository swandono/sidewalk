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
	Run: command,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func command(cmd *cobra.Command, args []string) {
	oss := getOss()
	if oss == nil {
		log.Fatal("OS not supported")
	}

	data := getYaml()
	for k, v := range data {
		fmt.Printf("\n")
		fmt.Printf("Name: %v \n", k)
		if v.(map[interface{}]interface{})["exe"] != nil {
			_, err := oss.check(v.(map[interface{}]interface{})["exe"].(string))
			fmt.Printf("executable: %v \n", v.(map[interface{}]interface{})["exe"].(string))
			if err != nil {
				fmt.Println(" - Not installed")
			} else {
				fmt.Println(" - Already installed")
			}
		}
		if v.(map[interface{}]interface{})["exe"] != nil && v.(map[interface{}]interface{})["dependencies"] != nil {
			fmt.Println("Dependencies:")
			for _, dep := range v.(map[interface{}]interface{})["dependencies"].([]interface{}) {
				_, err := oss.check(dep.(string))
				if err != nil {
					fmt.Printf(" - %v: Not installed\n", dep.(string))
				} else {
					fmt.Printf(" - %v: Already installed\n", dep.(string))
				}
			}
		}
		if v.(map[interface{}]interface{})["exe"] != nil && v.(map[interface{}]interface{})["config"] != nil {
			fmt.Println("Config:")
			for _, v := range v.(map[interface{}]interface{})["config"].([]interface{}) {
				fmt.Printf(" - %v\n", v)
			}
		}
	}
}
