package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type RecuperacionContrasena_20260520_154142 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &RecuperacionContrasena_20260520_154142{}
	m.Created = "20260520_154142"

	migration.Register("RecuperacionContrasena_20260520_154142", m)
}

// Run the migrations
// Agrega indice para acelerar busquedas por codigo de verificacion y usuario
func (m *RecuperacionContrasena_20260520_154142) Up() {
	m.SQL(`INSERT INTO autenticacion.recuperacion_contrasena(id_recuperacion, id_usuario, codigo_verificacion, fecha_expiracion, usado) 
    VALUES (1, 1, 'abc123xyz789def456', NOW() + INTERVAL '1 hour', false);`)

}	
// Reverse the migrations
func (m *RecuperacionContrasena_20260520_154142) Down() {
	m.SQL(`DELETE FROM autenticacion.recuperacion_contrasena WHERE id_recuperacion = 1`)

}
