package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Model

	Name     string     `gorm:"UNIQUE;NOT NULL;UNIQUE_INDEX:name;" json:"name"`
	Mail     *string    `json:"Mail"`
	Password string     `gorm:"NOT NULL" json:"password"`
	Nickname *string    `json:"nickname"`
	LoginAt  *time.Time `json:"login_at"`
	Salt     string     `gorm:"NOT NULL" json:"salt"`
}

func GetUsers(maps interface{}) (users []User) {
	userDB.Where(maps).Find(&users)

	return
}

func GetUserCount(maps interface{}) (count int64) {
	userDB.Model(&User{}).Where(maps).Count(&count)

	return
}

func ExistUserByName(name string) bool {
	var count int64
	userDB.Model(&User{}).Where("name = ?", name).Count(&count)
	if int(count) > 0 {
		return true
	}
	return false
}

func ExistUserByNameAndPassword(name string, password string) bool {
	if !ExistUserByName(name) {
		return false
	}
	var user User
	userDB.Select("password", "salt").Where("name = ?", name).First(&user)
	// encrypted, _ := util.EncryptWithSalt(password, []byte(user.Salt))
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password+user.Salt))
	if err != nil {
		return false
	} else {
		return true
	}
}

func AddUser(name string, password string, salt string) bool {
	user := User{Name: name, Password: password, Salt: salt}
	result := userDB.Create(&user)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func AddUserItem(user *User) bool {
	result := userDB.Create(&user)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}
