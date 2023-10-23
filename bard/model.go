package bard

import "net/http"

type Chatbot struct {
	ReqID          int
	SNlM0e         string
	ConversationID string
	ResponseID     string
	ChoiceID       string
	Client         *http.Client
	SessionID      string
}

type Response struct {
	Content           string
	ConversationID    string
	ResponseID        string
	FactualityQueries []interface{}
	TextQuery         string
	Choices           []Choice
}

type Choice struct {
	ID      string
	Content string
}

const (
	HOST       = "bard.google.com"
	ORIGIN_URL = "https://" + HOST
	BASE_URL   = "https://" + HOST + "/"
	ASK_URL    = BASE_URL + "_/BardChatUi/data/assistant.lamda.BardFrontendService/StreamGenerate"
)
