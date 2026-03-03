package model

import "time"

// 도메인 엔티티 (API <-> DB)
type Category struct {
	ID				int32		`json:"id" db:"id"`
	ParentID		int32		`json:"parent_id" db:"parent_id"`
	Path			string		`json:"path" db:"path"`
	Name			string		`json:"name" db:"name"`
	Slug			string		`json:"slug" db:"slug"`
	CreatedAt		time.Time	`json:"created_at" db:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at" db:"updated_at"`
}

// 요청 DTO (Client -> API)

type CreateCategoryRequest struct {
	ParentID		int32		`json:"parent_id"`
	Name			string		`json:"name" binding:"required,max=100"`
	Slug			string		`json:"slug" binding:"required,max=150"`
}

type ReadCategoryRequest struct {
	Search			string		`json:"search"`
	SortBy			string		`json:"sort_by"`
	SortDir			string		`json:"sort_dir"`
}

type UpdateCategoryRequest struct {
	ParentID		int32		`json:"parent_id"`
	Name			string		`json:"name" binding:"max=100"`
	Slug			string		`json:"slug" binding:"max=150"`
}


// 응답 DTO (Client <- API)

type CreateCategoryResponse struct {
	ID				int32		`json:"id"`
	ParentID		int32		`json:"parent_id"`
	Name			string		`json:"name"`
	Slug			string		`json:"slug"`
	Path			string		`json:"path"`
	Breadcrumb		string		`json:"breadcrumb"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

type ReadCategoriesResponse struct {
	Result struct {
		Meta	struct{
			Sort struct {
				SortBy		string	`json:"by"`
				SortDir		string	`json:"dir"`
			}	`json:"sort"`
			Total 		int32		`json:"total"`
			Filters struct {
				Search 	string		`json:"search"`
			}
		}		`json:"meta"`
		Data	[]CategoriesObject		`json:"data"`
	} `json:"result"`
}

type ReadCategoryResponse struct {
	ID				int32		`json:"id"`
	ParentID		int32		`json:"parent_id"`
	Name			string		`json:"name"`
	Slug			string		`json:"slug"`
	Path			string		`json:"path"`
	Breadcrumb		string		`json:"breadcrumb"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time	`json:"updated_at"`
}

type UpdateCategoryResponse struct {
	Result struct {
		Data		CreateCategoryResponse	`json:"data"`
		Diff		CommonUpdateDiff		`json:"diff"`
	} `json:"result"`
}

type CategoriesObject struct {
	ID				int32		`json:"id"`
	Name			string		`json:"name"`
	Slug			string		`json:"slug"`
	Breadcrumb		string		`json:"breadcrumb"`
	Depth			int32		`json:"depth"`
}