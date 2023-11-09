package app

import (
	"encoding/json"
	"encoding/xml"
	"log"
	"net/http"
)

func setResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Panic("Error in generating response ", err)
	}

}
func setResponseXML(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/xml")
	w.WriteHeader(code)
	err := xml.NewEncoder(w).Encode(data)
	if err != nil {
		log.Panic("Error in generating response ", err)
	}

}
