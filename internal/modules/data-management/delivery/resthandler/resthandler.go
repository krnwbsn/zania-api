// Code generated by candi v1.18.0.

package resthandler

import (
	"net/http"

	"zania-api/internal/modules/data-management/domain"
	"zania-api/pkg/shared/usecase"

	"github.com/golangid/candi/candihelper"
	"github.com/golangid/candi/codebase/factory/dependency"
	"github.com/golangid/candi/codebase/interfaces"
	"github.com/golangid/candi/tracer"
	"github.com/golangid/candi/wrapper"
)

// RestHandler handler
type RestHandler struct {
	mw        interfaces.Middleware
	uc        usecase.Usecase
	validator interfaces.Validator
}

// NewRestHandler create new rest handler
func NewRestHandler(uc usecase.Usecase, deps dependency.Dependency) *RestHandler {
	return &RestHandler{
		uc: uc, mw: deps.GetMiddleware(), validator: deps.GetValidator(),
	}
}

// Mount handler with root "/"
// handling version in here
func (h *RestHandler) Mount(root interfaces.RESTRouter) {
	v1DataManagement := root.Group(candihelper.V1 + "/data-management")

	v1DataManagement.GET("/", h.getAllDataManagement)
}

// GetAllDataManagement documentation
// @Summary			Get All DataManagement
// @Description		API for get all data-management
// @Tags			DataManagement
// @Accept			json
// @Produce			json
// @Param			page	query	string	false	"Page with default value is 1"
// @Param			limit	query	string	false	"Limit with default value is 10"
// @Param			search	query	string	false	"Search"
// @Param			orderBy	query	string	false	"Order By"
// @Param			sort	query	string	false	"Sort (ASC DESC)"
// @Success			200	{object}	domain.ResponseDataManagementList
// @Success			400	{object}	wrapper.HTTPResponse
// @Security		ApiKeyAuth
// @Router			/v1/data-management [get]
func (h *RestHandler) getAllDataManagement(rw http.ResponseWriter, req *http.Request) {
	trace, ctx := tracer.StartTraceWithContext(req.Context(), "DataManagementDeliveryREST:GetAllDataManagement")
	defer trace.Finish()

	var filter domain.FilterDataManagement
	if err := candihelper.ParseFromQueryParam(req.URL.Query(), &filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed parse filter", err).JSON(rw)
		return
	}

	if err := h.validator.ValidateDocument("data-management/get_all", filter); err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, "Failed validate filter", err).JSON(rw)
		return
	}

	result, err := h.uc.DataManagement().GetAllDataManagement(ctx, &filter)
	if err != nil {
		wrapper.NewHTTPResponse(http.StatusBadRequest, err.Error()).JSON(rw)
		return
	}

	message := "Success"
	response := wrapper.NewHTTPResponse(http.StatusOK, message, result.Data)
	response.Meta = result.Meta
	response.JSON(rw)
}
