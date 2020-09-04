package main

import (
	"dbSchema/export"
	"dbSchema/format"
	"io"
	"os"
	"unicode/utf8"
)

func main()  {
	writer := parseWriter()
	delimiter, _ := utf8.DecodeLastRuneInString(",")
	f := format.NewCsvFormat(writer,delimiter, false)

	export.Export(f)

	//fmt.Println(data)


}



func parseWriter() io.Writer {
	outputFilename := "./output.csv"
	_, err := os.Stat(outputFilename)
	//err==nil，说明存在
	if err == nil {
		os.Remove(outputFilename)
	}

	if outputFilename != "" {
		f, err := os.Create(outputFilename)
		if err != nil {
			panic(err)
		}
		return f
	}
	return os.Stdout
}

