package commands

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"syscall"

	"github.com/limaJavier/secure/internal/encryption"
	"github.com/limaJavier/secure/internal/persistence"
	"golang.org/x/term"
)

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

func auth() (persistence.LoggedUser, error) {
	userRepository, err := persistence.NewUserRepository()
	if err != nil {
		log.Fatalf("an unexpected error occurred: %v", err)
	}

	username, err := readInput("Enter username", false)
	if err != nil {
		return persistence.LoggedUser{}, err
	}
	password, err := readInput("Enter password", true)
	if err != nil {
		return persistence.LoggedUser{}, err
	}

	user, err := userRepository.Retrieve(username)
	if err != nil {
		return persistence.LoggedUser{}, err
	}

	keyProvider := encryption.NewKeyProvider()
	encryptor := encryption.NewEncryptor(keyProvider)
	hasher := encryption.NewHasher(keyProvider)
	encoder := encryption.NewEncoder()

	passwordHash, err := encoder.DecodeHash(user.Password)
	if err != nil {
		return persistence.LoggedUser{}, err
	}

	match, err := hasher.Verify([]byte(password), passwordHash)
	if err != nil {
		return persistence.LoggedUser{}, err
	} else if !match {
		return persistence.LoggedUser{}, fmt.Errorf("wrong password")
	}

	encryptedKey, err := encoder.DecodeEncryption(user.Key)
	if err != nil {
		return persistence.LoggedUser{}, err
	}

	key, err := encryptor.Decrypt([]byte(password), encryptedKey)
	if err != nil {
		return persistence.LoggedUser{}, err
	}

	return persistence.LoggedUser{
		Username: user.Username,
		Key:      key,
	}, err
}
