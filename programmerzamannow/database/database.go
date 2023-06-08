package database

import "fmt"

var connection string

func init() {
	//ini bukan konek ke database, hanya ingin memperlihatkan function ini akan langsung dieksekusi ketika di import
	fmt.Println("Function init executed")
	connection = "MySQL"

}

func GetDatabase() string {
	return connection
}
