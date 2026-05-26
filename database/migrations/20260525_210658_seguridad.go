package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Seguridad_20260525_210658 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Seguridad_20260525_210658{}
	m.Created = "20260525_210658"

	migration.Register("Seguridad_20260525_210658", m)
}

// Run the migrations
func (m *Seguridad_20260525_210658) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Seguridad_20260525_210658) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
