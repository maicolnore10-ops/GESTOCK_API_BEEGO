package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Usuarios_20260520_132117 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Usuarios_20260520_132117{}
	m.Created = "20260520_132117"

	migration.Register("Usuarios_20260520_132117", m)
}

// Run the migrations
func (m *Usuarios_20260520_132117) Up() {
	m.SQL("INSERT INTO autenticacion.usuarios (correo, nombres, apellidos, telefono, fecha_nacimiento, documento) VALUES ('MAICOL@gestock.com', 'Maicol', 'Nore', '3145571810', '2007-03-31', 1082895940)")
	m.SQL("INSERT INTO autenticacion.usuarios (correo, nombres, apellidos, telefono, fecha_nacimiento, documento) VALUES ('MONO@gestock.com', 'MONO', 'Fonseca', '3106863520', '2000-01-01', 1024548654)")

}

// Reverse the migrations
func (m *Usuarios_20260520_132117) Down() {
	m.SQL("DELETE FROM autenticacion.usuarios WHERE correo IN ('MAICOL@gestock.com', 'MONO@gestock.com')")

}
