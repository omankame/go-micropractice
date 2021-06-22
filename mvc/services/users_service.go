package services

import (
         "github.com/omankame/mvc/utils"
         "github.com/omankame/mvc/domain"
)

type userService struct{}

var (
     UsersService userService
)

func(u *userService) GetUser(id int) (*domain.User, *utils.ApplicationError) {
     user, err := domain.UserDao.GetUser(id)
     if err != nil {
        return nil, err
     }

     return user, nil
}



