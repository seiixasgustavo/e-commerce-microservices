package v1

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct for auth
type User struct {
	gorm.Model
	Username string `gorm:"unique;not_null"`
	Email    string `gorm:"unique;not_null"`
	Password string `gorm:"size:255"`
}

func (u *User) hash() error {
	cryptPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), 10)

	if err != nil {
		return err
	}

	u.Password = string(cryptPass)

	return nil
}

// ValidPassword checks if the ...
func (u *User) ValidPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}

// Create save an instance of the user to the database
func (u *User) Create(db *gorm.DB) error {
	hashErr := u.hash()

	if hashErr != nil {
		return hashErr
	}

	err := db.Save(&u).Error

	if err != nil {
		return err
	}

	return nil
}

// Update an instance of the user to the database
func (u *User) Update(db *gorm.DB, id uint) error {
	var user User
	err := db.Where("id = ?", u.ID).First(&user).Error

	if err != nil {
		return err
	}

	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	user.hash()

	saveErr := db.Save(&user).Error

	if saveErr != nil {
		return saveErr
	}

	return nil
}

// Delete deletes an instance of the user to the database
func (u *User) Delete(db *gorm.DB, id uint) error {
	err := db.Where("id = ?", id).Delete(User{}).Error

	if err != nil {
		return err
	}

	return nil
}

// ChangePassword changes the password of an instance of the user on the database
func (u *User) ChangePassword(db *gorm.DB, id uint) error {
	user, err := u.FindByPk(db, id)

	if err != nil {
		return err
	}

	user.Password = u.Password
	user.hash()

	saveErr := db.Save(&user).Error
	if saveErr != nil {
		return saveErr
	}

	return nil
}

// FindByPk returns the user that matches the id
func (u *User) FindByPk(db *gorm.DB, id uint) (*User, error) {
	var user User
	dbErr := db.Where("id = ?", id).First(&user).Error

	if dbErr != nil {
		return nil, dbErr
	}

	return &user, nil
}

// FindByUsername returns the user that matches the username
func (u *User) FindByUsername(db *gorm.DB, username string) (*User, error) {
	var user User
	dbErr := db.Where("Username = ?", username).First(&user).Error

	if dbErr != nil {
		return nil, dbErr
	}

	return &user, nil
}
