

CREATE TABLE consecutivos.consecutivo(
    id  serial not null,
    contexto_id integer not null,
    year numeric(4,0) not null,
    consecutivo integer not null,
    descripcion CHARACTER VARYING(250),
    activo boolean NOT NULL,
    fecha_creacion timestamp NOT NULL DEFAULT now(),
    fecha_modificacion timestamp NOT NULL DEFAULT now(),
    CONSTRAINT pk_consecutivo PRIMARY KEY (id)
);

-- ddl-end --
COMMENT ON TABLE consecutivos.consecutivo IS 'Tabla que almacena los consecutivos generados secuencialmente para cada contexto en un año especifico';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.id IS 'Identificador del consecutivo';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.contexto_id IS 'Campo obligatorio que referencia el esquema parametros. Hace referencia al contexto o modelo de negocio al cual aplica el consecutivo';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.year IS 'Campo que indica el año para el cual es vigente el consecutivo.';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.consecutivo IS 'Campo que indica el el numero de consecutivo para un contexto y año especifico.';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.descripcion IS 'Descripción del consecutivo. Campo opcional que puede ser usado para concatenar el formato del consecutivo';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.activo IS 'Campo para indicar el estado del registro';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.fecha_creacion IS 'Fecha de creación del registro';
-- ddl-end --
COMMENT ON COLUMN consecutivos.consecutivo.fecha_modificacion IS 'Fecha de la última modificación del registro';
-- ddl-end --
COMMENT ON CONSTRAINT pk_consecutivo ON consecutivos.consecutivo  IS 'Llave primaria de la tabla consecutivo';
-- ddl-end --


-- Permisos de usuario
GRANT USAGE ON SCHEMA consecutivos TO desarrollooas;
GRANT SELECT, INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA consecutivos TO desarrollooas;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA consecutivos TO desarrollooas;
