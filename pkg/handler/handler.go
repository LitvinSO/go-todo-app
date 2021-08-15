package handler

import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	auth := router.Group("/path")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		lists := api.Group("/list")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			item := api.Group("/")
			{
				item.POST("/", h.createItem)
				item.GET("/", h.getAllItem)
				item.GET("/:item_id", h.getItemById)
				item.PUT("/:item_id", h.updateItem)
				item.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
