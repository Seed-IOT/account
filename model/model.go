package model

import (
	"account/config"
	"database/sql"
	"regexp"

	"github.com/gomodule/redigo/redis"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"

	_ "github.com/jinzhu/gorm/dialects/mysql" //mysql dialects
)

// Service is the interface of all model service.
// type Service interface {
// 	GetDB() *gorm.DB
// 	Login(account string, password string) (*UserInfo, error)
// }

// Service is the interface of all model service.
type Service struct {
	config    *config.Config
	DB        *gorm.DB
	redisConn redis.Conn
}

var db *gorm.DB

// New returns a Service instance for operating all model service.
func New(dbCfg *config.Database) (*Service, error) {
	db, err := newDB(dbCfg) //needs to pass db as
	serv := &Service{}
	if err != nil {
		return serv, errors.Wrap(err, "Failed to initialize db of grom")
	}

	serv.DB = db

	// 创建RedisConn
	redisConn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		return serv, errors.Wrap(err, "Failed to redis")
	}
	serv.redisConn = redisConn

	return serv, nil
}

func newDB(dbCfg *config.Database) (*gorm.DB, error) {
	db, err := gorm.Open("mysql", dbCfg.URL)
	// 数据库不存在 则创建数据库
	if err != nil {
		db, err := initDB(dbCfg.URL)
		return db, err
	}
	db.DB().SetMaxOpenConns(dbCfg.MaxActive)
	db.DB().SetMaxIdleConns(dbCfg.MaxIdle)

	db.LogMode(dbCfg.LogMode)

	db.AutoMigrate(&AppInfo{})
	db.AutoMigrate(&UserInfo{})

	// 检查表是否存在，否则创建

	return db, nil
}

// 初始化数据库
func initDB(url string) (*gorm.DB, error) {
	// 需要转换 root:Kylewang1331@/sys?charset=utf8&parseTime=True&loc=Local
	re3, _ := regexp.Compile(`\@\/(.*?)\?`)
	rep := re3.ReplaceAllString(url, "@/sys?")
	basedb, err := sql.Open("mysql", rep)
	stmt, err := basedb.Prepare("CREATE DATABASE  IF NOT EXISTS `seed`")
	stmt.Exec()
	basedb.Close()
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to open db URL")
	}
	return db, nil
}
