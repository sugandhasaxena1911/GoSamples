package utils

import (
	"encoding/json"
	"github.com/sugandhasaxena1911/GoSamples/main/models"
	"net/http"
)

func RespondError(w http.ResponseWriter, statuscode int, err models.Error) {
	w.WriteHeader(statuscode)
	json.NewEncoder(w).Encode(err)
	return
}
