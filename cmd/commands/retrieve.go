package commands

import (
	"log"

	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves all password entries",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Process logged-user...
		user := persistence.LoggedUser{
			Username: "",
			Key:      []byte("key"),
		}

		repository, err := persistence.NewEntryRepository(user) // Initialize repository
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

		entries, err := repository.Retrieve() // Retrieve entries
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

		// TODO: Properly print entries
		log.Println(entries)
	},
}
