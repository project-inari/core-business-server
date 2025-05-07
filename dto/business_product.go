package dto

type CreateNewProductReq struct {
	BusinessID int                    `json:"businessId" validate:"required"`
	SupplierID int                    `json:"supplierId" validate:"required"`
	Name       string                 `json:"name" validate:"required"`
	Brand      string                 `json:"brand" validate:"required"`
	CategoryID int                    `json:"categoryId" validate:"required"`
	Variants   []BusinessVariantModel `json:"variants"`
}

type CreateNewProductRes struct {
	ProductID int  `json:"productId"`
	Success   bool `json:"success"`
}

type ListBusinessProductsRes struct {
	Products []BusinessProductModel `json:"products"`
}

type ListBusinessProductVariantsRes struct {
	Variants []BusinessVariantModel `json:"variants"`
}

type BusinessProductModel struct {
	ID         int                    `json:"id"`
	BusinessID int                    `json:"businessId"`
	SupplierID int                    `json:"supplierId"`
	Name       string                 `json:"name"`
	Brand      string                 `json:"brand"`
	CategoryID int                    `json:"categoryId"`
	Variants   []BusinessVariantModel `json:"variants"`
	CreatedAt  string                 `json:"createdAt"`
	UpdatedAt  string                 `json:"updatedAt"`
}

type BusinessVariantModel struct {
	ID                int            `json:"variantId"`
	ProductID         int            `json:"productId"`
	VariantName       string         `json:"variantName"`
	SKUNo             string         `json:"skuNo"`
	PictureURL        string         `json:"pictureUrl"`
	BaseSellingPrice  float64        `json:"baseSellingPrice"`
	BasePurchasePrice float64        `json:"basePurchasePrice"`
	Note              string         `json:"note"`
	TagIds            []int          `json:"tagIds"`
	QtyInWarehouse    []WarehouseQty `json:"qtyInWarehouse"`
	CreatedAt         string         `json:"createdAt"`
	UpdatedAt         string         `json:"updatedAt"`
}

type WarehouseQty struct {
	WarehouseID int `json:"warehouseId" sql:"warehouse_id"`
	Qty         int `json:"qty" sql:"quantity"`
}
