package handler

import (
	"app/api/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create PromoCode godoc
// @ID create_promo_code
// @Router /promo_code [POST]
// @Summary Create PromoCode
// @Description Create PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param promoCode body models.CreatePromoCode true "CreatePromoCodeRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreatePromoCode(c *gin.Context) {

	var createPromoCode models.CreatePromoCode

	err := c.ShouldBindJSON(&createPromoCode) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "create promo_code", http.StatusBadRequest, err.Error())
		return
	}

	if createPromoCode.DiscountType == "percent" && (createPromoCode.Discount > 100 || createPromoCode.Discount < 0) {
		h.handlerResponse(c, "create promo_code", http.StatusBadRequest, "Discount percent error")
		return
	}

	id, err := h.storages.PromoCode().Create(context.Background(), &createPromoCode)
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.create", http.StatusInternalServerError, err.Error())
		return
	}

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{Name: id})
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "create promo_code", http.StatusCreated, resp)
}

// Get By ID PromoCode godoc
// @ID get_by_id_promo_code
// @Router /promo_code/{id} [GET]
// @Summary Get By ID PromoCode
// @Description Get By ID PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIdPromoCode(c *gin.Context) {

	promoName := c.Param("id")

	resp, err := h.storages.PromoCode().GetByID(context.Background(), &models.PromoCodePrimaryKey{Name: promoName})
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get promo_code by id", http.StatusCreated, resp)
}

// Get List PromoCode godoc
// @ID get_list_promo_code
// @Router /promo_code [GET]
// @Summary Get List PromoCode
// @Description Get List PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param offset query string false "offset"
// @Param limit query string false "limit"
// @Param search query string false "search"
// @Success 200 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListPromoCode(c *gin.Context) {

	offset, err := h.getOffsetQuery(c.Query("offset"))
	if err != nil {
		h.handlerResponse(c, "get list promo_code", http.StatusBadRequest, "invalid offset")
		return
	}

	limit, err := h.getLimitQuery(c.Query("limit"))
	if err != nil {
		h.handlerResponse(c, "get list promo_code", http.StatusBadRequest, "invalid limit")
		return
	}

	resp, err := h.storages.PromoCode().GetList(context.Background(), &models.GetListPromoCodeRequest{
		Offset: offset,
		Limit:  limit,
		Search: c.Query("search"),
	})
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.getlist", http.StatusInternalServerError, err.Error())
		return
	}

	h.handlerResponse(c, "get list promo_code response", http.StatusOK, resp)
}

// DELETE PromoCode godoc
// @ID delete_promo_code
// @Router /promo_code/{id} [DELETE]
// @Summary Delete PromoCode
// @Description Delete PromoCode
// @Tags PromoCode
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param promo_code body models.PromoCodePrimaryKey true "DeletePromoCodeRequest"
// @Success 204 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeletePromoCode(c *gin.Context) {

	promoName := c.Param("id")

	rowsAffected, err := h.storages.PromoCode().Delete(context.Background(), &models.PromoCodePrimaryKey{Name: promoName})
	if err != nil {
		h.handlerResponse(c, "storage.promo_code.delete", http.StatusInternalServerError, err.Error())
		return
	}
	if rowsAffected <= 0 {
		h.handlerResponse(c, "storage.promo_code.delete", http.StatusBadRequest, "now rows affected")
		return
	}

	h.handlerResponse(c, "delete promo_code", http.StatusNoContent, nil)
}
