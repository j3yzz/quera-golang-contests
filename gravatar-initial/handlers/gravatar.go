package handlers

import (
	"crypto/md5"
	hex2 "encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
)

type GravatarResponse struct {
	Ok          bool   `json:"ok"`
	GravatarUrl string `json:"gravatar_url"`
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

const GRAVATAR_BASE_URL = "https://www.gravatar.com/avatar/"

func HandleGravatarRequest(w http.ResponseWriter, r *http.Request) {
	var errorResponse ErrorResponse
	var gr GravatarResponse
	query := r.URL.Query()
	email := query.Get("email")
	if len(email) == 0 {
		errorResponse.Ok = false
		errorResponse.Message = "No email address provided"
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal, Error: %s", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	_, err := mail.ParseAddress(email)

	if err != nil {
		errorResponse.Ok = false
		errorResponse.Message = "Invalid email address"
		w.Header().Set("Content-Type", "application/json")
		jsonResp, err := json.Marshal(errorResponse)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal, Error: %s", err)
		}
		w.WriteHeader(http.StatusBadRequest)
		w.Write(jsonResp)
		return
	}

	hash := md5.Sum([]byte(email))
	hex := hex2.EncodeToString(hash[:])
	gr.Ok = true
	gr.GravatarUrl = GRAVATAR_BASE_URL + hex
	w.Header().Set("Content-Type", "application/json")
	jsonResp, err := json.Marshal(gr)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal, Error: %s", err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
	return
}
