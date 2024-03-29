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

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update software from the repository",
	Long:  `Update software from the repository. It will update the software and dependencies.`,
	Args:  cobra.ExactArgs(1),
	Run:   update,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func update(cmd *cobra.Command, args []string) {
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

	listDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range listDir {
		fmt.Println(v.Name())
	}

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
			if err != nil {
				err := oss.update(k)
				fmt.Printf("Executable: %v \n", v.Exe)
				if err != nil {
					fmt.Println(" - Updating Failed")
				} else {
					fmt.Println(" - Updating Successfull")
				}
			} else {
				fmt.Println(" - Already up to date")
			}
		}
		if v.Exe != "" && v.Dependencies != nil {
			fmt.Println("Dependencies:")
			for _, dep := range v.Dependencies {
				_, err := oss.check(dep)
				if err != nil {
					err := oss.update(dep)
					if err != nil {
						fmt.Printf(" - %v: Updating Failed\n", dep)
					} else {
						fmt.Printf(" - %v: Updating Successfull\n", dep)
					}
				} else {
					fmt.Printf(" - %v: Already up to date\n", dep)
				}
			}
		}
	}
}
