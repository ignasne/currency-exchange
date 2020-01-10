## Back-end API

Api written using Go programming language. How to setup Go environment
[https://golang.org/doc/install](https://golang.org/doc/install) 

Api converts given currency amount to needed currency by fetching these
two currencies exchange ratio from third party api.

Api uses database for cache purpose, please not forget to run database migrations.

Available Swagger documentation [swagger.yaml](.doc/swagger.yaml)
 
### Running API
Using [Makefile](https://en.wikipedia.org/wiki/Makefile)

### For development
Use [direnv](https://direnv.net/) for development variables set. 
Copy ```.envrc.dist``` to ```.envrc```

1. Add env variables:
    ```bash
    direnv allow
    ```
1. Update vendors:
    ```bash
    make vendor
    ```
1. Run migrations:
    ```bash
    make migrate
    ```
1. Run api:
    ```bash
    make run-dev
    ```

Api will be available at ```http://localhost:8080```

### Using Docker
```bash
make build && make start
```
Api docker containers will be created and running. There are
database migrations which should be executed manually.
Locate api docker container while executing:
```bash
docker ps
```
Search for `quotes-api` image and copy `CONTAINER ID`

Login to docker container:
```bash
docker exec -it CONTAINER_ID bin/sh
```

Run the following command inside `quotes-api` container:
```bash
quotes-api migrate
```

Successful result should output message similar to this one:
```
{"level":"info","msg":"migrations applied","time":"2020-01-08T23:50:56Z","total_migrations":2}
```

For stopping Docker containers:
```bash
make stop
```

Api will be available at ```http://localhost:8080``` (the same port as for development environment)