# Movies API

## Start - Debug

```
go run main.go
```

## Docker

### Build

```
docker build -t movies-api:latest .
```

### Run

```
docker run -d -p 8080:8080 --name movies-api -it movies-api:latest
```

### Logs

```
docker logs --follow --tail all movies-api
```