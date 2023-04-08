package handler

import (
	"app/api/models"
	"app/config"
	"app/pkg/helper"
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Create Register
// @Description Create Register
// @Tags Register
// @Accept json
// @Produce json
// @Param register body models.Register true "CreateRegisterRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Register(c *gin.Context) {

	var createUser models.CreateUser

	err := c.ShouldBindJSON(&createUser) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "register user", http.StatusBadRequest, err.Error())
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(createUser.Password), 7)
	if err != nil {
		h.handlerResponse(c, "error while hashing password", http.StatusBadRequest, err.Error())
		return
	}

	createUser.Password = string(hashedPassword)

	id, err := h.storages.User().Create(context.Background(), &createUser)
	if err != nil {
		if err.Error() == `ERROR: duplicate key value violates unique constraint "users_login_key" (SQLSTATE 23505)` {
			h.handlerResponse(c, "storage.user.create", http.StatusBadRequest, "user already exists please login!")
			return
		}
		h.handlerResponse(c, "storage.user.create", http.StatusInternalServerError, err.Error())
		return
	}

	user, err := h.storages.User().GetByID(context.Background(), &models.UserPrimaryKey{Id: id})
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, user)
}

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Create Login
// @Description Create Login
// @Tags Login
// @Accept json
// @Produce json
// @Param logim body models.Login true "LoginRequest"
// @Success 201 {object} Response{data=string} "Success Request"
// @Response 400 {object} Response{data=string} "Bad Request"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {

	var login models.Login

	err := c.ShouldBindJSON(&login) // parse req body to given type struct
	if err != nil {
		h.handlerResponse(c, "login user", http.StatusBadRequest, err.Error())
		return
	}

	resp, err := h.storages.User().GetByID(context.Background(), &models.UserPrimaryKey{Login: login.Login})
	if err != nil {
		if err.Error() == "no rows in result set" {
			h.handlerResponse(c, "storage.user.getByID", http.StatusBadRequest, "user not found please register first")
			return
		}

		h.handlerResponse(c, "storage.user.getByID", http.StatusInternalServerError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(resp.Password), []byte(login.Password))
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusBadRequest, "credentials are wrong")
		return
	}

	data := map[string]interface{}{
		"Id":         resp.Id,
		"name":       resp.Name,
		"login":      resp.Login,
		"password":   resp.Password,
		"created_at": resp.CreatedAt,
		"updated_at": resp.UpdatedAt,
	}

	token, err := helper.GenerateJWT(data, config.TimeExpiredAt, h.cfg.SecretKey)
	if err != nil {
		h.handlerResponse(c, "storage.user.getByID", http.StatusBadRequest, errors.New("token error"))
		return
	}
	// var bearer = "Bearer " + token

	c.JSON(http.StatusCreated, models.LoginResponse{Token: token, UserData: resp})
}
