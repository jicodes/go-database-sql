Connect the PostgreSQL with database/sql package in golang

Run the postgres db in the Docker

```sh
docker run --name pg-container -p 5432:5432 -e POSTGRES_PASSWORD=xxx -d postgres:16-alpine
```

Create the user and db with Docker exec

```sh
docker exec -it pg-container createdb -U postgres testdb
```

Connect to the PostgreSQL with Docker exec

```sh
docker exec -it pg-container psql -U postgres -d testdb
```

Install loadenv package

```sh
go get github.com/joho/godotenv
```

Install the pq driver for Go's database/sql package

```sh
go get github.com/lib/pq
```

The `database/sql` package provides a general interface around SQL (or SQL-like) databases, and `pq` driver implements this interface for PostgreSQL.
