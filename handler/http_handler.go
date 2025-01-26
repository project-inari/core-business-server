package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/core-business-server/dto"
	"github.com/project-inari/core-business-server/pkg/request"
	"github.com/project-inari/core-business-server/pkg/response"
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

	businessName := c.Param("businessName")
	if businessName == "" {
		return response.ErrorResponse(c, http.StatusBadRequest, "error - [BusinessInquiry] bad request: business name is required", "")
	}

	res, err := h.d.Service.BusinessInquiry(ctx, businessName)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [BusinessInquiry] internal server error: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
