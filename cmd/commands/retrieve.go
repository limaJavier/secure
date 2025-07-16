package commands

import (
	"fmt"
	"log"

	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieves all password entries",
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

		// Retrieve entries
		entries, err := repository.Retrieve()
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

		for _, entry := range entries {
			fmt.Println("--------------------")
			fmt.Printf("ID: %v\n", entry.ID)
			fmt.Printf("Name: %v\n", entry.Name)
			fmt.Printf("Description: %v\n", entry.Description)
			fmt.Printf("Password: %v\n", entry.Password)
		}
		fmt.Println("--------------------")
	},
}
