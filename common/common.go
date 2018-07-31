package common

import (
	"time"
)

type Store interface {
	Init() error
	StoreSecret(secret string, expiration time.Duration) (string, error)
	GetSecret(id string) (string, error)
}

type Secret struct {
	Id         string    `json:"id" binding:"required"`
	Content    string    `json:"content" binding:"required"`
	Expiration time.Time `json:"expiration" binding:"required"`
}
