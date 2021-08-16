package main

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initMysql() {
	config := loadConfigFileName(".", ".env")
	fmt.Println(config)
	ConnectInstance = NewConnect("mysql", config)
	if ConnectInstance.GetConnection().(*sql.DB) != nil {
		fmt.Println("Connected")
	} else {
		fmt.Println("Not Connect")
	}

}

func Test_mysql(t *testing.T) {
	initMysql()
	assert.True(t, ConnectInstance != nil)
	defer ConnectInstance.Disconnect()
}
