package models

import "github.com/google/uuid"

type Item struct {
	Id    uuid.UUID `json:"id"`
	Value string    `json:"value"`
	Done  bool      `json:"done"`
}

type ItemFromDB struct {
	Id     uuid.UUID `json:"id"`
	Value  string    `json:"value"`
	Author string    `json:"author"`
	Done   bool      `json:"done"`
}
