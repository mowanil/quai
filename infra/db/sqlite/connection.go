package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	// TODO: connect to database and do real operations
	_, _ = gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
}
