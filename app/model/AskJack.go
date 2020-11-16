package model

type AskJackLog struct {
	IP       string
	Port     int
	Email    string
	Question string
}

func NewAskJackLog() *AskJackLog {
	return &AskJackLog{
		IP:       "",
		Port:     0000,
		Email:    "",
		Question: "",
	}
}
