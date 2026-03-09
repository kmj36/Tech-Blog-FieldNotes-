package model

import "time"

// 게시물 모델
type Post struct {
	ID				int32		`json:"id" db:"id"`
	AccountID 		string		`json:"account_id" db:"account_id"`
	Slug			string		`json:"slug" db:"slug"`
	Title			string		`json:"title" db:"title"`
	Content			string		`json:"content" db:"content"`
	Thumbnail		string		`json:"thumbnail" db:"thumbnail"`
	CategoryID		int16		`json:"category_id" db:"category_id"`
	PublishedAt		time.Time	`json:"published_at" db:"published_at"`
	IsPrivate		bool		`json:"is_private" db:"is_private"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at" db:"updated_at"`
	DeletedAt		time.Time	`json:"deleted_at" db:"deleted_at"`
}