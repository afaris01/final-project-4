package services

import (
	"errors"
	"final-project-4/models"
	"final-project-4/models/inputs"
	"final-project-4/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(userInput inputs.RegisterUserInput) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	TopUp(saldo int, IDUser int) (models.User, error)
	UpdateUser(ID int, input inputs.UpdateUserInput) (models.User, error)
	GetUserByID(ID int) (models.User, error)
	DeleteUser(ID int) (models.User, error)
	CheckUserAdmin(ID int) (bool, error)
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *userService {
	return &userService{userRepository}
}

func (s *userService) CreateUser(input inputs.RegisterUserInput) (models.User, error) {
	newUser := models.User{}
	newUser.Email = input.Email
	newUser.FullName = input.FullName
	newUser.Role = "customer"
	newUser.Balance = 0

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil {
		return models.User{}, err
	}

	newUser.Password = string(passwordHash)

	checkSameUser, err := s.userRepository.CheckSameEmail(input.Email)

	if err != nil {
		return models.User{}, err
	}

	if checkSameUser.ID != 0 {
		return models.User{}, errors.New("Email already registered!")
	}

	createdUser, err := s.userRepository.Save(newUser)

	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (s *userService) GetUserByEmail(email string) (models.User, error) {
	userResult, err := s.userRepository.GetByEmail(email)

	if err != nil {
		return models.User{}, err
	}

	return userResult, nil
}

func (s *userService) TopUp(saldo int, IDUser int) (models.User, error) {
	userResult, err := s.userRepository.GetByID(IDUser)

	if err != nil {
		return models.User{}, err
	}

	if userResult.ID == 0 {
		return models.User{}, errors.New("user not found!")
	}

	saldoTerkini := userResult.Balance + saldo

	updated, err := s.userRepository.Update(IDUser, models.User{Balance: saldoTerkini})

	if err != nil {
		return models.User{}, err
	}

	return updated, nil
}

func (s *userService) GetUserByID(ID int) (models.User, error) {
	user, err := s.userRepository.GetByID(ID)

	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return models.User{}, nil
	}

	return user, nil
}

func (s *userService) UpdateUser(ID int, input inputs.UpdateUserInput) (models.User, error) {
	userResult, err := s.userRepository.GetByID(ID)

	if err != nil {
		return models.User{}, err
	}

	if userResult.ID == 0 {
		return models.User{}, errors.New("user not found!")
	}

	updatedUser := models.User{
		FullName: input.FullName,
		Email:    input.Email,
	}

	userUpdate, err := s.userRepository.Update(ID, updatedUser)

	if err != nil {
		return userUpdate, err
	}

	return userUpdate, nil
}

func (s *userService) DeleteUser(ID int) (models.User, error) {

	userdata, err := s.GetUserByID(ID)

	if err != nil {
		return models.User{}, err
	}

	if userdata.Role == "admin" {
		return models.User{}, errors.New("Admin can not destroy self!")
	}

	_, err = s.userRepository.Delete(ID)

	if err != nil {
		return models.User{}, err
	}

	return models.User{}, nil

}

func (s *userService) CheckUserAdmin(ID int) (bool, error) {
	userdata, err := s.GetUserByID(ID)

	if err != nil {
		return false, err
	}

	if userdata.Role != "admin" {
		return false, nil
	}

	return true, nil
}
