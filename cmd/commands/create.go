package commands

import (
	"log"

	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a new password entry",
	Long: `Creates a new password entry comprised of:
- Name
- Description
- Password`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Auth user
		user, err := auth()
		if err != nil {
			log.Fatalf("cannot auth user: %v", err)
		}

		// Initialize repository
		repository, err := persistence.NewEntryRepository(user)
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

		// Read user input
		name, err := readInput("Enter name", false)
		if err != nil {
			log.Fatal(err)
		}
		description, err := readInput("Enter description", false)
		if err != nil {
			log.Fatal(err)
		}
		password, err := readInput("Enter password", true)
		if err != nil {
			log.Fatal(err)
		}
		confirmedPassword, err := readInput("Confirm password", true)
		if err != nil {
			log.Fatal(err)
		} else if password != confirmedPassword {
			log.Fatal("passwords don't match")
		}

		// Create entry
		entry := persistence.Entry{
			Name:        name,
			Description: description,
			Password:    password,
			Username:    user.Username,
		}
		err = repository.Create(entry)
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}
	},
}
