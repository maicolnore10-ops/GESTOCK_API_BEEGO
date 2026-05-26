-- ============================================================================
-- API CONFIGURACIÓN - Script de Datos Iniciales (SEED)
-- ============================================================================
-- Ejecutar este script después de crear las tablas
-- Inserta datos de ejemplo para pruebas
-- ============================================================================

SET search_path TO configuracion, public;

-- ============================================================================
-- INSERTAR CONFIGURACIÓN DEL SISTEMA
-- ============================================================================

INSERT INTO configuracion.sistema (nombre_empresa, descripcion, url, telefono, correo, logo_url, activo)
VALUES 
    ('Tecnología Avanzada S.A.', 'Empresa líder en soluciones tecnológicas', 'https://tecavanzada.com', '3105551234', 'info@tecavanzada.com', 'https://tecavanzada.com/logo.png', true),
    ('Desarrollo Web Pro', 'Especialistas en desarrollo web y móvil', 'https://devweb.pro', '3129876543', 'contacto@devweb.pro', 'https://devweb.pro/logo.png', true)
ON CONFLICT (nombre_empresa) DO NOTHING;

-- ============================================================================
-- INSERTAR CONFIGURACIONES DE SEGURIDAD
-- ============================================================================

INSERT INTO configuracion.seguridad (tipo_autenticacion, requiere_autenticacion, intentos_maximos_fallidos, tiempo_bloqueo_minutos, requiere_two_factor, encriptacion_datos, sesion_maxima_horas, descripcion, activo)
VALUES 
    ('OAuth2', true, 5, 15, true, true, 8, 'Configuración de seguridad estricta con OAuth2 y 2FA', true),
    ('JWT', true, 3, 20, false, true, 4, 'Configuración JWT con sesiones cortas', true),
    ('LDAP', true, 5, 30, true, true, 12, 'Integración con LDAP corporativo', false),
    ('MultiAuthenticación', true, 4, 15, true, true, 6, 'Soporte para múltiples métodos de autenticación', true)
ON CONFLICT DO NOTHING;

-- ============================================================================
-- INSERTAR CONFIGURACIONES DE BACKUP
-- ============================================================================

INSERT INTO configuracion.backup (nombre_backup, tipo_backup, fecha_ejecucion, fecha_proximo_backup, frecuencia, ubicacion_guardado, tamano_mb, estado, descripcion, activo)
VALUES 
    ('Backup Completo Diario', 'Completo', '2026-05-20 02:00:00', '2026-05-21 02:00:00', 'Diaria', '/backups/completo/bd_20260520.bak', 2048.50, 'Exitoso', 'Backup automático diario de toda la BD', true),
    ('Backup Incremental Diario', 'Incremental', '2026-05-20 06:00:00', '2026-05-21 06:00:00', 'Diaria', '/backups/incremental/bd_20260520_inc.bak', 512.75, 'Exitoso', 'Backup incremental después del completo', true),
    ('Backup Semanal Completo', 'Completo', '2026-05-18 04:00:00', '2026-05-25 04:00:00', 'Semanal', '/backups/semanal/bd_20260518.bak', 2150.25, 'Exitoso', 'Backup semanal de la BD', true),
    ('Backup de Configuración', 'Completo', '2026-05-19 23:00:00', '2026-05-20 23:00:00', 'Diaria', '/backups/configuracion/config_20260519.bak', 125.00, 'Exitoso', 'Backup de archivos de configuración', true),
    ('Backup en la Nube', 'Incremental', '2026-05-20 03:30:00', '2026-05-21 03:30:00', 'Diaria', 'cloud://aws-s3/backups/bd_incremental', 450.00, 'Exitoso', 'Backup incrementado enviado a AWS S3', true),
    ('Backup Mensual', 'Completo', '2026-05-01 02:00:00', '2026-06-01 02:00:00', 'Mensual', '/backups/mensual/bd_202605.bak', 2200.00, 'Exitoso', 'Backup completo mensual archivado', true),
    ('Backup de Prueba', 'Diferencial', '2026-05-20 18:00:00', '2026-05-21 18:00:00', 'Bajo demanda', '/backups/prueba/db_test_20260520.bak', 856.30, 'Exitoso', 'Backup de prueba para validación', true)
ON CONFLICT DO NOTHING;

-- ============================================================================
-- INSERTAR NOTIFICACIONES
-- ============================================================================

INSERT INTO configuracion.notificacion (titulo, descripcion, tipo, canal, destinatario, fecha_envio, fecha_programada, estado, intentos_envio, prioridad, activo)
VALUES 
    ('Backup Completado', 'El backup diario se completó exitosamente a las 02:00', 'Informativa', 'Email', 'admin@tecavanzada.com', '2026-05-20 02:15:00', '2026-05-20 02:00:00', 'Enviada', 1, 'Media', true),
    ('Seguridad: Nueva Contraseña Requerida', 'Por política de seguridad, debe cambiar su contraseña', 'Alerta', 'Email', 'usuarios@tecavanzada.com', NULL, '2026-05-21 08:00:00', 'Pendiente', 0, 'Alta', true),
    ('Actualización del Sistema Disponible', 'Una nueva versión del sistema está disponible', 'Informativa', 'Notificación Push', 'todos@tecavanzada.com', '2026-05-19 15:30:00', '2026-05-19 15:00:00', 'Enviada', 1, 'Media', true),
    ('Error en Copia de Seguridad', 'El backup incremental falló. Revise los logs', 'Error', 'Email', 'ops@tecavanzada.com', NULL, '2026-05-20 09:00:00', 'Pendiente', 2, 'Urgente', true),
    ('Mantenimiento Programado', 'Se realizará mantenimiento del servidor el domingo', 'Alerta', 'SMS', 'admin@tecavanzada.com', NULL, '2026-05-22 10:00:00', 'Pendiente', 0, 'Alta', true),
    ('Bienvenida al Sistema', 'Bienvenido al sistema de configuración', 'Informativa', 'Email', 'nuevo@tecavanzada.com', '2026-05-20 10:00:00', '2026-05-20 10:00:00', 'Enviada', 1, 'Baja', true),
    ('Límite de Licencia Próximo', 'El período de licencia vence en 30 días', 'Alerta', 'Email', 'licencias@tecavanzada.com', NULL, '2026-06-19 09:00:00', 'Pendiente', 0, 'Media', true),
    ('Revisión de Logs de Seguridad', 'Se detectaron 5 intentos fallidos de acceso', 'Alerta', 'Email', 'seguridad@tecavanzada.com', '2026-05-20 08:30:00', '2026-05-20 08:30:00', 'Enviada', 1, 'Alta', true),
    ('Notificación de Prueba', 'Esta es una notificación de prueba del sistema', 'Informativa', 'Email', 'test@tecavanzada.com', '2026-05-20 11:00:00', '2026-05-20 11:00:00', 'Enviada', 1, 'Baja', true),
    ('Reporte Mensual Disponible', 'El reporte de actividades de mayo está disponible', 'Informativa', 'Email', 'reportes@tecavanzada.com', '2026-05-20 09:00:00', '2026-05-20 09:00:00', 'Enviada', 1, 'Media', true)
ON CONFLICT DO NOTHING;

-- ============================================================================
-- ACTUALIZAR SECUENCIAS
-- ============================================================================

-- Actualizar secuencia de sistema
SELECT setval('configuracion.sistema_id_sistema_seq', (SELECT MAX(id_sistema) FROM configuracion.sistema));

-- Actualizar secuencia de seguridad
SELECT setval('configuracion.seguridad_id_seguridad_seq', (SELECT MAX(id_seguridad) FROM configuracion.seguridad));

-- Actualizar secuencia de backup
SELECT setval('configuracion.backup_id_backup_seq', (SELECT MAX(id_backup) FROM configuracion.backup));

-- Actualizar secuencia de notificación
SELECT setval('configuracion.notificacion_id_notificacion_seq', (SELECT MAX(id_notificacion) FROM configuracion.notificacion));

-- ============================================================================
-- ESTADÍSTICAS
-- ============================================================================

SELECT 'Configuraciones del sistema insertadas' as descripcion, COUNT(*) as cantidad FROM configuracion.sistema UNION ALL
SELECT 'Políticas de seguridad insertadas', COUNT(*) FROM configuracion.seguridad UNION ALL
SELECT 'Configuraciones de backup insertadas', COUNT(*) FROM configuracion.backup UNION ALL
SELECT 'Notificaciones insertadas', COUNT(*) FROM configuracion.notificacion;

-- ============================================================================
-- RESUMEN DE DATOS
-- ============================================================================

SELECT 
    (SELECT nombre_empresa FROM configuracion.sistema WHERE activo = true LIMIT 1) as empresa_principal,
    (SELECT COUNT(*) FROM configuracion.seguridad WHERE activo = true) as politicas_activas,
    (SELECT COUNT(*) FROM configuracion.backup WHERE estado = 'Exitoso') as backups_exitosos,
    (SELECT COUNT(*) FROM configuracion.notificacion WHERE estado = 'Pendiente') as notificaciones_pendientes;

-- ============================================================================
-- FIN DEL SCRIPT
-- ============================================================================
-- Script de datos iniciales para API Configuración
-- Fecha: Mayo 2026
