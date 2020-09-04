package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type Mysql struct {
	Db *gorm.DB
}

func NewMysqlConn()  *Mysql{
	conn := Connect()
	return &Mysql{
		Db : conn,
	}
}

func Connect() *gorm.DB{
	conn, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/information_schema?charset=utf8&parseTime=True&loc=Local")
	defer conn.Close()

	if err != nil {
		fmt.Println(err);
		panic(err)
	}

	return conn;
}
