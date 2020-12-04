package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	//仅引入init方法
	"gin-blog/pkg/setting"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	//``内容为成员变量标签
	/**
	u := &User{UserId: 1, UserName: "tony"}
	j, _ := json.Marshal(u)
	fmt.Println(string(j))
	// 输出内容：{"user_id":1,"user_name":"tony"}
	*/
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

//golang程序初始化先于main函数执行，由runtime进行初始化
func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	//为表名添加前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	//禁用复数形式
	db.SingularTable(true)
	//设置最大连接数和最大闲置数
	// 默认0表示不限制
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
func CloseDB() {
	defer db.Close()
}
