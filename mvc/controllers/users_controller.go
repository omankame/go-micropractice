package controllers

import (
        "net/http"
        "strconv"
        "strings"
        "github.com/omankame/mvc/services"
        "github.com/omankame/mvc/utils"
)

func GetUser(w http.ResponseWriter, req *http.Request) {
     id := strings.TrimLeft(req.URL.Path, "/users/")  //get the userid

     userId, err := strconv.Atoi(id)
     if err != nil {
        apiErr := &utils.ApplicationError {
                  Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
         _ = apiErr
         http.Error(w, err.Error(), http.StatusBadRequest)
         return
     }    


     user, apiErr := services.UsersService.GetUser(userId)
     if apiErr != nil {
        utils.RespondError(w,apiErr) 
        return
     }

     utils.Respond(w,user)

     
     
} 
