package utils

import (
        "net/http"
        "encoding/json"
)

func Response(w http.ResponseWriter, body interface{}) {
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(body)

} 

func ResponseError(w http.ResponseWriter, err *ApplicationError) {
     w.Header().Set("Content-Type", "application/json")
     json.NewEncoder(w).Encode(err)

} 
