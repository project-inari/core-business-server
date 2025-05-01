package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/pkg/request"
	"github.com/project-inari/core-business-server/pkg/response"
	"github.com/project-inari/core-business-server/pkg/utils"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

// CreateNewBusiness creates a new business
func (h *httpHandler) CreateNewBusiness(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewBusinessReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewBusiness] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewBusiness(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewBusiness] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// BusinessInquiry retrieves a business information
func (h *httpHandler) BusinessInquiry(c echo.Context) error {
	ctx := c.Request().Context()

	businessID := c.Param("businessID")
	if businessID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [BusinessInquiry] bad request: business name is required", "")
	}

	res, err := h.d.Service.BusinessInquiry(ctx, utils.ConvertStringToInt(businessID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [BusinessInquiry] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) CreateNewCategory(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewCategoryReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewCategory] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewCategory(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewCategory] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) ListBusinessCategories(c echo.Context) error {
	ctx := c.Request().Context()

	businessID := c.Param("businessID")
	if businessID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [ListBusinessCategories] bad request: business name is required", "")
	}

	res, err := h.d.Service.ListBusinessCategories(ctx, utils.ConvertStringToInt(businessID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [ListBusinessCategories] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) CreateNewTag(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewTagReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewTag] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewTag(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewTag] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) ListBusinessTags(c echo.Context) error {
	ctx := c.Request().Context()

	businessID := c.Param("businessID")
	if businessID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [ListBusinessTags] bad request: business name is required", "")
	}

	res, err := h.d.Service.ListBusinessTags(ctx, utils.ConvertStringToInt(businessID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [ListBusinessTags] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) CreateNewWarehouse(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewWarehouseReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewWarehouse] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewWarehouse(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewWarehouse] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) ListBusinessWarehouses(c echo.Context) error {
	ctx := c.Request().Context()

	businessID := c.Param("businessID")
	if businessID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [ListBusinessWarehouses] bad request: business name is required", "")
	}

	res, err := h.d.Service.ListBusinessWarehouses(ctx, utils.ConvertStringToInt(businessID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [ListBusinessWarehouses] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) CreateNewSupplier(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewSupplierReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewSupplier] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewSupplier(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewSupplier] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) CreateNewSupplierContact(c echo.Context) error {
	ctx := c.Request().Context()
	wrapper := request.ContextWrapper(c)

	req := new(dto.CreateNewSupplierContactReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [CreateNewSupplierContact] bad request: %v", err), "")
	}

	res, err := h.d.Service.CreateNewSupplierContact(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [CreateNewSupplierContact] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) ListBusinessSuppliers(c echo.Context) error {
	ctx := c.Request().Context()

	businessID := c.Param("businessID")
	if businessID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [ListBusinessSuppliers] bad request: business name is required", "")
	}

	res, err := h.d.Service.ListBusinessSuppliers(ctx, utils.ConvertStringToInt(businessID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [ListBusinessSuppliers] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

func (h *httpHandler) InquiryBusinessSupplier(c echo.Context) error {
	ctx := c.Request().Context()

	supplierID := c.Param("supplierID")
	if supplierID == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [InquiryBusinessSupplier] bad request: supplier ID is required", "")
	}

	res, err := h.d.Service.InquiryBusinessSupplier(ctx, utils.ConvertStringToInt(supplierID))
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [InquiryBusinessSupplier] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
