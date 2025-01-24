package dto

// CreateNewBusinessReq represents the request body for creating a new business
type CreateNewBusinessReq struct {
	Name             string         `json:"name" validate:"required"`
	IndustryType     string         `json:"industryType" validate:"required"`
	BusinessType     string         `json:"businessType" validate:"required"`
	Description      string         `json:"description"`
	PhoneNo          string         `json:"phoneNo" validate:"required"`
	OperatingHours   OperatingHours `json:"operatingHours"`
	Address          string         `json:"address"`
	BusinessImageURL string         `json:"businessImageUrl"`
	OwnerUsername    string         `json:"ownerUsername" validate:"required"`
}

// OperatingHours represents the operating hours of a business for the week
type OperatingHours struct {
	Monday    OpenTime `json:"monday" validate:"required"`
	Tuesday   OpenTime `json:"tuesday" validate:"required"`
	Wednesday OpenTime `json:"wednesday" validate:"required"`
	Thursday  OpenTime `json:"thursday" validate:"required"`
	Friday    OpenTime `json:"friday" validate:"required"`
	Saturday  OpenTime `json:"saturday" validate:"required"`
	Sunday    OpenTime `json:"sunday" validate:"required"`
}

// OpenTime represents the open and close time of a business
type OpenTime struct {
	Open      bool   `json:"open"`
	OpenTime  string `json:"openTime,omitempty"`
	CloseTime string `json:"closeTime,omitempty"`
}

// CreateNewBusinessRes represents the response body for creating a new business
type CreateNewBusinessRes struct {
	BusinessID   int    `json:"businessId"`
	BusinessName string `json:"businessName"`
	Success      bool   `json:"success"`
}
