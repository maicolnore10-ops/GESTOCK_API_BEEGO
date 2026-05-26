package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Notificacion_20260525_210707 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Notificacion_20260525_210707{}
	m.Created = "20260525_210707"

	migration.Register("Notificacion_20260525_210707", m)
}

// Run the migrations
func (m *Notificacion_20260525_210707) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update

}

// Reverse the migrations
func (m *Notificacion_20260525_210707) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update

}
