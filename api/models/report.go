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
