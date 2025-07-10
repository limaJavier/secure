package commands

import (
	"log"
	"strconv"

	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <entry-id>",
	Short: "Deletes an existing password entry",
	Long:  "Deletes an existing password entry given its ID.",
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

		err = repository.Delete(uint(entryId))
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}
	},
}
