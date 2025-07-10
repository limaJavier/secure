package commands

import (
	"log"
	"strconv"

	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update <entry-id>",
	Short: "Updates an existing password entry",
	Long:  "Updates an existing password entry's properties given its ID. If the new value is an empty string the property won't be updated.",
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Validate arg
		entryId, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("entry-id argument must be of type unsigned-integer")
		}

		user, err := auth()
		if err != nil {
			log.Fatalf("cannot auth user: %v", err)
		}

		repository, err := persistence.NewEntryRepository(user) // Initialize repository
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

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
		if password != "" {
			confirmedPassword, err := readInput("Confirm password", true)
			if err != nil {
				log.Fatal(err)
			} else if password != confirmedPassword {
				log.Fatal("passwords don't match")
			}
		}

		entry := persistence.Entry{
			ID:          uint(entryId),
			Name:        name,
			Description: description,
			Password:    password,
			Username:    user.Username,
		}

		repository.Update(entry)
	},
}
