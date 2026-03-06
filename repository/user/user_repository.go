package users

import (
	user_service "github.com/andrewidianto12/Manufacture-Rental/service/user"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}
func (r *UserRepository) RegisterUser(user *user_service.User) error {
	user.ID = 0
	return r.db.Create(user).Error
}

func (r *UserRepository) LoginUser(email string, password string) (*user_service.User, error) {
	var user user_service.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) DeleteUser(ID uint) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Exec("DELETE FROM rentals WHERE user_id = ?", ID).Error; err != nil {
			return err
		}
		if err := tx.Delete(&user_service.User{}, ID).Error; err != nil {
			return err
		}
		return nil
	})
}
