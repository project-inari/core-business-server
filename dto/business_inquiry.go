package dto

// BusinessInquiryRes represents the response of business inquiry
type BusinessInquiryRes struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	IndustryType     string         `json:"industryType"`
	BusinessType     string         `json:"businessType"`
	Description      string         `json:"description"`
	PhoneNo          string         `json:"phoneNo"`
	OperatingHours   OperatingHours `json:"operatingHours"`
	Address          string         `json:"address"`
	BusinessImageURL string         `json:"businessImageUrl"`
	CreatedAt        string         `json:"createdAt"`
	UpdatedAt        string         `json:"updatedAt"`
}
