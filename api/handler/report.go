package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// TASK1

// Exchange Products In Stores godoc
// @ID create_exchange
// @Router /report/exchange [PUT]
// @Summary Exchange Products In Stores
// @Description Exchange Products In Stores
// @Tags Report
// @Accept json
// @Produce json
// @Param exchange body models.ExchangeProduct true "ExchangeProductRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) ExchangeStoreProductHandler(c *gin.Context) {
	var exchange models.ExchangeProduct

	err := c.ShouldBindJSON(&exchange) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "exchange  product", http.StatusBadRequest, err.Error())
		return
	}

	// check amount of products in store with given amount
	resp, err := h.storages.Stock().GetByIdProductStock(context.Background(), exchange.FromStoreId, exchange.ProductId)
	if err != nil {
		h.handlerResponse(c, "storage.exchange_product.getByID", http.StatusInternalServerError, err.Error())
		return
	}
	if resp.Quantity == 0 || resp.Quantity < exchange.Quantity {
		h.handlerResponse(c, "storage.exchange_product", http.StatusBadRequest, "not enough products for exchange")
		return
	}

	fromId, toId, err := h.storages.Report().ExchangeStoreProduct(context.Background(), &exchange)
	if err != nil {
		h.handlerResponse(c, "storage.exchange_product.update", http.StatusInternalServerError, err.Error())
		return
	}

	// get stock datas after exchange
	fromData, err := h.storages.Stock().GetByIdProductStock(context.Background(), fromId, exchange.ProductId)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}
	toData, err := h.storages.Stock().GetByIdProductStock(context.Background(), toId, exchange.ProductId)
	if err != nil {
		h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "exchange_product", http.StatusCreated, models.ResponseExchange{FromData: fromData, ToData: toData})

}

// TASK2

// Get LEmployee Report
// @ID get_list_emmployee_report
// @Router /report/employee [GET]
// @Summary Report
// @Description Report
// @Tags Report
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetEmployeeReport(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list product", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.Report().StaffSaleReport(context.Background(), &models.GetListEmployeeReportRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.product.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list product response", http.StatusOK, resp)
}
