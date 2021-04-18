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
   v1.0.0

DESCRIPTION:
   Convert a stream of json objects to csv
   If no file is given stdin is used

COMMANDS:
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
   --stdout                    Print to stdout as well as to output (default: false)
   --help, -h                  show help (default: false)
   --version, -v               print the version (default: false)
```