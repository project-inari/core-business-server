package dto

// BusinessEntity represents the business entity in the database for table tbl_businesses
type BusinessEntity struct {
	ID               int    `sql:"id"`
	Name             string `sql:"name"`
	IndustryType     string `sql:"industry_type"`
	BusinessType     string `sql:"business_type"`
	Description      string `sql:"description"`
	PhoneNo          string `sql:"phone_no"`
	OperatingHours   string `sql:"operating_hours"`
	Address          string `sql:"address"`
	BusinessImageURL string `sql:"business_image_url"`
	CreatedAt        string `sql:"created_at"`
	UpdatedAt        string `sql:"updated_at"`
}

// BusinessMemberEntity represents the business member entity in the database for table tbl_business_members
type BusinessMemberEntity struct {
	ID         int    `sql:"id"`
	BusinessID int    `sql:"business_id"`
	Username   string `sql:"username"`
	Role       string `sql:"role"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}

// BusinessJoiningEntity represents the business joining entity in the database for table tbl_business_joining
type BusinessJoiningEntity struct {
	ID         int    `sql:"id"`
	BusinessID int    `sql:"business_id"`
	Username   string `sql:"username"`
	Status     string `sql:"status"`
	ActionedBy string `sql:"actioned_by"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}

// BusinessCategoryEntity represents the business category entity in the database for table tbl_business_categories
type BusinessCategoryEntity struct {
	ID                 int    `sql:"id"`
	BusinessID         int    `sql:"business_id"`
	CategoryName       string `sql:"category_name"`
	CategoryPictureURL string `sql:"category_picture_url"`
	Description        string `sql:"description"`
	ParentCategoryID   *int   `sql:"parent_category_id"`
	CreatedAt          string `sql:"created_at"`
	UpdatedAt          string `sql:"updated_at"`
}

type BusinessTagEntity struct {
	ID         int    `sql:"id"`
	BusinessID int    `sql:"business_id"`
	TagName    string `sql:"tag_name"`
	Color      string `sql:"color"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}
