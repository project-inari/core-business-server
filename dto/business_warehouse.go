package dto

type CreateNewWarehouseReq struct {
	BusinessID          int    `json:"businessId" validate:"required"`
	WarehouseName       string `json:"warehouseName" validate:"required"`
	Description         string `json:"description"`
	WarehousePictureURL string `json:"warehousePictureUrl"`
}

type CreateNewWarehouseRes struct {
	WarehouseID int  `json:"warehouseId"`
	Success     bool `json:"success"`
}

type ListBusinessWarehousesRes struct {
	Warehouses []BusinessWarehouseModel `json:"warehouses"`
}

type BusinessWarehouseModel struct {
	ID                  int    `json:"id"`
	BusinessID          *int   `json:"businessId,omitempty"`
	Name                string `json:"name"`
	Description         string `json:"description"`
	WarehousePictureURL string `json:"warehousePictureUrl"`
	CreatedAt           string `json:"createdAt"`
	UpdatedAt           string `json:"updatedAt"`
}
