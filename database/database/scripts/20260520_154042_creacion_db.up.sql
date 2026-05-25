CREATE SCHEMA IF NOT EXISTS autenticacion;

CREATE TABLE autenticacion.roles (
    id_rol               SERIAL          PRIMARY KEY,
    nombre_rol           VARCHAR(100)    NOT NULL UNIQUE,
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE autenticacion.usuarios (
    id_usuario           SERIAL          PRIMARY KEY,
    id_rol               INTEGER         NOT NULL,
    correo               VARCHAR(150)    NOT NULL UNIQUE,
    nombres              VARCHAR(100)    NULL,
    apellidos            VARCHAR(100)    NULL,
    telefono             VARCHAR(20)     NULL,
    fecha_nacimiento     DATE            NULL,
    documento            INTEGER         NULL,
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_usuario_rol
        FOREIGN KEY (id_rol)
        REFERENCES autenticacion.roles(id_rol)
        ON DELETE RESTRICT
);

CREATE TABLE autenticacion.contrasena_hash (
    id_contrasena_hash   SERIAL          PRIMARY KEY,
    id_usuario           INTEGER         NOT NULL,
    contrasena_hash      VARCHAR(255)    NOT NULL,
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_contrasena_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES autenticacion.usuarios(id_usuario)
        ON DELETE CASCADE
);

CREATE TABLE autenticacion.recuperacion_contrasena (
    id_recuperacion      SERIAL          PRIMARY KEY,
    id_usuario           INTEGER         NOT NULL,
    codigo_verificacion  VARCHAR(100)    NULL,
    fecha_expiracion     TIMESTAMP       NULL,
    usado                BOOLEAN         NULL DEFAULT FALSE,
    activo               BOOLEAN         NOT NULL DEFAULT TRUE,
    fecha_creacion       TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,
    fecha_actualizacion  TIMESTAMP       DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_recuperacion_usuario
        FOREIGN KEY (id_usuario)
        REFERENCES autenticacion.usuarios(id_usuario)
        ON DELETE CASCADE
);
