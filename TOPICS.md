Docker as:

- tool for runtime version control
- tool for launching applications without installation
- tool for packaging application
- tool for artifacts delivery

Docker components:

- docker demon
- docker CLI
- docker images
- docker containers
- docker register
- docker compose

Docker for Go apps:

https://hub.docker.com/_/golang
https://hub.docker.com/_/mysql

1. server with database

- bundle server app
- run database in container
- run app and database as compose services

2. different runtimes - same code

- run code with last version of go runtime
- run same code with older go

docker build -t warehouse .

docker run -it --rm --name warehouse-test -p 8080:8080 --network shared warehouse

docker run --rm --name warehouse \
 -e MYSQL_ROOT_PASSWORD=root \
 -e MYSQL_DATABASE=warehouse \
 -v ./warehouse/db/init.sql:/docker-entrypoint-initdb.d/init.sql \
 --network shared \
 --network-alias mysql \
 -p 8083:3306 \
 mysql:8
