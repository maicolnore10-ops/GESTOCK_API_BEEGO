package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type UnidadMedida_20260525_133455 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &UnidadMedida_20260525_133455{}
	m.Created = "20260525_133455"

	migration.Register("UnidadMedida_20260525_133455", m)
}

// Run the migrations
func (m *UnidadMedida_20260525_133455) Up() {
	m.SQL("INSERT INTO empresa.unidades_medida (nombre_unidad_medida, sigla_unidad_medida) VALUES ('Kilogramo', 'kg')")
	m.SQL("INSERT INTO empresa.unidades_medida (nombre_unidad_medida, sigla_unidad_medida) VALUES ('Metro', 'm')")
	m.SQL("INSERT INTO empresa.unidades_medida (nombre_unidad_medida, sigla_unidad_medida) VALUES ('Centimetro', 'cm')")
	m.SQL("INSERT INTO empresa.unidades_medida (nombre_unidad_medida, sigla_unidad_medida) VALUES ('Milimetro', 'mm')")

}

// Reverse the migrations
func (m *UnidadMedida_20260525_133455) Down() {
	m.SQL("DELETE FROM empresa.unidades_medida WHERE nombre_unidad_medida = 'Kilogramo'")
	m.SQL("DELETE FROM empresa.unidades_medida WHERE nombre_unidad_medida = 'Metro'")
	m.SQL("DELETE FROM empresa.unidades_medida WHERE nombre_unidad_medida = 'Centimetro'")
	m.SQL("DELETE FROM empresa.unidades_medida WHERE nombre_unidad_medida = 'Milimetro'")

}
