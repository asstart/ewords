# Endlish-Words app

### Compile and run

```shell

cd app
go build -o ../ewords
./ewords
```

### Run integration test with env

```shell
NOTION_API_KEY="API_KEY" go test -v -tags=integration
```

### Run with notion

```shell
NOTION_API_KEY="API_KEY" ewords -n <db_id> -m
```