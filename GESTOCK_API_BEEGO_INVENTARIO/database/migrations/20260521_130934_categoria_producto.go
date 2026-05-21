package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type CategoriaProducto_20260521_130934 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CategoriaProducto_20260521_130934{}
	m.Created = "20260521_130934"

	migration.Register("CategoriaProducto_20260521_130934", m)
}

// Run the migrations
func (m *CategoriaProducto_20260521_130934) Up() {
	m.SQL("INSERT INTO inventario.categoria_producto( 'nombre' ) VALUES ( 'Electrónica');") // to make schema update

}

// Reverse the migrations
func (m *CategoriaProducto_20260521_130934) Down() {
m.SQL("DELETE FROM inventario.categoria_producto WHERE nombre = 'Electrónica';")  // to reverse schema update

}
