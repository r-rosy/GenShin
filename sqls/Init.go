package sqls

import (
	"fmt"
	"game/info/sqldet"
	"game/info/tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlDb struct {
	Db *gorm.DB
}

func initConnection() *SqlDb {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", sqldet.Username, sqldet.Password, sqldet.Host, sqldet.Port, sqldet.Dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	sqldb := &SqlDb{
		db,
	}
	return sqldb
}
func (sqldb *SqlDb) tablesInit() {
	sqldb.Db.AutoMigrate(&tables.User{})
	sqldb.Db.AutoMigrate(&tables.Weapon{})
	sqldb.Db.AutoMigrate(&tables.Character{})
	sqldb.Db.AutoMigrate(&tables.Cookie{})
	sqldb.Db.AutoMigrate(&tables.GodLeftValue{})
}
func tableinit() *SqlDb {
	sqldb := initConnection()
	sqldb.tablesInit()
	return sqldb
}
