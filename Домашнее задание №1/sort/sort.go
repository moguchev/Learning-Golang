package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func sortStrings(table [][]string, isF bool, isR bool, isN bool, k int) error {
	if k < 0 {
		return errors.New("Column under zero")
	}

	if isN {
		if isR {
			sort.Slice(table, func(i, j int) bool {
				f1, err := strconv.ParseFloat(table[i][k], 64)
				handleError(err)
				f2, err := strconv.ParseFloat(table[j][k], 64)
				handleError(err)
				return f1 > f2
			})
		} else {
			sort.Slice(table, func(i, j int) bool {
				f1, err := strconv.ParseFloat(table[i][k], 64)
				handleError(err)
				f2, err := strconv.ParseFloat(table[j][k], 64)
				handleError(err)
				return f1 < f2
			})
		}
	} else {
		if isF && isR {
			sort.Slice(table, func(i, j int) bool {
				return strings.ToLower(table[i][k]) > strings.ToLower(table[j][k])
			})
		} else if isF && !isR {
			sort.Slice(table, func(i, j int) bool {
				return strings.ToLower(table[i][k]) < strings.ToLower(table[j][k])
			})
		} else if !isF && isR {
			sort.Slice(table, func(i, j int) bool {
				return table[i][k] > table[j][k]
			})
		} else if !isF && !isR {
			sort.Slice(table, func(i, j int) bool {
				return table[i][k] < table[j][k]
			})
		}
	}

	return nil
}

func printSlice(slice []string, out io.Writer) {
	fmt.Fprintln(out, strings.Trim(fmt.Sprint(slice), "[]"))
}

func writeStrings(table [][]string, isU bool, isF bool, k int, out io.Writer) error {
	if k < 0 {
		return errors.New("Column under zero")
	}

	if isU && isF {
		for i := range table {
			if i > 0 && strings.ToLower(table[i][k]) != strings.ToLower(table[i-1][k]) {
				printSlice(table[i], out)
			} else if i == 0 {
				printSlice(table[i], out)
			}
		}
	} else if isU && !isF {
		for i := range table {
			if i > 0 && table[i][k] != table[i-1][k] {
				printSlice(table[i], out)
			} else if i == 0 {
				printSlice(table[i], out)
			}
		}
	} else if !isU {
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
	if len(args) == 0 {
		fmt.Println("No arguments")
		os.Exit(1)
	}

	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile(args[0])
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))
	for i, line := range lines {
		if *kPtr > 0 {
			tableOfStings[i] = strings.Split(line, " ")
		} else {
			tableOfStings[i] = make([]string, 1)
			tableOfStings[i][0] = line
		}
	}

	// пользователь ведёт нумерацию с 1
	if *kPtr > 0 {
		*kPtr = *kPtr - 1
	}
	errSort := sortStrings(tableOfStings, *fPtr, *rPtr, *nPtr, *kPtr)
	handleError(errSort)

	if *oPtr != "" {
		out, err := os.Create(*oPtr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer out.Close()
		errWrite := writeStrings(tableOfStings, *uPtr, *fPtr, *kPtr, out)
		handleError(errWrite)
	} else {
		errWrite := writeStrings(tableOfStings, *uPtr, *fPtr, *kPtr, os.Stdout)
		handleError(errWrite)
	}
}
