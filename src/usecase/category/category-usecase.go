package category

import "Parameters/src/domain"

type CategoryUsecase interface {
	CreateCategory(category domain.Category)
	GetAllCategories() []domain.Category
	DeleteCategory(id int)
	UpdateCategory(category domain.Category, id int)
}
