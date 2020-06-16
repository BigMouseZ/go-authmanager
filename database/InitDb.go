package database

import "github.com/jinzhu/gorm"
import _ "github.com/jinzhu/gorm/dialects/postgres"
var (
	Db *gorm.DB
)
func Dbinit() *gorm.DB {
	Db := NewConn()
	//SetMaxOpenConns用于设置最大打开的连接数
	//SetMaxIdleConns用于设置闲置的连接数
	Db.DB().SetMaxIdleConns(10)
	Db.DB().SetMaxOpenConns(100)
	// 启用Logger，显示详细日志
	Db.LogMode(true)
	// 自动迁移模式
	//db.AutoMigrate(&Model.UserModel{},
	//	&Model.UserDetailModel{},
	//	&Model.UserAuthsModel{},
	//)
	return Db
}

func NewConn() *gorm.DB {
	var err error
	Db, err = gorm.Open("postgres", "host=192.168.147.129 port=5432 user=postgres dbname=authtest password=1234zxcv sslmode=disable")
	if err != nil {
		panic("连接数据库失败:" + err.Error())
	}
	return Db
}
