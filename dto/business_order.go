package dto

type CreateNewCustomerOrderReq struct {
	BusinessID         int                 `json:"businessId" validate:"required"`
	OrderID            string              `json:"orderId" validate:"required"`
	CustomerID         int                 `json:"customerId" validate:"required"`
	ChannelID          int                 `json:"channelId" validate:"required"`
	StatusID           int                 `json:"statusId" validate:"required"`
	ShippingMethod     string              `json:"shippingMethod" validate:"required"`
	ShippingFee        float64             `json:"shippingFee" validate:"required"`
	ShippingCost       float64             `json:"shippingCost" validate:"required"`
	CustomerOrderItems []CustomerOrderItem `json:"customerOrderItems" validate:"required"`
}

type CreateNewCustomerOrderRes struct {
	CustomerOrderID int  `json:"customerOrderId"`
	Success         bool `json:"success"`
}

type CustomerOrderItem struct {
	VariantID       int     `json:"variantId" validate:"required"`
	WarehouseID     int     `json:"warehouseId" validate:"required"`
	Quantity        int     `json:"quantity" validate:"required"`
	PricePerUnit    float64 `json:"pricePerUnit" validate:"required"`
	DiscountPerUnit float64 `json:"discountPerUnit" validate:"required"`
}

type ListCustomerOrdersRes struct {
	CustomerOrders []CustomerOrderModel `json:"customerOrders"`
}

type CustomerOrderModel struct {
	ID             int                 `json:"id"`
	OrderID        string              `json:"orderId"`
	CustomerID     int                 `json:"customerId"`
	ChannelID      int                 `json:"channelId"`
	StatusID       int                 `json:"statusId"`
	ShippingMethod string              `json:"shippingMethod"`
	ShippingFee    float64             `json:"shippingFee"`
	ShippingCost   float64             `json:"shippingCost"`
	Items          []CustomerOrderItem `json:"items"`
	CreatedAt      string              `json:"createdAt"`
	UpdatedAt      string              `json:"updatedAt"`
}
