package commands

import (
	"log"

	"github.com/limaJavier/secure/internal/encryption"
	"github.com/limaJavier/secure/internal/persistence"
	"github.com/spf13/cobra"
)

const KeyLen = 32

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Creates a new user",
	Long: `Creates a new user comprised of:
- Username
- Password`,
	Args: cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize repository
		userRepository, err := persistence.NewUserRepository()
		if err != nil {
			log.Fatalf("an unexpected error occurred: %v", err)
		}

		// Read user input
		username, err := readInput("Enter username", false)
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
			log.Fatalf("passwords don't match")
		}

		// Initialize encryption dependencies
		keyProvider := encryption.NewKeyProvider()
		encryptor := encryption.NewEncryptor(keyProvider)
		hasher := encryption.NewHasher(keyProvider)
		encoder := encryption.NewEncoder()

		// Generate user's random key
		key, _ := keyProvider.GenerateRandomKey(KeyLen)

		// Encrypt key using the password
		encryptedKey, err := encryptor.Encrypt([]byte(password), key)
		if err != nil {
			log.Fatalf("cannot encrypt key: %v", err)
		}

		// Encode key-encryption as a plain string
		encryptedKeyEncoding, err := encoder.EncodeEncryption(encryptedKey)
		if err != nil {
			log.Fatalf("cannot encrypt key: %v", err)
		}

		// Hash password
		passwordHash, err := hasher.Hash([]byte(password))
		if err != nil {
			log.Fatalf("cannot hash password: %v", err)
		}

		// Encode password-hash as a plain
		passwordHashEncoding, err := encoder.EncodeHash(passwordHash)
		if err != nil {
			log.Fatalf("cannot encode password-hash: %v", err)
		}

		// Create user
		user := persistence.User{
			Username: username,
			Password: passwordHashEncoding,
			Key:      encryptedKeyEncoding,
		}
		err = userRepository.Create(user)
		if err != nil {
			log.Fatal(err)
		}
	},
}
