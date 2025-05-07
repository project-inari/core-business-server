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
	SupplierID int    `json:"supplierId"`
	FullName   string `json:"fullName"`
	Email      string `json:"email"`
	PhoneNo    string `json:"phoneNo"`
	Address    string `json:"address"`
	Remarks    string `json:"remarks"`
	Status     string `json:"status"`
	CreatedAt  string `json:"createdAt"`
	UpdatedAt  string `json:"updatedAt"`
}

type CreateNewSupplierOrderReq struct {
	BusinessID         int                 `json:"businessId" validate:"required"`
	SupplierID         int                 `json:"supplierId" validate:"required"`
	ReceiveID          string              `json:"receiveId" validate:"required"`
	WarehouseID        int                 `json:"warehouseId" validate:"required"`
	Status             string              `json:"status" validate:"required"`
	ShippingMethod     string              `json:"shippingMethod" validate:"required"`
	ShippingCost       float64             `json:"shippingCost" validate:"required"`
	SupplierOrderItems []SupplierOrderItem `json:"supplierOrderItems" validate:"required"`
}

type CreateNewSupplierOrderRes struct {
	SupplierOrderID int  `json:"supplierOrderId"`
	Success         bool `json:"success"`
}

type SupplierOrderItem struct {
	VariantID int `json:"variantId" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type ListSupplierOrdersRes struct {
	SupplierOrders []SupplierOrderModel `json:"supplierOrders"`
}

type SupplierOrderModel struct {
	ID             int                      `json:"id"`
	ReceiveID      string                   `json:"receiveId"`
	SupplierID     int                      `json:"supplierId"`
	WarehouseID    int                      `json:"warehouseId"`
	Status         string                   `json:"status"`
	ShippingMethod string                   `json:"shippingMethod"`
	ShippingCost   float64                  `json:"shippingCost"`
	OrderItems     []SupplierOrderItemModel `json:"orderItems"`
	CreatedAt      string                   `json:"createdAt"`
	UpdatedAt      string                   `json:"updatedAt"`
}

type SupplierOrderItemModel struct {
	VariantID         int     `json:"variantId"`
	ProductName       string  `json:"productName"`
	VariantName       string  `json:"variantName"`
	SKUNo             string  `json:"skuNo"`
	BasePurchasePrice float64 `json:"basePurchasePrice"`
	PictureURL        string  `json:"pictureUrl"`
	CategoryID        int     `json:"categoryId"`
	Quantity          int     `json:"quantity"`
}
