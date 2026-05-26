-- ============================================================================
-- API CONFIGURACIÓN - Script de Creación de Tablas
-- ============================================================================
-- Ejecutar este script en PostgreSQL después de crear la base de datos
-- ============================================================================

-- Crear el esquema si no existe
CREATE SCHEMA IF NOT EXISTS configuracion;

-- Establecer el search_path para usar el esquema por defecto
SET search_path TO configuracion, public;

-- ============================================================================
-- TABLA: sistema
-- Descripción: Configuración general del sistema
-- ============================================================================
CREATE TABLE IF NOT EXISTS configuracion.sistema (
    id_sistema SERIAL PRIMARY KEY,
    nombre_empresa VARCHAR(255) NOT NULL UNIQUE,
    descripcion TEXT,
    url VARCHAR(255),
    telefono VARCHAR(20),
    correo VARCHAR(255),
    logo_url VARCHAR(255),
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear índices para búsquedas rápidas
CREATE INDEX idx_sistema_nombre_empresa ON configuracion.sistema(nombre_empresa);
CREATE INDEX idx_sistema_activo ON configuracion.sistema(activo);

-- ============================================================================
-- TABLA: seguridad
-- Descripción: Configuración de políticas de seguridad
-- ============================================================================
CREATE TABLE IF NOT EXISTS configuracion.seguridad (
    id_seguridad SERIAL PRIMARY KEY,
    tipo_autenticacion VARCHAR(100) NOT NULL,
    requiere_autenticacion BOOLEAN DEFAULT true,
    intentos_maximos_fallidos INTEGER DEFAULT 3,
    tiempo_bloqueo_minutos INTEGER DEFAULT 15,
    requiere_two_factor BOOLEAN DEFAULT false,
    encriptacion_datos BOOLEAN DEFAULT true,
    sesion_maxima_horas INTEGER DEFAULT 8,
    descripcion TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear índices para búsquedas rápidas
CREATE INDEX idx_seguridad_tipo_autenticacion ON configuracion.seguridad(tipo_autenticacion);
CREATE INDEX idx_seguridad_activo ON configuracion.seguridad(activo);

-- ============================================================================
-- TABLA: backup
-- Descripción: Configuración de backups del sistema
-- ============================================================================
CREATE TABLE IF NOT EXISTS configuracion.backup (
    id_backup SERIAL PRIMARY KEY,
    nombre_backup VARCHAR(255) NOT NULL,
    tipo_backup VARCHAR(100) NOT NULL,
    fecha_ejecucion TIMESTAMP NOT NULL,
    fecha_proximo_backup TIMESTAMP,
    frecuencia VARCHAR(100),
    ubicacion_guardado VARCHAR(500),
    tamano_mb DECIMAL(10,2),
    estado VARCHAR(50) NOT NULL,
    descripcion TEXT,
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear índices para búsquedas rápidas
CREATE INDEX idx_backup_nombre_backup ON configuracion.backup(nombre_backup);
CREATE INDEX idx_backup_tipo_backup ON configuracion.backup(tipo_backup);
CREATE INDEX idx_backup_fecha_ejecucion ON configuracion.backup(fecha_ejecucion);
CREATE INDEX idx_backup_estado ON configuracion.backup(estado);
CREATE INDEX idx_backup_activo ON configuracion.backup(activo);

-- ============================================================================
-- TABLA: notificacion
-- Descripción: Sistema de notificaciones del sistema
-- ============================================================================
CREATE TABLE IF NOT EXISTS configuracion.notificacion (
    id_notificacion SERIAL PRIMARY KEY,
    titulo VARCHAR(255) NOT NULL,
    descripcion TEXT NOT NULL,
    tipo VARCHAR(100) NOT NULL,
    canal VARCHAR(100) NOT NULL,
    destinatario VARCHAR(255),
    fecha_envio TIMESTAMP,
    fecha_programada TIMESTAMP,
    estado VARCHAR(50) NOT NULL DEFAULT 'Pendiente',
    intentos_envio INTEGER DEFAULT 0,
    mensaje_error TEXT,
    prioridad VARCHAR(50),
    activo BOOLEAN DEFAULT true,
    fecha_creacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    fecha_modificacion TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Crear índices para búsquedas rápidas
CREATE INDEX idx_notificacion_tipo ON configuracion.notificacion(tipo);
CREATE INDEX idx_notificacion_canal ON configuracion.notificacion(canal);
CREATE INDEX idx_notificacion_estado ON configuracion.notificacion(estado);
CREATE INDEX idx_notificacion_fecha_envio ON configuracion.notificacion(fecha_envio);
CREATE INDEX idx_notificacion_prioridad ON configuracion.notificacion(prioridad);
CREATE INDEX idx_notificacion_activo ON configuracion.notificacion(activo);

-- ============================================================================
-- COMENTARIOS DE TABLAS
-- ============================================================================
COMMENT ON TABLE configuracion.sistema IS 'Tabla que almacena la configuración general del sistema';
COMMENT ON TABLE configuracion.seguridad IS 'Tabla que almacena las políticas de seguridad';
COMMENT ON TABLE configuracion.backup IS 'Tabla que registra la configuración y ejecución de backups';
COMMENT ON TABLE configuracion.notificacion IS 'Tabla que gestiona las notificaciones del sistema';

-- ============================================================================
-- COMENTARIOS DE COLUMNAS
-- ============================================================================

-- Sistema
COMMENT ON COLUMN configuracion.sistema.id_sistema IS 'Identificador único de la configuración';
COMMENT ON COLUMN configuracion.sistema.nombre_empresa IS 'Nombre comercial de la empresa';
COMMENT ON COLUMN configuracion.sistema.descripcion IS 'Descripción general de la empresa';
COMMENT ON COLUMN configuracion.sistema.url IS 'Sitio web o URL de la empresa';
COMMENT ON COLUMN configuracion.sistema.telefono IS 'Teléfono de contacto principal';
COMMENT ON COLUMN configuracion.sistema.correo IS 'Correo electrónico de contacto';
COMMENT ON COLUMN configuracion.sistema.logo_url IS 'URL del logo de la empresa';
COMMENT ON COLUMN configuracion.sistema.activo IS 'Indica si la configuración está activa';

-- Seguridad
COMMENT ON COLUMN configuracion.seguridad.id_seguridad IS 'Identificador único de la política';
COMMENT ON COLUMN configuracion.seguridad.tipo_autenticacion IS 'Tipo de autenticación (OAuth2, JWT, etc)';
COMMENT ON COLUMN configuracion.seguridad.requiere_autenticacion IS 'Si se requiere autenticación obligatoria';
COMMENT ON COLUMN configuracion.seguridad.intentos_maximos_fallidos IS 'Número máximo de intentos fallidos';
COMMENT ON COLUMN configuracion.seguridad.tiempo_bloqueo_minutos IS 'Minutos de bloqueo después de fallos';
COMMENT ON COLUMN configuracion.seguridad.requiere_two_factor IS 'Si se requiere autenticación de dos factores';
COMMENT ON COLUMN configuracion.seguridad.encriptacion_datos IS 'Si los datos están encriptados';
COMMENT ON COLUMN configuracion.seguridad.sesion_maxima_horas IS 'Duración máxima de sesión en horas';

-- Backup
COMMENT ON COLUMN configuracion.backup.id_backup IS 'Identificador único del backup';
COMMENT ON COLUMN configuracion.backup.nombre_backup IS 'Nombre descriptivo del backup';
COMMENT ON COLUMN configuracion.backup.tipo_backup IS 'Tipo (Completo, Incremental, Diferencial)';
COMMENT ON COLUMN configuracion.backup.fecha_ejecucion IS 'Cuándo se ejecutó o ejecutará';
COMMENT ON COLUMN configuracion.backup.fecha_proximo_backup IS 'Fecha del próximo backup programado';
COMMENT ON COLUMN configuracion.backup.frecuencia IS 'Frecuencia de ejecución (Diaria, Semanal, etc)';
COMMENT ON COLUMN configuracion.backup.ubicacion_guardado IS 'Ruta donde se guarda el backup';
COMMENT ON COLUMN configuracion.backup.tamano_mb IS 'Tamaño en MB del archivo de backup';
COMMENT ON COLUMN configuracion.backup.estado IS 'Estado (Exitoso, Fallido, En progreso, Pendiente)';

-- Notificación
COMMENT ON COLUMN configuracion.notificacion.id_notificacion IS 'Identificador único de la notificación';
COMMENT ON COLUMN configuracion.notificacion.titulo IS 'Título o asunto de la notificación';
COMMENT ON COLUMN configuracion.notificacion.descripcion IS 'Contenido detallado de la notificación';
COMMENT ON COLUMN configuracion.notificacion.tipo IS 'Tipo (Informativa, Alerta, Urgente, Error)';
COMMENT ON COLUMN configuracion.notificacion.canal IS 'Canal de envío (Email, SMS, Push, etc)';
COMMENT ON COLUMN configuracion.notificacion.destinatario IS 'A quién se envía la notificación';
COMMENT ON COLUMN configuracion.notificacion.fecha_envio IS 'Cuándo se envió la notificación';
COMMENT ON COLUMN configuracion.notificacion.fecha_programada IS 'Cuándo está programada para enviar';
COMMENT ON COLUMN configuracion.notificacion.estado IS 'Estado (Pendiente, Enviada, Fallida, etc)';
COMMENT ON COLUMN configuracion.notificacion.intentos_envio IS 'Número de intentos de envío';
COMMENT ON COLUMN configuracion.notificacion.prioridad IS 'Nivel de prioridad (Baja, Media, Alta)';

-- ============================================================================
-- FUNCIONES AUXILIARES
-- ============================================================================

-- Función para actualizar la fecha de modificación automáticamente
CREATE OR REPLACE FUNCTION configuracion.update_fecha_modificacion()
RETURNS TRIGGER AS $$
BEGIN
    NEW.fecha_modificacion = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- ============================================================================
-- TRIGGERS
-- ============================================================================

-- Trigger para actualizar fecha_modificacion en sistema
CREATE TRIGGER trigger_sistema_update
BEFORE UPDATE ON configuracion.sistema
FOR EACH ROW
EXECUTE FUNCTION configuracion.update_fecha_modificacion();

-- Trigger para actualizar fecha_modificacion en seguridad
CREATE TRIGGER trigger_seguridad_update
BEFORE UPDATE ON configuracion.seguridad
FOR EACH ROW
EXECUTE FUNCTION configuracion.update_fecha_modificacion();

-- Trigger para actualizar fecha_modificacion en backup
CREATE TRIGGER trigger_backup_update
BEFORE UPDATE ON configuracion.backup
FOR EACH ROW
EXECUTE FUNCTION configuracion.update_fecha_modificacion();

-- Trigger para actualizar fecha_modificacion en notificacion
CREATE TRIGGER trigger_notificacion_update
BEFORE UPDATE ON configuracion.notificacion
FOR EACH ROW
EXECUTE FUNCTION configuracion.update_fecha_modificacion();

-- ============================================================================
-- VISTAS ÚTILES
-- ============================================================================

-- Vista: Resumen de configuración del sistema
CREATE OR REPLACE VIEW configuracion.v_resumen_sistema AS
SELECT 
    s.id_sistema,
    s.nombre_empresa,
    s.url,
    s.correo,
    s.activo,
    COUNT(DISTINCT seg.id_seguridad) as politicas_seguridad,
    COUNT(DISTINCT b.id_backup) as backups_configurados,
    COUNT(DISTINCT n.id_notificacion) as notificaciones_totales
FROM configuracion.sistema s
LEFT JOIN configuracion.seguridad seg ON true
LEFT JOIN configuracion.backup b ON true
LEFT JOIN configuracion.notificacion n ON true
GROUP BY s.id_sistema, s.nombre_empresa, s.url, s.correo, s.activo;

-- Vista: Notificaciones pendientes
CREATE OR REPLACE VIEW configuracion.v_notificaciones_pendientes AS
SELECT 
    id_notificacion,
    titulo,
    tipo,
    canal,
    destinatario,
    fecha_programada,
    prioridad,
    estado
FROM configuracion.notificacion
WHERE estado = 'Pendiente' OR estado IS NULL
ORDER BY 
    CASE WHEN prioridad = 'Alta' THEN 1 
         WHEN prioridad = 'Media' THEN 2 
         ELSE 3 END,
    fecha_programada ASC;

-- Vista: Historial de backups recientes
CREATE OR REPLACE VIEW configuracion.v_backups_recientes AS
SELECT 
    id_backup,
    nombre_backup,
    tipo_backup,
    fecha_ejecucion,
    fecha_proximo_backup,
    tamano_mb,
    estado,
    frecuencia
FROM configuracion.backup
ORDER BY fecha_ejecucion DESC
LIMIT 10;

-- ============================================================================
-- DATOS INICIALES (EJEMPLO)
-- ============================================================================

-- Insertar configuración inicial del sistema
INSERT INTO configuracion.sistema (nombre_empresa, descripcion, url, telefono, correo, activo)
VALUES (
    'Mi Empresa S.A.',
    'Descripción de la empresa',
    'https://miempresa.com',
    '3123456789',
    'contacto@miempresa.com',
    true
) ON CONFLICT (nombre_empresa) DO NOTHING;

-- Insertar configuración de seguridad por defecto
INSERT INTO configuracion.seguridad (tipo_autenticacion, requiere_autenticacion, activo)
VALUES (
    'OAuth2',
    true,
    true
) ON CONFLICT DO NOTHING;

-- ============================================================================
-- FIN DEL SCRIPT
-- ============================================================================
-- Script creado para API Configuración
-- Fecha: Mayo 2026
-- Framework: Beego 2.3.10
-- Database: PostgreSQL 12+
