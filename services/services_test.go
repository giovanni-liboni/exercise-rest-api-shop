package services

import "github.com/DATA-DOG/go-txdb"

func init() {
	txdb.Register("mysql_txdb", "mysql", "root:secret@/test?multiStatements=true&parseTime=true")
}
