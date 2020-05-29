# Movies API

## Start - Debug

```{bash}
go run main.go
```

## Docker

### Build

```{bash}
docker build -t movies-api:latest .
```

### Run

```{bash}
docker run -d -p 8080:8080 --name movies-api -it movies-api:latest
```

### Logs

```{bash}
docker logs --follow --tail all movies-api
```
