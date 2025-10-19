package models

// Topic represents a Kafka topic
type Topic struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Description string `json:"description"`
}
