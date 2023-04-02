package handler

import (
	"app/api/models"
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

func (h *Handler) GetStockDataExcelDynamic(c *gin.Context) {
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

	// --------------------------------------------------------------------------------------------------------

	hash := map[int][]interface{}{}
	// get
	for i, stock := range stocks.Stocks {

		resp, err := h.storages.Stock().GetByID(context.Background(), &models.StockPrimaryKey{StoreId: stock.StoreId})
		if err != nil {
			h.handlerResponse(c, "storage.stock.getByID", http.StatusInternalServerError, err.Error())
			return
		}

		for _, product := range resp.Products {
			products := []interface{}{}

			if i == 0 {
				products = append(products, product.CategoryId)

				products = append(products, product.ProductName)
				products = append(products, product.ListPrice)
				products = append(products, product.Quantity)

				hash[product.ProductId] = products
			} else {
				hash[product.ProductId] = append(hash[product.ProductId], product.Quantity)
			}

		}

	}

	categoriesStocks, err := h.storages.Report().GetOnlyCategoryDataFromStock(context.Background())
	if err != nil {
		h.handlerResponse(c, "storage.stock.CategoryStock", http.StatusInternalServerError, err.Error())
		return
	}

	for _, cat := range categoriesStocks {
		obj := []interface{}{}
		obj = append(obj, cat.CategoryName, nil, cat.Quantity)

		data = append(data, obj)

		for _, val := range hash {
			if cat.CategoryId == val[0].(int) {
				obj := []interface{}{}
				obj = append(obj, val[1:]...)
				data = append(data, obj)
			}

		}

	}

	for idx, row := range data {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow("Sheet1", cell, &row)
	}

	if err := f.SaveAs("excel-data/stock_report.xlsx"); err != nil {
		fmt.Println(err)
	}

	h.handlerResponse(c, "get stock data", http.StatusCreated, "File saved successfully!")
}
