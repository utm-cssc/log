package handler

import (
	"gorm.io/gorm"
	"net/http"
)
// ExtractData from POST request for Ask Jack
// Currently extracted data includes: ip, question, tags, email
func ExtractData(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

}
