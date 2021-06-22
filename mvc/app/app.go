package app

import (
         "net/http"
         "github.com/omankame/mvc/controllers"
)

func StartApp() {
//     mapUrls()
     http.HandleFunc("/users/", controllers.GetUser)
     http.ListenAndServe(":8080", nil)

}
