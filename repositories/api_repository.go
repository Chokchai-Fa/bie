package repositories

import (
	"database/sql"
	"fastmath/models"

	"github.com/gin-gonic/gin"
)

type ApiRepository interface {
	Get(uuid string, ctx *gin.Context) (*models.APIRes, error)
	Create(uuid string, ctx *gin.Context) (*models.APIRes, error)
	Update(uuid string, ctx *gin.Context) (*models.APIRes, error)
	Delete(uuid string, ctx *gin.Context) (*models.APIRes, error)
}

type ApiRepositoryContext struct {
}

func NewApiRepository(db *sql.DB) ApiRepository {
	return &ApiRepositoryContext{}
}

func (h *ApiRepositoryContext) Get(uuid string, ctx *gin.Context) (*models.APIRes, error) {
	return nil, nil
}

func (h *ApiRepositoryContext) Create(uuid string, ctx *gin.Context) (*models.APIRes, error) {
	return nil, nil
}

func (h *ApiRepositoryContext) Update(uuid string, ctx *gin.Context) (*models.APIRes, error) {
	return nil, nil
}

func (h *ApiRepositoryContext) Delete(uuid string, ctx *gin.Context) (*models.APIRes, error) {
	return nil, nil
}
