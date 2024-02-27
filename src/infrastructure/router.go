package infrastructure

import (
	controllers "Parameters/src/interfaces/api"
	"Parameters/src/interfaces/database"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Init(g *echo.Group, db database.DbHandler) {
	categoryController := controllers.NewCategoryController(db)
	category := g.Group("/category")

	category.GET("", func(c echo.Context) error {
		category := categoryController.GetAllCategories()
		return c.JSON(http.StatusOK, category)
	})

	category.PUT("/:id", func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		categoryController.UpdateCategory(id, c)
		return c.JSON(http.StatusOK, id)
	})

	//category.DELETE("/:id", func(c echo.Context) error {
	//id:
	//})
}
