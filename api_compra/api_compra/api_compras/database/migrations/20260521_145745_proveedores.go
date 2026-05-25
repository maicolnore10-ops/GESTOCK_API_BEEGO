package main

import (
	"github.com/beego/beego/v2/client/orm/migration"
)

// DO NOT MODIFY
type Proveedores_20260521_145745 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &Proveedores_20260521_145745{}
	m.Created = "20260521_145745"

	migration.Register("Proveedores_20260521_145745", m)
}

// Run the migrations

// Up: Run the migrations
func (m *Proveedores_20260521_145745) Up() {
    m.SQL(`CREATE TABLE compras.proveedores (
        id SERIAL PRIMARY KEY,
        id_empresa INT,
        nombre VARCHAR(150),
        correo VARCHAR(150),
        telefono VARCHAR(50),
        activo BOOLEAN DEFAULT true
    );`)

    m.SQL(`INSERT INTO compras.proveedores (id_empresa, nombre, correo, telefono, activo) 
           VALUES (1, 'Ander', 'Ander29@gmail.com', '3223656923', true);`)
}


func (m *Proveedores_20260521_145745) Down() {
    // 3. Borrado seguro con CASCADE para limpiar dependencias
    m.SQL(`DROP TABLE IF EXISTS compras.proveedores CASCADE;`)
}