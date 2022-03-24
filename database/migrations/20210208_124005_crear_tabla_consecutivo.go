package main

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CrearTablaConsecutivo_20210208_124005 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CrearTablaConsecutivo_20210208_124005{}
	m.Created = "20210208_124005"

	migration.Register("CrearTablaConsecutivo_20210208_124005", m)
}

// Run the migrations
func (m *CrearTablaConsecutivo_20210208_124005) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	file, err := ioutil.ReadFile("../scripts/20210208_124005_crear_tabla_consecutivo_up.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}

// Reverse the migrations
func (m *CrearTablaConsecutivo_20210208_124005) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	file, err := ioutil.ReadFile("../scripts/20210208_124005_crear_tabla_consecutivo_down.sql")

	if err != nil {
		// handle error
		fmt.Println(err)
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		fmt.Println(request)
		m.SQL(request)
		// do whatever you need with result and error
	}

}
