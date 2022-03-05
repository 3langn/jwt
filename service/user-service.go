package service

import (
	"github.com/3langn/learn-go/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type IUserService interface {
	InsertUser(user models.User) models.User
	UpdateUser(user models.User) models.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) models.User
	ProfileUser(userID string) models.User
}

type UserService struct {
}

func (m UserService) InsertUser(user models.User) models.User {
	user.Password = hashAndSalt([]byte(user.Password))
	models.GetDb().Save(&user)
	return user
}

func (m UserService) UpdateUser(user models.User) models.User {
	if user.Password != "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var tempUser models.User
		models.GetDb().Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}

	models.GetDb().Save(&user)
	return user
}

func (m UserService) VerifyCredential(email string, password string) interface{} {
	var user models.User
	res := models.GetDb().Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (m UserService) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user models.User
	return models.GetDb().Where("email = ?", email).Take(&user)
}

func (m UserService) FindByEmail(email string) models.User {
	var user models.User
	models.GetDb().Where("email = ?", email).Take(&user)
	return user
}

func (m UserService) ProfileUser(userID string) models.User {
	var user models.User
	models.GetDb().Preload("Books").Preload("Books.models.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
