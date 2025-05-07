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

type BusinessWarehouseEntity struct {
	ID                  int    `sql:"id"`
	BusinessID          int    `sql:"business_id"`
	Name                string `sql:"name"`
	Description         string `sql:"description"`
	WarehousePictureURL string `sql:"warehouse_picture_url"`
	CreatedAt           string `sql:"created_at"`
	UpdatedAt           string `sql:"updated_at"`
}

type BusinessSupplierEntity struct {
	ID          int    `sql:"id"`
	BusinessID  int    `sql:"business_id"`
	Name        string `sql:"name"`
	Type        string `sql:"type"`
	Description string `sql:"description"`
	CreatedAt   string `sql:"created_at"`
	UpdatedAt   string `sql:"updated_at"`
}

type BusinessSupplierContactEntity struct {
	ID         int    `sql:"id"`
	SupplierID int    `sql:"supplier_id"`
	FullName   string `sql:"full_name"`
	Email      string `sql:"email"`
	PhoneNo    string `sql:"phone_no"`
	Address    string `sql:"address"`
	Remarks    string `sql:"remarks"`
	Status     string `sql:"status"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}

type BusinessProductEntity struct {
	ID         int    `sql:"id"`
	BusinessID int    `sql:"business_id"`
	SupplierID int    `sql:"supplier_id"`
	ItemName   string `sql:"item_name"`
	Brand      string `sql:"brand"`
	CategoryID int    `sql:"category_id"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}

type BusinessProductVariantEntity struct {
	ID                int     `sql:"id"`
	ProductID         int     `sql:"product_id"`
	VariantName       string  `sql:"variant_name"`
	SKUNo             string  `sql:"sku_no"`
	PictureURL        string  `sql:"picture_url"`
	BaseSellingPrice  float64 `sql:"base_selling_price"`
	BasePurchasePrice float64 `sql:"base_purchase_price"`
	Note              string  `sql:"note"`
	CreatedAt         string  `sql:"created_at"`
	UpdatedAt         string  `sql:"updated_at"`
}

type BusinessProductVariantTagEntity struct {
	VariantID int `sql:"variant_id"`
	TagID     int `sql:"tag_id"`
}
