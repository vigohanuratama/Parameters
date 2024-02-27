package database

import "Parameters/src/domain"

type CategoryRepository struct {
	DbHandler
}

func (db *CategoryRepository) UpdateCategory(category domain.Category, id int) {
	category.ID = id
	db.Update(&category)
}

func (db *CategoryRepository) DeleteCategory(id int) {
	category := []domain.Category{}
	db.Delete(&category, id)
}

func (db *CategoryRepository) CreateCategory(category domain.Category) {
	db.Create(&category)
}

func (db *CategoryRepository) GetAllCategories() []domain.Category {
	category := []domain.Category{}
	db.FindAll(&category)
	return category
}
