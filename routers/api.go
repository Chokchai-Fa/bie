package routers

import (
	"fastmath/controllers"
	"fastmath/repositories"
	"fastmath/services"
	"fastmath/utils/config"

	"github.com/gin-gonic/gin"
)

func initAPIRoute(group *gin.RouterGroup, cfg config.AppConfig) {
	{
		allRepo := services.AllRepository{
			ApiRepository: repositories.NewApiRepository(cfg.DBPGCli),
		}

		apiService := services.NewApiService(allRepo, cfg)

		allSerivce := services.AllService{
			ApiService: apiService,
		}

		questionHandler := controllers.NewApiController(allSerivce)

		group.POST("get/:uuid", questionHandler.Get)
		group.POST("post/:uuid", questionHandler.Create)
		group.POST("put/:uuid", questionHandler.Update)
		group.POST("delete/:uuid", questionHandler.Delete)
	}
}
