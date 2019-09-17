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

func sortStrings(tabel [][]string, isF bool, isR bool, isN bool, k int) error {
	if k < 0 {
		return errors.New("Column under zero")
	}

	if isN {
		if isR {
			sort.Slice(tabel, func(i, j int) bool {
				f1, err := strconv.ParseFloat(tabel[i][k], 64)
				handleError(err)
				f2, err := strconv.ParseFloat(tabel[j][k], 64)
				handleError(err)
				return f1 > f2
			})
		} else {
			sort.Slice(tabel, func(i, j int) bool {
				f1, err := strconv.ParseFloat(tabel[i][k], 64)
				handleError(err)
				f2, err := strconv.ParseFloat(tabel[j][k], 64)
				handleError(err)
				return f1 < f2
			})
		}
	} else {
		if isF && isR {
			sort.Slice(tabel, func(i, j int) bool {
				return strings.ToLower(tabel[i][k]) > strings.ToLower(tabel[j][k])
			})
		} else if isF && !isR {
			sort.Slice(tabel, func(i, j int) bool {
				return strings.ToLower(tabel[i][k]) < strings.ToLower(tabel[j][k])
			})
		} else if !isF && isR {
			sort.Slice(tabel, func(i, j int) bool {
				return tabel[i][k] > tabel[j][k]
			})
		} else if !isF && !isR {
			sort.Slice(tabel, func(i, j int) bool {
				return tabel[i][k] < tabel[j][k]
			})
		}
	}

	return nil
}

func printSlice(slice []string, file io.Writer) {
	if file == nil {
		fmt.Println(strings.Trim(fmt.Sprint(slice), "[]"))
	} else {
		fmt.Fprintln(file, strings.Trim(fmt.Sprint(slice), "[]"))
	}
}

func writeStrings(tabel [][]string, isU bool, isF bool, k int, o string) error {
	if k < 0 {
		return errors.New("Column under zero")
	}
	file, err := os.Create(o)
	if err != nil && o != "" {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if isU && isF {
		for i := range tabel {
			if i > 0 && strings.ToLower(tabel[i][k]) != strings.ToLower(tabel[i-1][k]) {
				printSlice(tabel[i], file)
			}
		}
	} else if isU && !isF {
		for i := range tabel {
			if i > 0 && tabel[i][k] != tabel[i-1][k] {
				printSlice(tabel[i], file)
			}
		}
	} else if !isU {
		for i := range tabel {
			printSlice(tabel[i], file)
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

	errWrite := writeStrings(tableOfStings, *uPtr, *fPtr, *kPtr, *oPtr)
	handleError(errWrite)

}
