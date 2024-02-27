package database

type DbHandler interface {
	FindAll(object interface{})
	Create(object interface{})
	Delete(object interface{}, id any)
	Update(object interface{})
}
