package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type ProveedorProductos_20260521_150319 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &ProveedorProductos_20260521_150319{}
	m.Created = "20260521_150319"

	migration.Register("ProveedorProductos_20260521_150319", m)
}

// Run the migrations
func (m *ProveedorProductos_20260521_150319) Up() {
	m.SQL(`CREATE TABLE IF NOT EXISTS compras.proveedores (...);`)
    m.SQL(`CREATE TABLE IF NOT EXISTS compras.proveedor_productos (
        id_proveedor_producto SERIAL PRIMARY KEY,
        id_proveedor INT,
        id_producto INT,
        precio NUMERIC(10,2),
        activo BOOLEAN DEFAULT true
    );`)

    m.SQL(`INSERT INTO compras.proveedores (...) VALUES (...);`)
    m.SQL(`INSERT INTO compras.proveedor_productos (id_proveedor, id_producto, precio) 
           VALUES (1, 10, 250000) ON CONFLICT DO NOTHING;`)

}

// Reverse the migrations
func (m *ProveedorProductos_20260521_150319) Down() {
	m.SQL(`DELETE FROM compras.proveedor_productos WHERE id_proveedor = 1 AND id_producto = 10;`)
}

