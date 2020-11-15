package handler

import (
	"fmt"
	"gorm.io/gorm"
	"net"
	"net/http"
)
// ExtractData from POST request for Ask Jack
// Currently extracted data includes: ip, question, tags, email
func AddQuestionEntry(db *gorm.DB, r *http.Request) {
	ip, port, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		fmt.Println(ip)
		fmt.Println(port)
	}
}
