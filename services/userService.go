package services

import "gin_frame/models"

func GetUsers(where map[string]interface{}) *models.Users  {
	var users models.Users
	models.DB().Debug().Where(where).Unscoped().Find(&users)
	return &users
}