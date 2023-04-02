package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/xuri/excelize/v2"
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

// Get Employee Report
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

// TASK4

// Get TotalOrderPrice
// @ID get_order_price
// @Router /total_order_price/{id} [GET]
// @Summary Report
// @Description Report
// @Tags Report
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param promo_code query string false "promo_code"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) TotalOrderPrice(c *gin.Context) {

	id := c.Param("id")
	promo := c.Query("promo_code")

	fmt.Println(id)
	orderIdInt, err := strconv.Atoi(id)
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.getByID", http.StatusBadRequest, "id incorrect")
		return
	}

	resp, err := h.storages.Order().GetByID(context.Background(), &models.OrderPrimaryKey{OrderId: orderIdInt})
	if err != nil {
		h.handlerResponse(c, "storage.order.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	var response models.TotalOrderPrice
	// calc total price
	var totalPrice float64
	for _, item := range resp.OrderItems {
		totalPrice += item.ListPrice
	}

	response.TotalPrice = totalPrice

	fmt.Println(promo)
	if len(promo) > 0 {
		// get promocode
		promoCode, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{Name: promo})
		if err != nil {
			h.handlerResponse(c, "storage.promo_code.getByID", http.StatusInternalServerError, err.Error())
			return
		}

		response.PromoCode = promoCode.Name

		if promoCode.DiscountType == "fixed" && totalPrice >= promoCode.OrderLimitPrice {
			totalPrice = totalPrice - promoCode.Discount
			response.Discount = promoCode.Discount
		} else if promoCode.DiscountType == "percent" && totalPrice >= promoCode.OrderLimitPrice {
			totalPrice = totalPrice - totalPrice*promoCode.Discount/100
			response.Discount = math.Round((totalPrice*promoCode.Discount/100)*100) / 100
		}

		if totalPrice <= 0 {
			totalPrice = 0
		}
	}

	response.ResultPrice = math.Round(totalPrice*100) / 100
	h.handlerResponse(c, "get list product response", http.StatusOK, response)
}

// TASK6

// Get GetStockDataExcel
// @ID get_report_stock
// @Router /report/stock_excel [GET]
// @Summary GetStockDataExcel
// @Description GetStockDataExcel
// @Tags Report
// @Accept json
// @Produce json
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetStockDataExcel(c *gin.Context) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	stocks, err := h.storages.Stock().GetList(context.Background(), &models.GetListStockRequest{
		Offset: 0,
		Limit:  10,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.stock.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	data := [][]interface{}{}
	// dynamic titles
	stores := []interface{}{
		"Намеклатура", "Цена",
	}
	// get stores
	for _, stock := range stocks.Stocks {
		store, err := h.storages.Store().GetByID(context.Background(), &models.StorePrimaryKey{StoreId: stock.StoreId})
		if err != nil {
			h.handlerResponse(c, "storage.store.getByID", http.StatusInternalServerError, err.Error())
			return
		}
		stores = append(stores, store.StoreName)
	}
	data = append(data, stores)

	for idx, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow("Sheet1", cell, &row)
	}

	// get category products Data
	for index, stock := range stocks.Stocks {

		resp, err := h.storages.Report().GetCategoryData(context.Background(), stock.StoreId)
		if err != nil {
			h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
			return
		}

		k := 2
		for _, categoryProduct := range resp {

			location := fmt.Sprintf("A%d", k)
			location2 := fmt.Sprintf("E%d", k)

			style, err := f.NewStyle(&excelize.Style{
				Fill:      excelize.Fill{Type: "pattern", Color: []string{"#808080"}, Pattern: 1},
				Alignment: &excelize.Alignment{Horizontal: "center"},
			})
			if err != nil {
				fmt.Println(err)
			}

			f.SetCellValue("Sheet1", location, categoryProduct.CategoryName)
			f.SetColWidth("Sheet1", "A", "A", 50)
			f.SetColWidth("Sheet1", "B", "B", 16)
			f.SetColWidth("Sheet1", "C", "C", 16)
			f.SetColWidth("Sheet1", "D", "D", 16)
			f.SetColWidth("Sheet1", "E", "E", 16)

			locationAmountCategory := ""
			if index == 0 {
				locationAmountCategory = fmt.Sprintf("C%d", k)
			} else if index == 1 {
				locationAmountCategory = fmt.Sprintf("D%d", k)
			} else if index == 2 {
				locationAmountCategory = fmt.Sprintf("E%d", k)
			}

			f.SetCellValue("Sheet1", locationAmountCategory, categoryProduct.Quantity)

			f.SetCellStyle("Sheet1", location, location2, style)

			for j, p := range categoryProduct.CategoryShopProducts {

				location3 := fmt.Sprintf("A%d", k)
				locationPrice := fmt.Sprintf("B%d", k)

				if j == 0 {
					location3 = fmt.Sprintf("A%d", k+1)
					locationPrice = fmt.Sprintf("B%d", k+1)
					k++
				}

				locationAmount := ""
				if index == 0 {
					locationAmount = fmt.Sprintf("C%d", k)
				} else if index == 1 {
					locationAmount = fmt.Sprintf("D%d", k)
				} else if index == 2 {
					locationAmount = fmt.Sprintf("E%d", k)
				}

				f.SetCellValue("Sheet1", location3, p.ProductName)
				f.SetCellValue("Sheet1", locationPrice, p.ListPrice)
				f.SetCellValue("Sheet1", locationAmount, p.Quantity)
				k++
			}

		}

	}

	if err := f.SaveAs("excel-data/stock_report.xlsx"); err != nil {
		fmt.Println(err)
	}

	h.handlerResponse(c, "get stock data", http.StatusCreated, "File saved successfully! (path: ./excel-data/stock_report.xlsx)")
}
