package sqlite

import (
	"fmt"
	user "go-hex/internal/user/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func OpenDatabase(name string) *gorm.DB {
	db, err := gorm.Open("sqlite3", name)
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&user.User{})
	return db
}

type userRepository struct {
	db *gorm.DB
}

func NewSQLiteUserRepository(db *gorm.DB) user.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *user.User) error {
	if dbc := r.db.Create(user); dbc.Error != nil { return fmt.Errorf("database error: %w", dbc.Error) }
	return nil
}

func (r *userRepository) FindById(id string) (*user.User, error) {
	var u user.User
	if dbc := r.db.Where(user.User{ID: id}).First(&u); dbc.Error != nil {
		return nil, fmt.Errorf("database error: %w", dbc.Error)
	}
	return &u, nil
}

func (r *userRepository) FindAll() ([]*user.User, error) {
	users := make([]*user.User, 5)
	if dbc := r.db.Find(&users); dbc.Error != nil { return nil, fmt.Errorf("database error: %w", dbc.Error)}
	return users, nil
}

func (r *userRepository) Save(user *user.User) error {
	if dbc := r.db.Save(&user); dbc.Error != nil { return fmt.Errorf("database error: %w", dbc.Error)}
	return nil
}
