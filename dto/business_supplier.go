package dto

type CreateNewSupplierReq struct {
	BusinessID   int    `json:"businessId" validate:"required"`
	SupplierName string `json:"supplierName" validate:"required"`
	Description  string `json:"description"`
	Type         string `json:"type"`
}

type CreateNewSupplierRes struct {
	SupplierID int  `json:"supplierId"`
	Success    bool `json:"success"`
}

type ListBusinessSuppliersRes struct {
	Suppliers []BusinessSupplierModel `json:"suppliers"`
}

type BusinessSupplierModel struct {
	ID               int                             `json:"id"`
	BusinessID       *int                            `json:"businessId,omitempty"`
	Name             string                          `json:"name"`
	Type             string                          `json:"type"`
	Description      string                          `json:"description"`
	SupplierContacts *[]BusinessSupplierContactModel `json:"supplierContacts"`
	CreatedAt        string                          `json:"createdAt"`
	UpdatedAt        string                          `json:"updatedAt"`
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

type CreateNewSupplierContactReq struct {
	BusinessID int    `json:"businessId" validate:"required"`
	SupplierID int    `json:"supplierId" validate:"required"`
	FullName   string `json:"fullName" validate:"required"`
	Email      string `json:"email" validate:"required,email"`
	PhoneNo    string `json:"phoneNo" validate:"required"`
	Address    string `json:"address" validate:"required"`
	Remarks    string `json:"remarks"`
	Status     string `json:"status" validate:"required"`
}

type CreateNewSupplierContactRes struct {
	SupplierContactID int  `json:"supplierContactId"`
	Success           bool `json:"success"`
}

type BusinessSupplierContactModel struct {
	ID         int    `json:"id"`
	SupplierID *int   `json:"supplierId,omitempty"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	PhoneNo    string `json:"phoneNo"`
	Address    string `json:"address"`
	Remarks    string `json:"remarks"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type BusinessSupplierContactEntity struct {
	ID         int    `sql:"id"`
	SupplierID *int   `sql:"supplier_id"`
	FullName   string `sql:"full_name"`
	Email      string `sql:"email"`
	PhoneNo    string `sql:"phone_no"`
	Address    string `sql:"address"`
	Remarks    string `sql:"remarks"`
	Status     string `sql:"status"`
	CreatedAt  string `sql:"created_at"`
	UpdatedAt  string `sql:"updated_at"`
}
