package persistence

import (
	"testing"

	"github.com/limaJavier/secure/internal/encryption"
	"github.com/stretchr/testify/require"
)

func TestEntryRepository_Integration_I(t *testing.T) {
	//** Arrange
	db, err := getDb(true) // Initialize in-memory db
	require.NoError(t, err)
	user := LoggedUser{
		Username: "johndoe",
		Key:      []byte("random_key"),
	}
	// Initialize repository
	repository := &entryRepository{
		db:        db,
		user:      user,
		encryptor: encryption.NewEncryptor(encryption.NewKeyProvider()),
		encoder:   encryption.NewEncoder(),
	}
	entry := Entry{
		Username:    user.Username,
		Name:        "entry-name",
		Description: "entry-description",
		Password:    "entry-password",
	}

	//** Act and Assert
	err = repository.Create(entry) // Create entry
	require.NoError(t, err)

	entries, err := repository.Retrieve() // Retrieve all user's entries
	require.NoError(t, err)

	require.Len(t, entries, 1)
	require.Equal(t, entry.Name, entries[0].Name)               // Check name
	require.Equal(t, entry.Description, entries[0].Description) // Check description
	require.Equal(t, entry.Password, entries[0].Password)       // Check password

	updatedEntry := Entry{
		ID:          entries[0].ID,
		Username:    entry.Username,
		Name:        "entry-name-updated", // Update only the name
		Description: "",
		Password:    "",
	}
	err = repository.Update(updatedEntry) // Update entry
	require.NoError(t, err)

	entries, err = repository.Retrieve()
	require.NoError(t, err)
	require.Equal(t, updatedEntry.Name, entries[0].Name)        // Check name update
	require.Equal(t, entry.Description, entries[0].Description) // Check description remains the same
	require.Equal(t, entry.Password, entries[0].Password)       // Check password remains the same

	err = repository.Delete(entries[0].ID) // Delete entry
	require.NoError(t, err)

	entries, err = repository.Retrieve()
	require.NoError(t, err)
	require.Len(t, entries, 0) // Check there are no entries
}
