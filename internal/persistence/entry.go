package persistence

type Entry struct {
	ID                          uint
	Name, Description, Password string
	UserID                      uint
}
