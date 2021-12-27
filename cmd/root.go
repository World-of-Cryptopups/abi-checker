/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	simplefiletest "github.com/TheBoringDude/simple-filetest"
	"github.com/World-of-Cryptopups/abi-checker/cmd/internal"
	"github.com/World-of-Cryptopups/abi-checker/cmd/utils"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "abi-checker",
	Short: "Check eosio smart contract abi file.",
	Long:  `This tool checks your abi file if there are problem and shows it.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("No ABI file paths provided.")
		}

		for _, v := range args {
			// Check if the file exists or not.
			if !simplefiletest.FileExists(v) {
				if !strings.HasSuffix(v, ".abi") {
					log.Fatalf("%s is not an ABI file.\n", v)
				} else {
					log.Fatalf("%s does not exist in path.\n", v)
				}
			}

			// Open file
			data, err := os.ReadFile(v)
			if err != nil {
				log.Fatalf("Failed to read abi file: %s\n", v)

			}

			var ABI *internal.ABIProps

			// try to unmarshall it to ABI
			if err = json.Unmarshal(data, &ABI); err != nil {
				log.Fatalf("Failed to unmarshall the content of: %s\nError: %v", v, err)
			}

			// CHECK structs
			fmt.Printf("[%s] Checking tables...\n", v)
			for _, j := range ABI.Tables {
				if len(j.Name) > 12 {
					log.Fatalf("[%s] Invalid struct name length `%s`\n", v, j.Name)
				}

				if !utils.IsValidBase32(j.Name) {
					log.Fatalf("[%s] Name `%s` does not comply with the Base32 standard naming.", v, j.Name)
				}
			}

			// CHECK actions
			fmt.Printf("[%s] Checking actions...\n", v)
			for _, j := range ABI.Actions {
				if len(j.Name) > 12 {
					log.Fatalf("[%s] Invalid struct name length `%s`\n", v, j.Name)
				}

				if !utils.IsValidBase32(j.Name) {
					log.Fatalf("[%s] Name `%s` does not comply with the Base32 standard naming.", v, j.Name)
				}
			}

			fmt.Printf("-> [%s] IS A VALID ABI!\n\n", v)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.abi-checker.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
