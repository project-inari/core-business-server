package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *httpHandler) initRoutes(e *echo.Echo) {
	e.GET("/health", h.HealthCheck)

	v1 := e.Group("/v1")
	v1.POST("/create", h.CreateNewBusiness)
	v1.GET("/inquiry/:businessID", h.BusinessInquiry)

	v1Category := v1.Group("/category")
	v1Category.POST("/create", h.CreateNewCategory)
	v1Category.GET("/list/:businessID", h.ListBusinessCategories)

	v1Tag := v1.Group("/tag")
	v1Tag.POST("/create", h.CreateNewTag)
	v1Tag.GET("/list/:businessID", h.ListBusinessTags)

	v1Warehouse := v1.Group("/warehouse")
	v1Warehouse.POST("/create", h.CreateNewWarehouse)
	v1Warehouse.GET("/list/:businessID", h.ListBusinessWarehouses)

	v1Supplier := v1.Group("/supplier")
	v1Supplier.POST("/create", h.CreateNewSupplier)
	v1Supplier.POST("/contact/create", h.CreateNewSupplierContact)
	v1Supplier.GET("/list/:businessID", h.ListBusinessSuppliers)
	v1Supplier.GET("/contact/inquiry/:supplierID", h.InquiryBusinessSupplier)
	v1Supplier.POST("/order/create", h.CreateNewSupplierOrder)
	v1Supplier.GET("/order/list/:businessID", h.ListSupplierOrders)

	v1Product := v1.Group("/product")
	v1Product.POST("/create", h.CreateNewProduct)
	v1Product.GET("/list/:businessID", h.ListBusinessProducts)
	v1Product.GET("/inquiry/:productID", h.InquiryBusinessProduct)

	v1Inventory := v1.Group("/inventory")
	v1Inventory.GET("/list/:businessID", h.ListBusinessInventory)
	v1Inventory.GET("/inquiry/:variantID", h.InquiryProductInventory)
}
