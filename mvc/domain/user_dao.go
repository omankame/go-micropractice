package domain

import (
        "fmt"
        "github.com/omankame/mvc/utils"
        "net/http"
        "log"
)

var (
      users = map[int]*User {
              1: {ID: 1, FirstName: "Onkar", LastName: "Mankame", Email: "myemail@gmail.com"},
              }  
      UserDao userDaoInterface //interface
 )

type userDaoInterface interface {     
     GetUser(int) (*User, *utils.ApplicationError)  //function in interface
}

type userDao struct{}  // structure

func init() {
     UserDao = &userDao{}     //structure instance and same is interface instance

}

func (u *userDao) GetUser(uid int) (*User, *utils.ApplicationError) {    //structure implement interface
      log.Println("We are accessing the database")
      user, ok  := users[uid]
      if ok { 
         return user, nil
      } else {
        return nil, &utils.ApplicationError { 
                    Message:    fmt.Sprintf("user %v does not exists", uid),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	      }
      }
     
}    
       

     
