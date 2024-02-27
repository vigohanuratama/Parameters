package category

import "Parameters/src/domain"

type CategoryInteractor struct {
	CategoryUsecase CategoryUsecase
}

func (interactor *CategoryInteractor) GetAllCategories() []domain.Category {
	return interactor.CategoryUsecase.GetAllCategories()
}

func (interactor *CategoryInteractor) DeleteCategory(id int) {
	interactor.CategoryUsecase.DeleteCategory(id)
}

func (interactor *CategoryInteractor) CreateCategory(u domain.Category) {
	interactor.CategoryUsecase.CreateCategory(u)
}

func (interactor *CategoryInteractor) UpdateCategory(category domain.Category, id int) {
	interactor.CategoryUsecase.UpdateCategory(category, id)
}
