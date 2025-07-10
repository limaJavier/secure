package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var rootCmd = &cobra.Command{
	Use:   "secure",
	Short: "Secure is an encrypted password manager",
	Long: `
  ___                      
 / __| ___ __ _  _ _ _ ___ 
 \__ \/ -_) _| || | '_/ -_)
 |___/\___\__|\_,_|_| \___|
                           

A fast and reliable application to manage your passwords`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(retrieveCmd)
}

func readInput(message string, hidden bool) (string, error) {
	fmt.Printf("%v: ", message)

	if hidden {
		input, err := term.ReadPassword(int(syscall.Stdin))
		defer fmt.Println() // Print new line since ReadPassword does not include "/n"
		if err != nil {
			return "", fmt.Errorf("failed to read input: %v", err)
		}
		return string(input), nil
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input: %v", err)
	}
	input = input[:len(input)-1] // Remove newline character from the end
	return input, err
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
