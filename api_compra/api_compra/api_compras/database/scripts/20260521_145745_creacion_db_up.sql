CREATE SCHEMA compras;

CREATE TABLE IF NOT EXISTS compras.proveedores compras.proveedores (
    id_proveedor SERIAL PRIMARY KEY,
    nombre VARCHAR(150) NOT NULL,
    correo VARCHAR(150),
    telefono VARCHAR(50),
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_proveedores
        FOREIGN KEY (id_proveedor)
        REFERENCES compras.proveedores(id_proveedor)
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS compras.guias_entrada (
    id_guia SERIAL PRIMARY KEY,
    id_proveedor INT,
    estado VARCHAR(50),
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_guias_entrada
        FOREIGN KEY (id_guia)
        REFERENCES compras.guias_entrada(id_guia)
        ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS compras.proveedor_productos (
    id_proveedor_producto SERIAL PRIMARY KEY,
    id_proveedor INT,
    id_producto INT,
    precio NUMERIC(10,2),
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_proveedor_productos
        FOREIGN KEY (id_proveedor_producto)
        REFERENCES compras.proveedor_productos(id_proveedor_producto)
        ON DELETE RESTRICT

);