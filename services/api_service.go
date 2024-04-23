package services

import (
	"fastmath/models"
	"fastmath/repositories"

	"github.com/gin-gonic/gin"
)

type ApiService interface {
	Get(uuid string, context *gin.Context) (*models.APIRes, error)
	Create(uuid string, context *gin.Context) (*models.APIRes, error)
	Update(uuid string, context *gin.Context) (*models.APIRes, error)
	Delete(uuid string, context *gin.Context) (*models.APIRes, error)
}

type ApiServiceContext struct {
	repository repositories.ApiRepository
}

func NewApiService(repositories AllRepository) ApiService {
	return &ApiServiceContext{
		repository: &repositories.ApiRepository,
	}
}

func (h *ApiServiceContext) Get(uuid string, context *gin.Context) (*models.APIRes, error) {
	res, err := h.repository.Get(uuid, context)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *ApiServiceContext) Create(uuid string, context *gin.Context) (*models.APIRes, error) {
	res, err := h.repository.Create(uuid, context)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *ApiServiceContext) Update(uuid string, context *gin.Context) (*models.APIRes, error) {
	res, err := h.repository.Update(uuid, context)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (h *ApiServiceContext) Delete(uuid string, context *gin.Context) (*models.APIRes, error) {
	res, err := h.repository.Delete(uuid, context)
	if err != nil {
		return nil, err
	}

	return res, nil
}
