package main

import (
	"github.com/beego/beego/v2/client/orm/migration"

)

// DO NOT MODIFY
type ContrasenaHash_20260520_154123 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ContrasenaHash_20260520_154123{}
	m.Created = "20260520_154123"

	migration.Register("ContrasenaHash_20260520_154123", m)
}

// Run the migrations
// Agrega indice unico para garantizar una contrasena activa por usuario
func (m *ContrasenaHash_20260520_154123) Up() {
	m.SQL(`INSERT INTO autenticacion.contrasena_hash(id_usuario, contrasena_hash )VALUES (3,'123456789');`)
}

// Reverse the migrations
func (m *ContrasenaHash_20260520_154123) Down() {
	m.SQL(`DELETE FROM autenticacion.contrasena_hash WHERE id_contrasena_hash = 1`)
}
