package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Empresas_20260525_133419 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Empresas_20260525_133419{}
	m.Created = "20260525_133419"

	migration.Register("Empresas_20260525_133419", m)
}

// Run the migrations
func (m *Empresas_20260525_133419) Up() {
	m.SQL("INSERT INTO empresa.empresas (nombre_empresa, correo) VALUES ('Empresa A', 'empresa@empresa.com')")
	m.SQL("INSERT INTO empresa.empresas (nombre_empresa, correo) VALUES ('Empresa B', 'empresa@empresa.com')")

}

// Reverse the migrations
func (m *Empresas_20260525_133419) Down() {
	m.SQL("DELETE FROM empresa.empresas WHERE nombre_empresa = 'Empresa A'")
	m.SQL("DELETE FROM empresa.empresas WHERE nombre_empresa = 'Empresa B'")

}
