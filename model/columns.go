package model

import "dbSchema/mysql"

type Columns struct {
	TableSchema string `gorm:"column:TABLE_SCHEMA"`
	Tablename string `gorm:"column:TABLE_NAME"`
	ColumnName string `gorm:"column:COLUMN_NAME"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
	ColumnType string `gorm:"column:COLUMN_TYPE"`
	IsNullable string `gorm:"column:is_nullable_"`
	ColumnKey string `gorm:"column:COLUMN_KEY"`
}


func (*Columns) TableName() string {
	return "COLUMNS"
}

/**
	获取表的所有列
	select column_name,column_comment,column_type,is_nullable, column_key from information_schema.columns  where  table_schema=? and table_name=?  group by column_name
 */
func GetAllTableColumns(tblSchema string, tblName string)  ([]*Columns, error){
	var columns []*Columns
	err := mysql.NewMysqlConn().Db.Where(&Columns{Tablename:tblName, TableSchema:tblSchema}).Find(&columns).Error
	if err != nil {
		return nil, err
	}

	return columns, err
}
