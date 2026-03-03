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

type ReadPostRequest struct {
	Page			int32		`json:"page" binding:"min=1,max=10000"`
	PageSize		int32		`json:"page_size" binding:"max=200"`
	Search			string		`json:"search" binding:"max=100"`
	CategorySlug	string		`json:"category_slug" binding:"max=100"`
	TagSlugs		[]string	`json:"tag_slugs" binding:"max=50"`
	TagSort			string		`json:"tag_sort"`
	SortBy			string		`json:"sort_by"`
	SortDir			string		`json:"sort_dir"`
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

type ReadPostsResponse struct {
	Result struct {
		Meta		PostMetadataObject	`json:"meta"`
		Data 		[]CommonPostObject	`json:"data"`
	} `json:"result"`
}

type ReadEachPostResponse struct {
	ID				int32		`json:"id"`
	Slug			string		`json:"slug"`
	Title			string		`json:"title"`
	Content			string		`json:"content"`
	Thumbnail		string		`json:"thumbnail"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
	CategoryID		ReadCategoryResponse	`json:"category"`
	TagID			[]ReadTagResponse		`json:"tags"`
}

type UpdatedPostResponse struct {
	Result struct {
		Data		CreatePostResponse	`json:"data"`
		Diff		CommonUpdateDiff	`json:"diff"`
	} `json:"result"`
}

type PostMetadataObject struct {
	Pagination		struct {
		Page		int32		`json:"page"`
		PageSize	int32		`json:"page_size"`
		Total		int32		`json:"total"`
		TotalPages	int32		`json:"total_pages"`
		HasNextPage	bool		`json:"has_nextpage"`
		HasPrevPage bool		`json:"has_prevpage"`
	}	`json:"pagination"`
	Sort			struct {
		SortBy		string		`json:"by"`
		SortDir		string		`json:"desc"`
	}	`json:"sort"`
	Filters			struct {
		Search		string		`json:"search"`
		CategorySlug	string	`json:"category_slug"`
		TagSlugs		[]string	`json:"tag_slugs"`
	}	`json:"filters"`
}
type CommonPostObject struct {
	ID				int32		`json:"id"`
	Slug			string		`json:"slug"`
	Title			string		`json:"title"`
	Excerpt			string		`json:"excerpt"`
	Thumbnail		string		`json:"thumbnail"`
	CreatedAt		time.Time	`json:"created_at"`
	CategoryID		ReadCategoryResponse	`json:"category"`
	TagID			[]ReadTagResponse		`json:"tags"`
}