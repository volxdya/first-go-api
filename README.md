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
go build -o ../bin/main.exe
```

create .env file at project's root

variables:

*** - non required variable

```
DB_ADDRESS="postgres://username:password@host:port/db_name"
SERVER_PORT="8080"
*** SERVER_HOST="localhost" # on current moment using only for logs
```