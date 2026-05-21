
CREATE SCHEMA IF NOT EXISTS inventario;



CREATE TABLE inventario.categoria_producto (
    id_categoria_producto SERIAL PRIMARY KEY,

    nombre VARCHAR(50),

    activo BOOLEAN DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE inventario.inventario (
    id_inventario SERIAL PRIMARY KEY,

    id_empresa INT REFERENCES empresa.empresas(id_empresa),

    id_bodega INT REFERENCES empresa.bodegas(id_bodegas),

    nombre VARCHAR(100),

    fecha DATE,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);



CREATE TABLE inventario.productos (
    id_producto SERIAL PRIMARY KEY,

    id_empresa INT REFERENCES empresa.empresas(id_empresa),

    id_bodega INT REFERENCES empresa.bodegas(id_bodegas),

    id_categoria_producto INT REFERENCES inventario.categoria_producto(id_categoria_producto),

    nombre VARCHAR(100),

    sku VARCHAR(50) UNIQUE,

    precio_unitario DECIMAL(10,2),

    stock_inicial INT DEFAULT 0,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_precio_producto
        CHECK (precio_unitario >= 0),

    CONSTRAINT chk_stock_producto
        CHECK (stock_inicial >= 0)
);


CREATE TABLE inventario.inventario_productos (
    id SERIAL PRIMARY KEY,

    id_inventario INT REFERENCES inventario.inventario(id_inventario),

    id_producto INT REFERENCES inventario.productos(id_producto),

    cantidad INT DEFAULT 0,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_cantidad_inventario
        CHECK (cantidad >= 0)
);


CREATE TABLE inventario.tipos_movimiento (
    id_tipo_movimiento SERIAL PRIMARY KEY,

    nombre VARCHAR(50) UNIQUE,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE inventario.movimientos (
    id_movimiento SERIAL PRIMARY KEY,

    id_tipo_movimiento INT REFERENCES inventario.tipos_movimiento(id_tipo_movimiento),

    id_producto INT REFERENCES inventario.productos(id_producto),

    id_bodegas INT REFERENCES empresa.bodegas(id_bodegas),

    cantidad INT,

    motivo VARCHAR(100),

    fecha TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_cantidad_movimiento
        CHECK (cantidad > 0)
);



CREATE TABLE inventario.ajustes (
    id_ajuste SERIAL PRIMARY KEY,

    id_producto INT REFERENCES inventario.productos(id_producto),

    cantidad_esperada INT,

    cantidad_real INT,

    diferencia INT GENERATED ALWAYS AS (
        cantidad_real - cantidad_esperada
    ) STORED,

    motivo VARCHAR(100),

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE inventario.transferencia (
    id_transferencia SERIAL PRIMARY KEY,

    id_producto INT REFERENCES inventario.productos(id_producto),

    cantidad INT,

    activo BOOLEAN NOT NULL DEFAULT true,

    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT chk_cantidad_transferencia
        CHECK (cantidad > 0)
);
