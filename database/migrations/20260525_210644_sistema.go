package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Sistema_20260525_210644 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Sistema_20260525_210644{}
	m.Created = "20260525_210644"

	migration.Register("Sistema_20260525_210644", m)
}

// Run the migrations
func (m *Sistema_20260525_210644) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Sistema_20260525_210644) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
