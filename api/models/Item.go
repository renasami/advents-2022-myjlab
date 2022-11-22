package models

import "github.com/google/uuid"

type Item struct{
	Id uuid.UUID `json:"id"`
	Value string `json:"value"`
	User string `json:"user"`
}