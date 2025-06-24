package commands

import (
	"log"

	"github.com/spf13/cobra"
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
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(updateCmd)
	rootCmd.AddCommand(retrieveCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
