package routes

import (
	"github.com/aswinap13/handlers"
	"github.com/aswinap13/services"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//Group the routes if more routes come
	arrayService := services.NewArrayService()
	arrayHandler := handlers.NewArrayHandler(arrayService)

	//Group new routes accordingly
	v1 := r.Group("/api/v1")
	{
		arrays := v1.Group("/arrays")
		{
			arrays.POST("/find-pairs", arrayHandler.FindPairsWithSum)
		}
	}

	return r
}
