package controllers

import (
	"io"

	"fastmath/services"
	"fastmath/utils/common"
	"fastmath/utils/errors"
	"fastmath/utils/handlers"
	"fastmath/utils/logs"

	"github.com/gin-gonic/gin"
)

type ApiController struct {
	apiService services.ApiService
}

func NewApiController(services services.AllService) *ApiController {
	return &ApiController{
		apiService: services.ApiService,
	}
}

func (h *ApiController) Get(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)

	var uuid string

	if err := common.BytesToStruct(body, &uuid); err != nil {
		err := errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	if err := common.ValidatorParam(uuid); err != nil {
		err = errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	res, err := h.apiService.Get(uuid, ctx)

	if err != nil {
		handlers.HandleError(ctx, err)
		return
	}

	handlers.HandleSuccess(ctx, res)
}

func (h *ApiController) Create(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)

	var uuid string

	if err := common.BytesToStruct(body, &uuid); err != nil {
		err := errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	if err := common.ValidatorParam(uuid); err != nil {
		err = errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	res, err := h.apiService.Create(uuid, ctx)

	if err != nil {
		handlers.HandleError(ctx, err)
		return
	}

	handlers.HandleSuccess(ctx, res)
}

func (h *ApiController) Update(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)

	var uuid string

	if err := common.BytesToStruct(body, &uuid); err != nil {
		err := errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	if err := common.ValidatorParam(uuid); err != nil {
		err = errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	res, err := h.apiService.Update(uuid, ctx)

	if err != nil {
		handlers.HandleError(ctx, err)
		return
	}

	handlers.HandleSuccess(ctx, res)
}

func (h *ApiController) Delete(ctx *gin.Context) {
	body, _ := io.ReadAll(ctx.Request.Body)

	var uuid string

	if err := common.BytesToStruct(body, &uuid); err != nil {
		err := errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	if err := common.ValidatorParam(uuid); err != nil {
		err = errors.NewValidationError(err.Error())
		logs.Error(err)
		handlers.HandleError(ctx, err)
		return
	}

	res, err := h.apiService.Delete(uuid, ctx)

	if err != nil {
		handlers.HandleError(ctx, err)
		return
	}

	handlers.HandleSuccess(ctx, res)
}
