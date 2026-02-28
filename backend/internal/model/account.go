package model

import "time"

// 도메인 엔티티 (API <-> DB)
type Account struct {
	ID				int16		`json:"id" db:"id"`
	AccountID		string		`json:"account_id" db:"account_id"`
	PasswordHash	string		`json:"password_hash" db:"password_hash"`
	Nickname		string		`json:"nickname" db:"nickname"`
	AvatarURL		string		`json:"avatar_url" db:"avatar_url"`
	Role			string		`json:"role" db:"role"`
	Status			string		`json:"status" db:"status"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at" db:"updated_at"`
}

// 요청 DTO (Client -> API)

type LoginAccountRequest struct {
	AccountID		string		`json:"id" binding:"required,min=3,max=32"`
	Password		string		`json:"password" binding:"required,min=8,max=80"`
}

type CreateAccountRequest struct {
	AccountID		string		`json:"id" binding:"required,min=3,max=32"`
	Password		string		`json:"password" binding:"required,min=8,max=80"`
	Nickname		string		`json:"nickname" binding:"min=2,max=20"`
	AvatarURL		string		`json:"avatar_url" binding:"max=2048"`
}

type UpdateAccountRequest struct {
	Password		string		`json:"password" binding:"required,min=8,max=80"`
	Nickname		string		`json:"nickname" binding:"min=2,max=20"`
	AvatarURL		string		`json:"avatar_url" binding:"max=2048"`
}

type DeleteAccountRequest struct {
	AccountID		string		`json:"id" binding:"required,min=3,max=32"`
}

// 응답 DTO (Client <- API)

type CreateAccountData struct {
    ID             string		`json:"id"`
    PasswordStatus string		`json:"password_status"`
    Nickname       string		`json:"nickname"`
    AvatarURL      string		`json:"avatar_url"`
}

type UpdateAccountResponse struct {
	Result struct {
		Data		CreateAccountData	`json:"data"`
		Diff		CommonUpdateDiff	`json:"diff"`
	} `json:"result"`
}

type ReadAccountResponse struct {
	ID				int16		`json:"id"`
	AccountID		string		`json:"account_id"`
	Nickname		string		`json:"nickname"`
	AvatarURL		string		`json:"avatar_url"`
	Role			string		`json:"role"`
	Status			string		`json:"status"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}