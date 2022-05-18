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

### Usage

```shell
Usage of ./ewords:
  -file string
        path to a source file
  -dir string
        path to a source directory
  -notion string
        id of source notion database
  -defenition
        do export defenitions
  -example
        do export examples
  -exdir string
        where store examples output (default "ewords_example")
  -defdir string
        where store defenitions output (default "ewords_defenition")
```

### Export examples and defenitions from notions to a file

```shell
NOTION_API_KEY="API_KEY" ./ewords -notion <db_id> -example -defenition -exdir <path> -defdir <path>
```

### Run tests from all packages

```shell
go test ./...
```

### Run tests from all packages by tag with set env var

```shell
 NOTION_API_KEY=<API_KEY>  go test ./... --tags=integration
```