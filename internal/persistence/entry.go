package persistence

type Entry struct {
	ID                          uint
	Name, Description, Password string
	Username                    string
}
