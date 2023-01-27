
# Questions API
Questions API was created to satisfy the requirements of the homework task provided by Toggle Hire. Api follows Clean Architecture principles and the communication layer is implemented in classic JSON (http) and gRPC

# Configuration
Specific configuration options can be provided in .yaml file or by environmental variables. Example configuration can be found in file `config_dev.yaml`


|   Name             |yaml     |env         | Desciption |
|----------------|---------|------------| ------| 
|`host`|`api.host`     |`API_HOST`   | Host for api. Default: `0.0.0.0`
|`port`          |`api.port`            |`PORT`            | Port for api. Default: `3000`
|`mode`          |`api.mode`|`API_MODE`| Comunication layer. `http` or `grpc`
|`jwtSecret`          |`auth.jwtSecret`|`AUTH_JWTSECRET`| Secret for JWT token validation
|`dbFile`          |`db.file`|`DB_FILE`| Path to sqlite .db file. File is created if doesn't exists 
|`dbMigrationsPath`          |`db.migrationsPath`|`DB_MIGRATIONSPATH`| Path to folder with migration files

# API endpoints documentation
API documentation generated from OpenAPI definition
https://pretty-lizard-62.redoc.ly/

# Testing
API contains 2 use case tests. It's a very simple crud so there is not much to test (a small amount of logic). My implementation is overcomplicated for such a simple API but I wanted to demonstrate how I would approach writing more complex APIs. Because of such a structure, it's easy to e.g. replace SQLite with Postgres. 

# SQLite3
API enabled `Foreign Key` support on every connection, so in order to properly delete data manually by `sqlite3 [db file name]` you have to execute `PRAGMA foreign_keys = ON;` before performing any deletion.

# How to start?

       git clone https://github.com/KiVo16/toggl_hw.git
       cd toggl_hw
       make docker-run-dev
       
  The above commands run Docker container which shares port `3000 ` with the host and API can be accessed on that port. Also `$(pwd)/docker/sqlite` folder is created and attached as volume to make the file available for convenient previewing. 