package persistence

import (
	"fmt"

	"github.com/limaJavier/secure/internal/encryption"
	"gorm.io/gorm"
)

type EntryRepository interface {
	Create(entry Entry) error
	Retrieve() ([]Entry, error)
	Update(entry Entry) error
	Delete(id uint) error
}

func NewEntryRepository(user LoggedUser) (EntryRepository, error) {
	db, err := getDb()
	if err != nil {
		return nil, fmt.Errorf("cannot initialize entry-repository: %v", err)
	}

	return &entryRepository{
		db:        db,
		user:      user,
		encryptor: encryption.NewEncryptor(encryption.NewKeyProvider()),
		encoder:   encryption.NewEncoder(),
	}, nil
}

type entryRepository struct {
	db        *gorm.DB
	user      LoggedUser
	encryptor encryption.Encryptor
	encoder   encryption.Encoder
}

func (repository *entryRepository) Create(entry Entry) error {
	// Assert user-id consistency
	if repository.user.Username != entry.Username {
		panic("cannot create entry: user-id does not match logged one")
	}

	entry, err := repository.encrypt(entry) // Encrypt entry
	if err != nil {
		return fmt.Errorf("cannot create entry: %v", err)
	}
	result := repository.db.Create(&entry) // Store on DB
	return result.Error
}
func (repository *entryRepository) Retrieve() ([]Entry, error) {
	// Query DB for user's entries
	encryptedEntries := make([]Entry, 0)
	result := repository.db.Where("username = ?", repository.user.Username).Find(&encryptedEntries)
	if result.Error != nil {
		return nil, fmt.Errorf("cannot retrieve entries: %v", result.Error)
	}

	// Decrypt entries
	entries := make([]Entry, 0, len(encryptedEntries))
	for _, entry := range encryptedEntries {
		entry, err := repository.decrypt(entry)
		if err != nil {
			return nil, fmt.Errorf("cannot retrieve entries: %v", err)
		}
		entries = append(entries, entry)
	}
	return entries, nil
}

func (repository *entryRepository) Update(entry Entry) error {
	// Query DB for entry with given id
	var storedEntry Entry
	result := repository.db.First(&storedEntry, entry.ID)
	if result.Error != nil {
		return fmt.Errorf("cannot update entry: %v", result.Error)
	}

	storedEntry, err := repository.decrypt(storedEntry) // Decrypt storedEntry
	if err != nil {
		return fmt.Errorf("cannot update entry: %v", err)
	}

	// Verify if any of the properties is an empty string, if so then don't update it
	if entry.Name == "" {
		entry.Name = storedEntry.Name
	}
	if entry.Description == "" {
		entry.Description = storedEntry.Description
	}
	if entry.Password == "" {
		entry.Password = storedEntry.Password
	}

	// Ensure logged-user cannot update other users' entries
	if repository.user.Username != entry.Username {
		return fmt.Errorf("cannot delete entry: entry with id %v does not belong to user %v", entry.ID, repository.user.Username)
	}

	entry, err = repository.encrypt(entry) // Encrypt entry
	if err != nil {
		return fmt.Errorf("cannot update entry: %v", err)
	}

	return repository.db.Save(&entry).Error // Update on DB
}

func (repository *entryRepository) Delete(id uint) error {
	// Query DB for entry with given id
	var entry Entry
	result := repository.db.First(&entry, id)
	if result.Error != nil {
		return fmt.Errorf("cannot delete entry: %v", result.Error)
	}

	// Ensure logged-user cannot delete other users' entries
	if repository.user.Username != entry.Username {
		return fmt.Errorf("cannot delete entry: entry with id %v does not belong to user with id %v", entry.ID, repository.user.Username)
	}

	return repository.db.Delete(&entry).Error // Delete entry from DB
}

func (repository *entryRepository) encrypt(entry Entry) (Entry, error) {
	encryptionData, err := repository.encryptor.Encrypt(repository.user.Key, []byte(entry.Name))
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	encoding, err := repository.encoder.EncodeEncryption(encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	entry.Name = encoding

	encryptionData, err = repository.encryptor.Encrypt(repository.user.Key, []byte(entry.Description))
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	encoding, err = repository.encoder.EncodeEncryption(encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	entry.Description = encoding

	encryptionData, err = repository.encryptor.Encrypt(repository.user.Key, []byte(entry.Password))
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	encoding, err = repository.encoder.EncodeEncryption(encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot encrypt entry: %v", err)
	}
	entry.Password = encoding

	return entry, nil
}

func (repository *entryRepository) decrypt(entry Entry) (Entry, error) {
	encryptionData, err := repository.encoder.DecodeEncryption(entry.Name)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	data, err := repository.encryptor.Decrypt(repository.user.Key, encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	entry.Name = string(data)

	encryptionData, err = repository.encoder.DecodeEncryption(entry.Description)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	data, err = repository.encryptor.Decrypt(repository.user.Key, encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	entry.Description = string(data)

	encryptionData, err = repository.encoder.DecodeEncryption(entry.Password)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	data, err = repository.encryptor.Decrypt(repository.user.Key, encryptionData)
	if err != nil {
		return Entry{}, fmt.Errorf("cannot decrypt entry: %v", err)
	}
	entry.Password = string(data)

	return entry, nil
}
