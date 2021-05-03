package main

import (
	"io"
	"log"
	"os"
	"sort"
	"strings"

	json_to_csv "github.com/suddani/json_to_csv/internal"
	"github.com/urfave/cli/v2"
)

func outputStream(file string, stdout bool) (io.Writer, error) {
	if file == "" || file == "-" {
		return os.Stdout, nil
	}
	writer, err := os.Create(file)
	if err != nil {
		return nil, err
	}
	if !stdout {
		return writer, nil
	}

	return io.MultiWriter(writer, os.Stdout), nil
}

func inputStream(inputFile string) (io.Reader, error) {
	if inputFile == "" || inputFile == "-" {
		return os.Stdin, nil
	}

	reader, err := os.Open(inputFile)
	if err != nil {
		return nil, err
	}
	return reader, nil
}

func SliceString(s string) []string {
	if s == "" {
		return nil
	}

	array := strings.Split(s, ",")
	if len(array) == 0 {
		return nil
	}

	return array
}

func main() {
	app := &cli.App{
		Name:                 "json_to_csv",
		Description:          "Convert a stream of json objects to csv\nIf no file is given stdin is used",
		EnableBashCompletion: true,
		Version:              "v0.0.2",
		Usage:                "Converts a file containing json objects to a csv",
		UsageText:            "json_to_csv [global options] [command] FILE",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Value:   "-",
				Usage:   "Sets the output `FILE`",
			},
			&cli.BoolFlag{
				Name:  "stdout",
				Value: false,
				Usage: "Print to stdout as well as to output",
			},
			&cli.BoolFlag{
				Name:  "no-header",
				Value: false,
				Usage: "Print no header line",
			},
			&cli.BoolFlag{
				Name:  "regex-filter",
				Value: false,
				Usage: "Treat filter as regex",
			},
			&cli.IntFlag{
				Name:    "limit",
				Aliases: []string{"l"},
				Value:   0,
				Usage:   "Print only `LIMIT` number of rows",
			},
			&cli.StringFlag{
				Name:    "keys",
				Aliases: []string{"k"},
				Value:   "",
				Usage:   "`KEYS` to use for csv. Comma sperated",
			},
			&cli.StringFlag{
				Name:  "key-file",
				Value: "",
				Usage: "Load keys from a `FILE`, one per line",
			},
			&cli.StringFlag{
				Name:    "filter",
				Aliases: []string{"f"},
				Value:   "",
				Usage:   "`FILTER` by key,value0,value1,value2",
			},
			&cli.StringFlag{
				Name:  "filter-file",
				Value: "",
				Usage: "Load filter from a `FILE`, per line instead of commad seperated",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "keys-only",
				Usage: "only print keys",
				Action: func(c *cli.Context) error {
					input, err := inputStream(c.Args().Get(0))
					if err != nil {
						return err
					}

					output, err := outputStream(c.String("output"), c.Bool("stdout"))
					if err != nil {
						return err
					}

					writer := json_to_csv.NewLimitCsvWriter(output, c.Int("limit"), false)
					err = json_to_csv.FindAllKeys(input, writer)
					writer.Flush()
					if err != nil && err.Error() != "maximum number of rows reached" {
						return err
					}
					return nil
				},
			},
		},
		Action: func(c *cli.Context) error {
			keys, err := json_to_csv.LoadKeys(SliceString(c.String("keys")), c.String("key-file"))
			if err != nil {
				return err
			}

			filterCreator := json_to_csv.NewArrayFilter
			if c.Bool("regex-filter") {
				filterCreator = json_to_csv.NewRegexFilter
			}

			filter := filterCreator(SliceString(c.String("filter")), c.String("filter-file"))

			input, err := inputStream(c.Args().Get(0))
			if err != nil {
				return err
			}

			output, err := outputStream(c.String("output"), c.Bool("stdout"))
			if err != nil {
				return err
			}

			writer := json_to_csv.NewLimitCsvWriter(output, c.Int("limit"), !c.Bool("no-header"))
			err = json_to_csv.ConvertToCsv(input, writer, keys, filter, !c.Bool("no-header"))
			writer.Flush()
			if err != nil && err.Error() != "maximum number of rows reached" {
				return err
			}
			return nil
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
