start server:

first method:

```
./build.sh
```

second method:

```
cd cmd
go run main.go
```

third method

```
go build -o ../bin/main
```

create .env file at project's root

variables:

```
DB_ADDRESS="postgres://username:password@host:port/db_name"
```