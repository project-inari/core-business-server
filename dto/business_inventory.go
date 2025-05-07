package dto

type ListBusinessInventoryRes struct {
	Inventory []BusinessInventoryModel `json:"inventory"`
}

type BusinessInventoryModel struct {
	VariantID         int                     `json:"variantId"`
	SKUNo             string                  `json:"skuNo"`
	ProductName       string                  `json:"productName"`
	VariantName       string                  `json:"variantName"`
	PictureURL        string                  `json:"pictureUrl"`
	BaseSellingPrice  float64                 `json:"baseSellingPrice"`
	BasePurchasePrice float64                 `json:"basePurchasePrice"`
	Note              string                  `json:"note"`
	SupplierID        int                     `json:"supplierId"`
	CategoryID        int                     `json:"categoryId"`
	Categories        []InventoryCategoryInfo `json:"categories"`
	Tags              []BusinessTagModel      `json:"tags"`
	QtyInWarehouse    []WarehouseQty          `json:"qtyInWarehouse"`
}

type InventoryCategoryInfo struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Parent []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"parent"`
}
