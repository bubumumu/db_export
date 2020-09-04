package model

import "dbSchema/mysql"

type Tables struct {
	TableSchema string `gorm:"column:TABLE_SCHEMA"`
	Tablename string `gorm:"column:TABLE_NAME"`
	TableComment string `gorm:"column:TABLE_NAME"`
}

func (*Tables) TableName() string {
	return "TABLES"
}


/**
获取某个db下的所有表
select table_name,table_comment from information_schema.tables where table_schema=?
*/
func GetAllTables(dbName string) ([]*Tables, error) {
	var tables []*Tables
	err := mysql.NewMysqlConn().Db.Where(&Tables{TableSchema:dbName}).Find(&tables).Error
	if err != nil {
		return nil, err
	}

	return tables, err
}