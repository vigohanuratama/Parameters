package infrastructure

import (
	"Parameters/src/interfaces/database"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DbHandler struct {
	Db *gorm.DB
}

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}

func NewDbHandler(postgresConfig *Config) (database.DbHandler, error) {
	host := postgresConfig.Host
	port := postgresConfig.Port
	user := postgresConfig.User
	password := postgresConfig.Password
	database := postgresConfig.Database

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", host, port, user, password, database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	if err != nil {
		return nil, err
	}

	return &DbHandler{Db: db}, nil
}

func (d DbHandler) FindAll(object interface{}) {
	d.Db.Find(object)
}

func (d DbHandler) Create(object interface{}) {
	d.Db.Create(object)
}

func (d DbHandler) Delete(object interface{}, id any) {
	d.Db.Delete(object, id)
}

func (d DbHandler) Update(object interface{}) {
	d.Db.Save(object)
}
