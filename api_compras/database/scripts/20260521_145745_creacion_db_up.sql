CREATE SCHEMA IF NOT EXISTS compras;

-- 1. Crear tabla proveedores
CREATE TABLE IF NOT EXISTS compras.proveedores (
    id SERIAL PRIMARY KEY,
    id_empresa INT,
    nombre VARCHAR(150) NOT NULL,
    correo VARCHAR(150),
    telefono VARCHAR(50),
    activo BOOLEAN DEFAULT TRUE
);

-- Datos de prueba para proveedores
INSERT INTO compras.proveedores (id_empresa, nombre, correo, telefono, activo) 
VALUES (1, 'Ander', 'Ander29@gmail.com', '3223656923', TRUE)
ON CONFLICT DO NOTHING;

-- 2. Crear tabla proveedor_productos
CREATE TABLE IF NOT EXISTS compras.proveedor_productos (
    id_proveedor_producto SERIAL PRIMARY KEY,
    id_proveedor INT,
    id_producto INT,
    precio NUMERIC(10,2),
    activo BOOLEAN DEFAULT TRUE,
    CONSTRAINT fk_productos_proveedor 
        FOREIGN KEY (id_proveedor) 
        REFERENCES compras.proveedores(id) 
        ON DELETE RESTRICT
);

-- Datos de prueba para proveedor_productos
INSERT INTO compras.proveedor_productos (id_proveedor, id_producto, precio) 
VALUES (1, 10, 250000)
ON CONFLICT DO NOTHING;

-- 3. Crear tabla guias_entrada
CREATE TABLE IF NOT EXISTS compras.guias_entrada (
    id_guia SERIAL PRIMARY KEY,
    id_proveedor INT,
    id_bodega INT,
    estado VARCHAR(50),
    activo BOOLEAN NOT NULL DEFAULT TRUE,
    CONSTRAINT fk_guias_entrada_proveedor
        FOREIGN KEY (id_proveedor)
        REFERENCES compras.proveedores(id)
        ON DELETE RESTRICT
);

-- Datos de prueba para guias_entrada
INSERT INTO compras.guias_entrada (id_proveedor, id_bodega, estado, activo) 
VALUES (1, 1, 'Recibido', TRUE)
ON CONFLICT DO NOTHING;