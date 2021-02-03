# consecutivos crud

API CRUD para la gestión de consecutivos definidos en diferentes contextos para las entidades de la universidad.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
# parametros de api
CONSECUTIVOS_CRUD_HTTP_PORT=[Puerto de exposición del API]
CONSECUTIVOS_CRUD_RUN_MODE=[Modo de ejecución del API]
# paramametros de bd
CONSECUTIVOS_CRUD_PGUSER=[Usuario de BD]
CONSECUTIVOS_CRUD_PGPASS=[Contraseña del usaurio de BD]
CONSECUTIVOS_CRUD_PGHOST=[URL, Dominio o EndPoint de la BD]
CONSECUTIVOS_CRUD_PGPORT=[Puerto de la BD]
CONSECUTIVOS_CRUD_PGDB=[Nombre de Base de Datos]
CONSECUTIVOS_CRUD_PGSCHEMA=[Nombre del Esquema de Base de Datos]
```
**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con CONSECUTIVOS_CRUD_...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/consecutivos_crud

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/consecutivos_crud

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
CONSECUTIVOS_CRUD_HTTP_PORT=8080 CONSECUTIVOS_CRUD_PGHOST=127.0.0.1:27017 CONSECUTIVOS_CRUD_SOME_VARIABLE=some_value bee run
```
### Ejecución Dockerfile
```shell
# Implementado para despliegue del Sistema de integración continua CI.
```

### Ejecución docker-compose
```shell
#1. Clonar el repositorio
git clone -b develop https://github.com/udistrital/consecutivos_crud

#2. Moverse a la carpeta del repositorio
cd consecutivos_crud

#3. Crear un fichero con el nombre **custom.env**
touch custom.env

#4. Crear la network **back_end** para los contenedores
docker network create back_end

#5. Ejecutar el compose del contenedor
docker-compose up --build

#6. Comprobar que los contenedores estén en ejecución
docker ps
```

### Ejecución Pruebas
Pruebas unitarias
```shell
# En Proceso
```

## Modelo de Datos
[Modelo de Datos consecutivos_crud](/sql/modelo_consecutivos_crud.png)


## Estado CI
| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/consecutivos_crud/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/consecutivos_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/consecutivos_crud/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/consecutivos_crud) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/consecutivos_crud/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/consecutivos_crud) |


## Licencia

This file is part of consecutivos_crud.

consecutivos_crud is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

consecutivos_crud is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with consecutivos_crud. If not, see https://www.gnu.org/licenses/.
