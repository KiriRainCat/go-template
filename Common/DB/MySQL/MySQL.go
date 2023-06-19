package MySQL

import (
	"fmt"
	"go-template/Common/Utils/ConfigUtil"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB // 指向Database的指针

func Init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		ConfigUtil.ReadConfigVal("database.username"),
		ConfigUtil.ReadConfigVal("database.password"),
		ConfigUtil.ReadConfigVal("database.host"),
		ConfigUtil.ReadConfigVal("database.name"))

	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ConfigUtil.ReadConfigVal("database.table_prefix").(string), // 表前缀
			SingularTable: true,                                                       // 禁用表名复数
		},
	})
	if err != nil {
		log.Panicf("Database connection failed: %e", err)
	}
}
