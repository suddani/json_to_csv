# JSON To Csv

Converts a file containing json objects to a csv


## Install

```
go get github.com/suddani/json_to_csv/cmd/json_to_csv
```

## Uninstall

```
go clean -i -n github.com/suddani/json_to_csv/...
```

## Usage

```
NAME:
   json_to_csv - Converts a file containing json objects to a csv

USAGE:
   json_to_csv [global options] [command] FILE

VERSION:
   v0.0.4

DESCRIPTION:
   Convert a stream of json objects to csv
   If no file is given stdin is used

COMMANDS:
   filter     only filters the original file and does not convert to csv
   keys-only  only print keys
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --filter FILTER, -f FILTER  FILTER by key,value0,value1,value2
   --filter-file FILE          Load filter from a FILE, per line instead of commad seperated
   --key-file FILE             Load keys from a FILE, one per line
   --keys KEYS, -k KEYS        KEYS to use for csv. Comma sperated
   --limit LIMIT, -l LIMIT     Print only LIMIT number of rows (default: 0)
   --no-header                 Print no header line (default: false)
   --output FILE, -o FILE      Sets the output FILE (default: "-")
   --regex-filter              Treat filter as regex (default: false)
   --stdout                    Print to stdout as well as to output (default: false)
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```

## Example

### Simle conversion

```bash
printf "{\"name\":\"user\",\"id\":1}\n{\"name\":\"other\",\"id\":2}" | json_to_csv
```

### Keys only

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv keys-only
```

### Filter json without conversion

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -f country,us filter
```

### Filter with conversion

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"it\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -f country,us.it
```

### Filter with conversion from filter file

```bash
printf "country\nus\nit" > simple_filter && \
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"it\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv --filter-file simple_filter
```

### Only print certain keys

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -k id,name
```

### Only print certain keys from file

```bash
printf "id\nname" > simple_keys && \
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv --key-file simple_keys
```

### Skip header line

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -f country,us --no-header
```

### Limit output

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -l 1
```

### Output to a file

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -l 1 -o somefile.csv
```

### Count values

```bash
printf "{\"name\":\"user\",\"id\":1,\"country\":\"us\"}\n{\"name\":\"other\",\"id\":2,\"country\":\"de\"}\n{\"name\":\"other2\",\"id\":3,\"country\":\"de\"}" | json_to_csv -k country --no-header|sort|uniq -c
```