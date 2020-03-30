package utils

import (
	"books-list/models"
	"encoding/json"
	"log"
	"net/http"
)

//LogFatal prints the fatal error that occurs in the console
func LogFatal(err error) {
	if err != nil {
		log.Println("A Fatal Error has occurred:")
		log.Fatal(err)
	}
}

func InternalError(w http.ResponseWriter, err error) {
	if err != nil {
		var error models.Error
		error.Message = "Server Error"
		SendError(w, http.StatusInternalServerError, error)
		return
	}
}


func SendError(w http.ResponseWriter, status int, err models.Error) {
	(w).Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
	//(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.WriteHeader(status)


	json.NewEncoder(w).Encode(err)
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	(w).Header().Set("Access-Control-Allow-Origin", "*") // Enable CORS
	//(w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	//(w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	json.NewEncoder(w).Encode(data)
}