package service

import (
	"github.com/3langn/learn-go/dto"
	"github.com/3langn/learn-go/models"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) models.User
	FindByEmail(email string) models.User
	IsDuplicateEmail(email string) bool
}

var userService = new(UserService)

type AuthService struct {
}
func (service AuthService) VerifyCredential(email string, password string) interface{} {
	res :=  userService.VerifyCredential(email, password)
	log.Printf("%v\n", res)
	if v, ok := res.(models.User); ok {
		comparedPassword := comparePassword(v.Password, []byte(password))
		if v.Email == email && comparedPassword {
			return res
		}
		return false
	}
	return false
}

func (service AuthService) CreateUser(user dto.RegisterDTO) models.User {
	userToCreate := models.User{}
	err := smapping.FillStruct(&userToCreate, smapping.MapFields(&user))
	if err != nil {
		log.Fatalf("Failed map %v", err)
	}
	res := userService.InsertUser(userToCreate)
	return res
}

func (service AuthService) FindByEmail(email string) models.User {
	return userService.FindByEmail(email)
}

func (service AuthService) IsDuplicateEmail(email string) bool {
	res := userService.IsDuplicateEmail(email)
	return !(res.Error == nil)
}

func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
