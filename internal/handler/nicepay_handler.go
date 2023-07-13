package handler

import (
	"fmt"
	"net/http"

	"github.com/doffy007/golang-nicepay/internal/constants"
	"github.com/doffy007/golang-nicepay/internal/request"
	"github.com/doffy007/golang-nicepay/internal/usecase"
	"github.com/doffy007/golang-nicepay/internal/utils"
	"github.com/gin-gonic/gin"
)

type NicepayHandler interface {
	NicepayRegister(c *gin.Context)
	NicepayStatus(c *gin.Context)
	NicepayPayment(c *gin.Context)
}

type nicepayHandler struct {
	usecase usecase.NicepayUsecase
}

func (h handler) NicepayHandler() NicepayHandler {
	return nicepayHandler{
		usecase: h.usecase.NicepayHandler(),
	}
}

// NicepayRegister implements.
func (h nicepayHandler) NicepayRegister(c *gin.Context) {
	var params request.RegistrationRequest

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayRegister Parse Param: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusBadRequest, constants.ERROR_PARSE_PARAM, err.Error())
		return
	}

	ok, res := h.usecase.NicepayRegistration(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayRegister From Usecase: %v\n", res.ResultMessage)
		utils.NewErrorResponse(c, res.ResultCode, res.ResultMessage, []string{res.ResultMessage})
		return
	}

	utils.NewSuccessResponse(c, res.ResultCode, res.ResultMessage, res)
}

// NcepayStatus implements.
func (h nicepayHandler) NicepayStatus(c *gin.Context) {
	var params request.StatusRequest

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayStatus Parse Param: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusBadRequest, constants.ERROR_PARSE_PARAM, err.Error())
		return
	}

	ok, res := h.usecase.NicepayStatus(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayStatus From Usecase: %v\n", res.ResultMessage)
		utils.NewErrorResponse(c, res.ResultCode, res.ResultMessage, []string{res.ResultMessage})
		return
	}

	utils.NewSuccessResponse(c, res.ResultCode, res.ResultMessage, res)
}

// NicepayPayment implements.
func (h nicepayHandler) NicepayPayment(c *gin.Context) {
	var params request.PaymentRequest

	if err := c.ShouldBind(&params); err != nil {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayStatus Parse Param: %v\n", err.Error())
		utils.NewErrorResponse(c, http.StatusBadRequest, constants.ERROR_PARSE_PARAM, err.Error())
		return
	}

	ok, res := h.usecase.NicepayPayment(params)
	if !ok {
		fmt.Printf("\033[1;31m [ERROR] \033[0m Handler NicepayStatus From Usecase: %v\n", res.ResultMessage)
		utils.NewErrorResponse(c, res.ResultCode, res.ResultMessage, []string{res.ResultMessage})
		return
	}

	utils.NewSuccessResponse(c, res.ResultCode, res.ResultMessage, res)
}
