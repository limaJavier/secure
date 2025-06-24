package persistence

type User struct {
	ID                      uint
	Username, Password, Key string
	Entries                 []Entry
}
