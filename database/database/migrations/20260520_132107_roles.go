package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Roles_20260520_132107 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Roles_20260520_132107{}
	m.Created = "20260520_132107"

	migration.Register("Roles_20260520_132107", m)
}

// Run the migrations
func (m *Roles_20260520_132107) Up() {
	m.SQL("INSERT INTO autenticacion.roles (nombre_rol ) VALUES ('Administrador')")
	m.SQL("INSERT INTO autenticacion.roles (nombre_rol ) VALUES ('Operador')")

}

// Reverse the migrations
func (m *Roles_20260520_132107) Down() {
	m.SQL("DELETE FROM autenticacion.roles WHERE nombre_rol IN ('Administrador', 'Operador')")

}
