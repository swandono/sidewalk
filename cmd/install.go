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
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	Run:  install,
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
	dir := cloneRepo(repo)
	defer os.RemoveAll(dir)

	listDir, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	for _, v := range listDir {
		fmt.Println(v.Name())
	}

	fmt.Println("................")
	data := getYaml(dir)
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
		if v.Exe != "" && v.Config != nil && v.Dir != "" {
			home, _ := os.UserHomeDir()
			target := home + "/" + v.Dir
			fmt.Println("Directory: ", target)

			// Create a directory
			err := os.MkdirAll(target, 0755)
			if err != nil {
				fmt.Println("Directory already exist")
			}

			// Copy file
			fmt.Println("Config:")
			for _, v := range v.Config {
				source := dir + "/" + v
				err := oss.init(target, source)
				if err != nil {
					fmt.Printf(" - %v: Copy Failed\n", v)
				} else {
					fmt.Printf(" - %v: Copy Successfull\n", v)
				}
			}
		}
	}
}
