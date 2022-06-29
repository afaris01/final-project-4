package repositories

import (
	"fmt"
	"final-project-4/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	Save(input models.User) (models.User, error)
	CheckSameEmail(email string) (models.User, error)
	GetByEmail(email string) (models.User, error)
	GetByID(ID int) (models.User, error)
	Update(ID int, user models.User) (models.User, error)
	Delete(ID int) (bool, error)
	UpdateSaldo(IDUser int, saldo int) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Save(user models.User) (models.User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) CheckSameEmail(email string) (models.User, error) {
	userSame := models.User{}

	err := r.db.Where("email = ?", email).Find(&userSame).Error

	if err != nil {
		return models.User{}, err
	}

	return userSame, nil
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	userResult := models.User{}

	err := r.db.Where("email = ?", email).Find(&userResult).Error

	if err != nil {
		return models.User{}, err
	}

	return userResult, nil
}

func (r *userRepository) GetByID(ID int) (models.User, error) {
	userResult := models.User{}

	err := r.db.Where("id = ?", ID).Find(&userResult).Error

	if err != nil {
		return models.User{}, err
	}

	return userResult, nil
}

func (r *userRepository) Update(ID int, user models.User) (models.User, error) {
	err := r.db.Where("id = ?", ID).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) Delete(ID int) (bool, error) {
	userDeleted := models.User{
		ID: ID,
	}

	err := r.db.Where("id = ?", ID).Delete(userDeleted).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *userRepository) UpdateSaldo(IDUser int, saldo int) (models.User, error) {
	userQuery := models.User{}

	err := r.db.Where("id = ?", IDUser).Find(&userQuery).Error

	if err != nil {
		return models.User{}, err
	}

	saldoSementara := userQuery.Balance - saldo

	err = r.db.Where("id = ?", IDUser).Update("balance", saldoSementara).Error

	if err != nil {
		return models.User{}, err
	}
	fmt.Println("update saldo")
	fmt.Println(saldoSementara)

	return userQuery, nil

}
