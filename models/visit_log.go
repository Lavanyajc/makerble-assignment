package models

type VisitLog struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	Endpoint  string `json:"endpoint"`
	Method    string `json:"method"`
	Timestamp string `json:"timestamp"`
}

