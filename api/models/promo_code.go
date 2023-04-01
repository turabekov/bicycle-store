package models

// task 3

type PromoCode struct {
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type PromoCodePrimaryKey struct {
	Name string `json:"name"`
}

type CreatePromoCode struct {
	Name            string  `json:"name"`
	Discount        float64 `json:"discount"`
	DiscountType    string  `json:"discount_type"`
	OrderLimitPrice float64 `json:"order_limit_price"`
}

type GetListPromoCodeRequest struct {
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
	Search string `json:"search"`
}

type GetListPromoCodeResponse struct {
	Count      int          `json:"count"`
	PromoCodes []*PromoCode `json:"promo_codes"`
}
