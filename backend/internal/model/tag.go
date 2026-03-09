package model

import "time"

// 태그 모델 (API <-> DB)
type Tag struct {
	ID				int32		`json:"id" db:"id"`
	Name			string		`json:"name" db:"name"`
	Slug			string		`json:"slug" db:"slug"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at" db:"updated_at"`
}