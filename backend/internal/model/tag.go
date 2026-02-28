package model

import "time"

// 도메인 엔티티 (API <-> DB)
type Tag struct {
	ID				int32		`json:"id" db:"id"`
	Name			string		`json:"name" db:"name"`
	Slug			string		`json:"slug" db:"slug"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at" db:"updated_at"`
}

// 요청 DTO (Client -> API)

type CreateTagRequest struct {
	Name			string		`json:"name" binding:"required,max=100"`
	Slug			string		`json:"slug" binding:"required,max=150"`
}

type UpdateTagRequest struct {
	Name			string		`json:"name" binding:"max=100"`
	Slug			string		`json:"slug" binding:"max=150"`
}

// 응답 DTO (Client <- API)

type CreateTagResponse struct {
	ID				int32		`json:"id"`
	Name			string		`json:"name"`
	Slug			string		`json:"slug"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

type UpdateTagResponse struct {
	Result struct {
		Data	CreateTagResponse	`json:"data"`
		Diff	CommonUpdateDiff	`json:"diff"`
	} `json:"result"`
}

type ReadTagResponse struct {
	ID				int32		`json:"id"`
	Name			string		`json:"name"`
	Slug			string		`json:"slug"`
}