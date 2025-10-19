package handler

import "github.com/gin-gonic/gin"

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler{
	return &Handler{services: services}
}

//функция инициализации маршрутов для api
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//объект для группы роутера для авторизации
	auth := router.Group("/auth")
	{
		//здесь обозначаем ендпоинты для логина и регистрации
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	//объект группы роутера для api
	api := router.Group("/api")
	{
		// мы можем создавать объекты из объектов, чтобы расширять ендпоинты
		lists := api.Group("/lists")
		{
			lists.POST("/", h.createList)
			lists.GET("/", h.getAllLists)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)

			//расширение ендпоинтов можно производить многократно
			items := lists.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItems)
				items.GET("/:item_id", h.getItemById)
				items.PUT("/:item_id", h.updateItem)
				items.DELETE("/:item_id", h.deleteItem)
			}
		}
	}
	return router
}
