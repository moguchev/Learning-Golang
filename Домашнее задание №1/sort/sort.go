package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Opts -
type Opts struct {
	Reverse  bool
	Numeric  bool
	Unique   bool
	FoldCase bool
	Key      int
	Output   string
}

func sortStrings(table [][]string, options Opts) error {
	if options.Key < 0 {
		return errors.New("Column under zero")
	}

	if options.Numeric {
		if options.Reverse {
			sort.Slice(table, func(i, j int) bool {
				f1, err := strconv.ParseFloat(table[i][options.Key], 64)
				f2, err := strconv.ParseFloat(table[j][options.Key], 64)
				if err != nil {
					return false
				}
				return f1 > f2
			})
		} else {
			sort.Slice(table, func(i, j int) bool {
				f1, err := strconv.ParseFloat(table[i][options.Key], 64)
				f2, err := strconv.ParseFloat(table[j][options.Key], 64)
				if err != nil {
					return false
				}
				return f1 < f2
			})
		}
	} else {
		if options.FoldCase && options.Reverse {
			sort.Slice(table, func(i, j int) bool {
				return strings.ToLower(table[i][options.Key]) > strings.ToLower(table[j][options.Key])
			})
		} else if options.FoldCase && !options.Reverse {
			sort.Slice(table, func(i, j int) bool {
				return strings.ToLower(table[i][options.Key]) < strings.ToLower(table[j][options.Key])
			})
		} else if !options.FoldCase && options.Reverse {
			sort.Slice(table, func(i, j int) bool {
				return table[i][options.Key] > table[j][options.Key]
			})
		} else if !options.FoldCase && !options.Reverse {
			sort.Slice(table, func(i, j int) bool {
				return table[i][options.Key] < table[j][options.Key]
			})
		}
	}

	return nil
}

func printSlice(slice []string, out io.Writer) {
	fmt.Fprintln(out, strings.Trim(fmt.Sprint(slice), "[]"))
}

func writeStrings(table [][]string, options Opts, out io.Writer) error {
	if options.Key < 0 {
		return errors.New("Column under zero")
	}

	if options.Unique && options.FoldCase {
		for i := range table {
			if i > 0 && strings.ToLower(table[i][options.Key]) !=
				strings.ToLower(table[i-1][options.Key]) {
				printSlice(table[i], out)
			} else if i == 0 {
				printSlice(table[i], out)
			}
		}
	} else if options.Unique && !options.FoldCase {
		for i := range table {
			if i > 0 && table[i][options.Key] != table[i-1][options.Key] {
				printSlice(table[i], out)
			} else if i == 0 {
				printSlice(table[i], out)
			}
		}
	} else if !options.Unique {
		for i := range table {
			printSlice(table[i], out)
		}
	}
	return nil
}

func main() {
	fPtr := flag.Bool("f", false, "")
	uPtr := flag.Bool("u", false, "")
	rPtr := flag.Bool("r", false, "")
	nPtr := flag.Bool("n", false, "")
	oPtr := flag.String("o", "", "")
	kPtr := flag.Int("k", 0, "")

	flag.Parse()
	args := flag.Args()
	var options Opts = Opts{
		FoldCase: *fPtr,
		Unique:   *uPtr,
		Reverse:  *rPtr,
		Numeric:  *nPtr,
		Output:   *oPtr,
		Key:      *kPtr,
	}

	if len(args) == 0 {
		log.Fatal("No arguments")
	}

	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile(args[0])
	if errRead != nil {
		log.Fatal(errRead)
	}

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))
	for i, line := range lines {
		if options.Key > 0 {
			tableOfStings[i] = strings.Split(line, " ")
		} else {
			tableOfStings[i] = make([]string, 1)
			tableOfStings[i][0] = line
		}
	}

	// пользователь ведёт нумерацию с 1
	if options.Key > 0 {
		options.Key--
	}
	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		log.Fatal(errSort)
	}

	if *oPtr != "" {
		out, err := os.Create(*oPtr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()
		errWrite := writeStrings(tableOfStings, options, nil)
		if errWrite != nil {
			log.Fatal(errWrite)
		}
	} else {
		errWrite := writeStrings(tableOfStings, options, os.Stdout)
		if errWrite != nil {
			log.Fatal(errWrite)
		}
	}
}
