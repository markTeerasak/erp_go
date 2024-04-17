package repository

import "erp/models"

type UserRepository interface {
	CreateUser(user *models.User) error
	Login(user *models.User) error
}