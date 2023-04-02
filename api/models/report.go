package models

// task1

type ExchangeProduct struct {
	FromStoreId int `json:"from_store_id"`
	ToStoreId   int `json:"to_store_id"`
	ProductId   int `json:"product_id"`
	Quantity    int `json:"quantity"`
}

type ResponseExchange struct {
	FromData *Stock `json:"from_data"`
	ToData   *Stock `json:"to_data"`
}

// task2
type EmployeeReport struct {
	EmployeeFullName string  `json:"employee"`
	StoreName        string  `json:"store_name"`
	CategoryName     string  `json:"category"`
	ProductName      string  `json:"product"`
	Quantity         int     `json:"quantity"`
	TotalPrice       float64 `json:"total_price"`
	Date             string  `json:"date"`
}

type GetListEmployeeReportRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListEmployeeReportResponse struct {
	Count           int               `json:"count"`
	EmployeeReports []*EmployeeReport `json:"employee_reports"`
}

// task 4

type TotalOrderPrice struct {
	TotalPrice  float64
	PromoCode   string
	Discount    float64
	ResultPrice float64
}

// task6
type StockProductData struct {
	StoreId      int       `json:"store_id"`
	ProductId    int       `json:"product_id"`
	ProductName  string    `json:"product_name"`
	BrandId      int       `json:"brand_id"`
	BrandData    *Brand    `json:"brand_data"`
	CategoryId   int       `json:"category_id"`
	CategoryData *Category `json:"category_data"`
	ModelYear    int       `json:"model_year"`
	ListPrice    float64   `json:"list_price"`
	Quantity     int       `json:"quantity"`
}

type CategoryStockProduct struct {
	CategoryId           int                `json:"category_id"`
	CategoryName         string             `json:"category_name"`
	Quantity             int                `json:"quantity"`
	CategoryShopProducts []StockProductData `json:"products_data"`
}

type CategoryStockProductData struct {
	CategoryId   int    `json:"category_id"`
	CategoryName string `json:"category_name"`
	Quantity     int    `json:"quantity"`
}
