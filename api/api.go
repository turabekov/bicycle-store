package api

import (
	_ "app/api/docs"

	"app/api/handler"
	"app/config"
	"app/pkg/logger"
	"app/storage"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func NewApi(r *gin.Engine, cfg *config.Config, store storage.StorageI, logger logger.LoggerI) {
	handler := handler.NewHandler(cfg, store, logger)

	// @securityDefinitions.apikey ApiKeyAuth
	// @in header
	// @name Authorization

	r.Use(customCORSMiddleware())

	r.POST("/register", handler.Register)
	r.POST("/login", handler.Login)
	// user api
	r.POST("/user", handler.CreateUser)
	r.GET("/user/:id", handler.GetByIdUser)
	r.GET("/user", handler.GetListUser)
	r.PUT("/user/:id", handler.UpdateUser)
	r.DELETE("/user/:id", handler.DeleteUser)

	// promo_code api
	r.POST("/promo_code", handler.CreatePromoCode)
	r.GET("/promo_code/:id", handler.GetByIdPromoCode)
	r.GET("/promo_code", handler.GetListPromoCode)
	r.DELETE("/promo_code/:id", handler.DeletePromoCode)

	// category api
	r.POST("/category", handler.CreateCategory)
	r.GET("/category/:id", handler.GetByIdCategory)
	r.GET("/category", handler.GetListCategory)
	r.PUT("/category/:id", handler.UpdateCategory)
	r.DELETE("/category/:id", handler.DeleteCategory)

	// brand api
	r.POST("/brand", handler.CreateBrand)
	r.GET("/brand/:id", handler.GetByIdBrand)
	r.GET("/brand", handler.GetListBrand)
	r.PUT("/brand/:id", handler.UpdateBrand)
	r.DELETE("/brand/:id", handler.DeleteBrand)

	// product api
	r.POST("/product", handler.CreateProduct)
	r.GET("/product/:id", handler.GetByIdProduct)
	r.GET("/product", handler.GetListProduct)
	r.PUT("/product/:id", handler.UpdateProduct)
	r.DELETE("/product/:id", handler.DeleteProduct)

	// stock api  -- not ready for using
	r.POST("/stock", handler.CreateStock)
	r.GET("/stock/:id", handler.GetByIdStock)
	r.GET("/stock", handler.GetListStock)
	r.PUT("/stock/:id", handler.UpdateStock)
	r.DELETE("/stock/:id", handler.DeleteStock)

	// store api
	r.POST("/store", handler.CreateStore)
	r.GET("/store/:id", handler.GetByIdStore)
	r.GET("/store", handler.GetListStore)
	r.PUT("/store/:id", handler.UpdateStore)
	r.PATCH("/store/:id", handler.UpdatePatchStore)
	r.DELETE("/store/:id", handler.DeleteStore)

	// customer api
	r.POST("/customer", handler.CreateCustomer)
	r.GET("/customer/:id", handler.GetByIdCustomer)
	r.GET("/customer", handler.GetListCustomer)
	r.PUT("/customer/:id", handler.UpdateCustomer)
	r.PATCH("/customer/:id", handler.UpdatePatchCustomer)
	r.DELETE("/customer/:id", handler.DeleteCustomer)

	// staff api
	r.POST("/staff", handler.CreateStaff)
	r.GET("/staff/:id", handler.GetByIdStaff)
	r.GET("/staff", handler.GetListStaff)
	r.PUT("/staff/:id", handler.UpdateStaff)
	r.PATCH("/staff/:id", handler.UpdatePatchStaff)
	r.DELETE("/staff/:id", handler.DeleteStaff)

	// order api
	r.POST("/order", handler.AuthMiddleware(), handler.CreateOrder)
	r.GET("/order/:id", handler.AuthMiddleware(), handler.GetByIdOrder)
	r.GET("/order", handler.AuthMiddleware(), handler.GetListOrder)
	r.PUT("/order/:id", handler.AuthMiddleware(), handler.UpdateOrder)
	r.PATCH("/order/:id", handler.AuthMiddleware(), handler.UpdatePatchOrder)
	r.DELETE("/order/:id", handler.AuthMiddleware(), handler.DeleteOrder)
	r.POST("/order_item/", handler.AuthMiddleware(), handler.CreateOrderItem)
	r.DELETE("/order_item/:id", handler.AuthMiddleware(), handler.DeleteOrderItem)

	//  report
	r.PUT("/report/exchange", handler.ExchangeStoreProductHandler)
	r.GET("/report/employee", handler.GetEmployeeReport)
	r.GET("/total_order_price/:id", handler.TotalOrderPrice)
	//  excel
	r.GET("/report/stock_excel", handler.GetStockDataExcel)

	// r.GET("/report/stock", handler.GetStockDataExcelDynamic)

	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}

func customCORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Acces-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH, DELETE, HEAD")
		c.Header("Access-Control-Allow-Headers", "Platform-Id, Content-Type, Content-Length, Accept-Encoding, X-CSF-TOKEN, Authorization, Cache-Control")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
