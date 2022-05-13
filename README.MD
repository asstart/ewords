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
  -ed
        export defenitions
  -ee
        export examples
  -f    export to a file
  -sd string
        path to a source directory
  -sf string
        path to a source file
  -sn string
        id of notion database
  -tdd string
        where store defenition output (default "ewords_defenition")
  -ted string
        where store example output (default "ewords_example")
```

### Export examples and defenitions from notions to a file

```shell
NOTION_API_KEY="API_KEY" ewords -n <db_id> -f -ted -tdd
```

### Run tests from all packages

```shell
go test ./...
```

### Run tests from all packages by tag with set env var

```shell
 NOTION_API_KEY=<API_KEY>  go test ./... --tags=integration
```