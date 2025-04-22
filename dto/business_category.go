package dto

type CreateNewCategoryReq struct {
	BusinessID         int    `json:"businessId" validate:"required"`
	CategoryName       string `json:"categoryName" validate:"required"`
	CategoryPictureURL string `json:"categoryPictureUrl"`
	Description        string `json:"description"`
	ParentCategoryID   int    `json:"parentCategoryId"`
	TagIDs             []int  `json:"tagIds" validate:"required"`
}

type CreateNewCategoryRes struct {
	CategoryID int  `json:"categoryId"`
	Success    bool `json:"success"`
}

type ListBusinessCategoriesRes struct {
	BusinessCategories []BusinessCategoryModel `json:"businessCategories"`
}

type BusinessCategoryModel struct {
	ID                 int                `json:"id"`
	BusinessID         *int               `json:"businessId,omitempty"`
	CategoryName       string             `json:"categoryName"`
	CategoryPictureURL string             `json:"categoryPictureUrl"`
	Description        string             `json:"description"`
	ParentCategoryID   *int               `json:"parentCategoryId"`
	Tags               []BusinessTagModel `json:"tags"`
	CreatedAt          string             `json:"createdAt"`
	UpdatedAt          string             `json:"updatedAt"`
}
