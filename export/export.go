package export

import (
	"dbSchema/model"
	"fmt"
)

type DataFormat interface {
	WriteHeader(columns []string) error
	WriteRow(values map[string]interface{}) error
	WriteCustomV(record []string) error
	Flush() error
}



func Export(format DataFormat)  {
	tables, err := model.GetAllTables("test_db");// db名称
	if err != nil{
		panic(err)
	}

	fmt.Println(tables)


	for _, tbl := range tables {
		h := []string{"字段名称","字段类型","约束","为空","字段含义"}
		err = format.WriteHeader(h)
		if err != nil {
			panic(err)
		}
		cols, _ :=model.GetAllTableColumns(tbl.TableSchema, tbl.Tablename)

		for _, col := range cols {
			//fmt.Printf("%v",col)
			r := []string{
				col.ColumnName,
				col.ColumnType,
				col.ColumnKey,
				col.IsNullable,
				col.ColumnComment,
			}
			err := format.WriteCustomV(r)
			if err !=nil {
				panic(err)
			}
		}

	}
	fmt.Println("export data ok！！！")
}
