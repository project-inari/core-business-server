package dto

type CreateNewTagReq struct {
	BusinessID  int    `json:"businessId" validate:"required"`
	TagName     string `json:"tagName" validate:"required"`
	Color       string `json:"color" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type CreateNewTagRes struct {
	TagID   int  `json:"tagId"`
	Success bool `json:"success"`
}

type ListBusinessTagsRes struct {
	BusinessTags []BusinessTagModel `json:"businessTags"`
}

type BusinessTagModel struct {
	ID          int    `json:"id"`
	BusinessID  *int   `json:"businessId,omitempty"`
	TagName     string `json:"tagName"`
	Color       string `json:"color"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}
