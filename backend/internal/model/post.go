package model

import (
	"time"
)

// 도메인 엔티티 (API <-> DB)
type Post struct {
	ID				int32		`json:"id"`
	AccountID 		string		`json:"account_id"`
	Slug			string		`json:"slug"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`
	Thumbnail		string		`json:"thumbnail"`
	CategoryID		int16		`json:"category_id"`
	PublishedAt		time.Time	`json:"published_at"`
	IsPrivate		bool		`json:"is_private"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
	DeletedAt		time.Time	`json:"deleted_at"`
}

// 요청 DTO (Client -> API)

type CreatePostRequest struct {
	Slug			string		`json:"slug" binding:"required,max=150"`
	Title			string		`json:"title" binding:"required,max=100"`
	Content			string		`json:"content" binding:"required,max=100000"`
	Thumbnail		string		`json:"thumbnail" binding:"max=2048"`
	CategoryID		string		`json:"category_id"`
	TagsID			[]string	`json:"tags_id"`
	IsPublish		string		`json:"is_publish"`
	IsPrivate		string		`json:"is_private"`
}

type UpdatePostRequest struct {
	Slug			string		`json:"slug" binding:"max=150"`
	Title			string		`json:"title" binding:"max=100"`
	Content			string		`json:"content" binding:"max=100000"`
	Thumbnail		string		`json:"thumbnail" binding:"max=2048"`
	CategoryID		string		`json:"category_id"`
	TagsID			[]string	`json:"tags_id"`
	IsPublish		string		`json:"is_publish"`
	IsPrivate		string		`json:"is_private"`
}

// 응답 DTO (Client <- API)

type CreatePostResponse struct {
	ID				int32		`json:"id"`
	Slug			string		`json:"slug"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`
	Thumbnail		string		`json:"thumbnail"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
	IsPublish		bool		`json:"is_publish"`
	IsPrivate		bool		`json:"is_private"`
	CategoryID		ReadCategoryResponse	`json:"category"`
	TagID			[]ReadTagResponse		`json:"tags"`
}

type UpdatedPostResponse struct {
	Result struct {
		Data		CreatePostResponse	`json:"data"`
		Diff		CommonUpdateDiff	`json:"diff"`
	} `json:"result"`
}