package dto

type CreateNewCustomerReq struct {
	BusinessID int    `json:"businessId" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Type       string `json:"type" validate:"required"`
	Address    string `json:"address" validate:"required"`
	PhoneNo    string `json:"phoneNo" validate:"required"`
}

type CreateNewCustomerRes struct {
	CustomerID int  `json:"customerId"`
	Success    bool `json:"success"`
}

type ListBusinessCustomersRes struct {
	Customers []BusinessCustomerModel `json:"customers"`
}

type BusinessCustomerModel struct {
	ID         int    `json:"id"`
	BusinessID *int   `json:"businessId,omitempty"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	Address    string `json:"address"`
	PhoneNo    string `json:"phoneNo"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}
