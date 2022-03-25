package infrastructure

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/so-heee/go-rest-api/api/interfaces/database"
)

type SQLConfig struct {
	user     string
	password string
	host     string
	port     int
	dbname   string
}

func NewSQLConfig(user, password, host, dbname string, port int) SQLConfig {
	return SQLConfig{
		user,
		password,
		host,
		port,
		dbname,
	}
}

type SQLHandler struct {
	conn *gorm.DB
}

func NewSQLHandler(conf *SQLConfig) (database.SQLHandler, error) {
	uri := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&collation=utf8mb4_general_ci", conf.user, conf.password, conf.host, conf.port, conf.dbname)
	log.Info(uri)
	conn, err := gorm.Open(mysql.Open(uri), &gorm.Config{
		Logger: logger.Default.LogMode(2),
	})
	if err != nil {
		log.Fatalf("db connect error: %v", err)
	}
	sqlHandler := new(SQLHandler)
	sqlHandler.conn = conn
	return sqlHandler, nil
}

func (handler *SQLHandler) Find(out interface{}, where ...interface{}) database.SQLHandler {
	db := handler.conn.Find(out, where...)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) First(out interface{}, where ...interface{}) database.SQLHandler {
	db := handler.conn.First(out, where...)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Create(value interface{}) database.SQLHandler {
	db := handler.conn.Create(value)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Delete(value interface{}, where ...interface{}) database.SQLHandler {
	db := handler.conn.Delete(value, where...)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Where(query interface{}, args ...interface{}) database.SQLHandler {
	db := handler.conn.Where(query, args...)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Model(value interface{}) database.SQLHandler {
	db := handler.conn.Model(value)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Update(column string, value interface{}) database.SQLHandler {
	db := handler.conn.Update(column, value)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Updates(values interface{}) database.SQLHandler {
	db := handler.conn.Updates(values)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Begin() database.SQLHandler {
	tx := handler.conn.Begin()
	return &SQLHandler{conn: tx}
}

func (handler *SQLHandler) Commit() database.SQLHandler {
	db := handler.conn.Commit()
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Preload(query string, args ...interface{}) database.SQLHandler {
	db := handler.conn.Preload(query, args...)
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Rollback() database.SQLHandler {
	db := handler.conn.Rollback()
	return &SQLHandler{conn: db}
}

func (handler *SQLHandler) Transaction(fc func(database.SQLHandler) error) error {
	ffc := func(tx *gorm.DB) error {
		driver := &SQLHandler{conn: tx}
		return fc(driver)
	}
	return handler.conn.Transaction(ffc)
}

func (handler *SQLHandler) Error() error {
	return handler.conn.Error
}

// Interface guards
var (
	_ database.SQLHandler = (*SQLHandler)(nil)
)
