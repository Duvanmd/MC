package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendData(rw http.ResponseWriter, data interface{}, status int) {

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)

	// Este metodo me convirte una lista de user a tipo Json
	output, _ := json.Marshal(data)
	fmt.Fprintln(rw, string(output))
}

func sendError(rw http.ResponseWriter, status int) {
	rw.WriteHeader(status)
	fmt.Fprintln(rw, "Resouce Not Found")
}
