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

// installCmd represents the install command
var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install software from the repository",
	Long:  `Install software from the repository. It will install the software and dependencies.`,
	Args:  cobra.ExactArgs(1),
	Run:   install,
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func install(cmd *cobra.Command, args []string) {
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
				err := oss.install(k)
				fmt.Printf("Executable: %v \n", v.Exe)
				if err != nil {
					fmt.Println(" - Installing Failed")
				} else {
					fmt.Println(" - Installing Successfull")
				}
			} else {
				fmt.Println(" - Already installed")
			}
		}
		if v.Exe != "" && v.Dependencies != nil {
			fmt.Println("Dependencies:")
			for _, dep := range v.Dependencies {
				_, err := oss.check(dep)
				if err != nil {
					err := oss.install(dep)
					if err != nil {
						fmt.Printf(" - %v: Installing Failed\n", dep)
					} else {
						fmt.Printf(" - %v: Installing Successfull\n", dep)
					}
				} else {
					fmt.Printf(" - %v: Already installed\n", dep)
				}
			}
		}
	}
}
