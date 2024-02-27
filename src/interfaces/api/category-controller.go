package controllers

import (
	"Parameters/src/domain"
	"Parameters/src/interfaces/database"
	"Parameters/src/usecase/category"
	"github.com/labstack/echo"
)

type CategoryController struct {
	Interactor category.CategoryInteractor
}

func NewCategoryController(dbHandler database.DbHandler) *CategoryController {
	return &CategoryController{
		Interactor: category.CategoryInteractor{
			CategoryUsecase: &database.CategoryRepository{
				DbHandler: dbHandler,
			},
		},
	}
}

func (controller *CategoryController) GetAllCategories() []domain.Category {
	res := controller.Interactor.GetAllCategories()
	return res
}

func (controller *CategoryController) DeleteCategory(id int) {
	controller.Interactor.DeleteCategory(id)
}

func (controller *CategoryController) CreateCategory(c echo.Context) {
	category := domain.Category{}
	c.Bind(&category)
	controller.Interactor.CreateCategory(category)
}

func (controller *CategoryController) UpdateCategory(id int, c echo.Context) {
	category := domain.Category{}
	c.Bind(&category)
	controller.Interactor.UpdateCategory(category, id)
}
