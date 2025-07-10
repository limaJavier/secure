package persistence

type User struct {
	Username      string `gorm:"primaryKey"`
	Password, Key string
	Entries       []Entry `gorm:"foreignKey:Username"`
}
