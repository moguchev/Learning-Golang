package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var uniqReverseSortedNum = `8
7
6
5
4
3
2
1
`
var uniqSortedNum = `1
2
3
4
5
6
7
8
`

var sortedNum = `1
2
3
3
4
4
5
6
7
8
`
var sortedUniqReverseText = `Napkin
January
Hauptbahnhof
Go
Book
Apple
`

func TestNumSort(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("numbers.txt")
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var n bool = true
	var k int = 0
	var f bool = false

	var r bool = false
	var u bool = false
	errSort := sortStrings(tableOfStings, f, r, n, k)
	handleError(errSort)

	out, err := os.Create("TestNumbers.txt")
	if err != nil {
		t.Errorf("Файл не создан")
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, u, f, k, out)
	handleError(errWrite)

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers.txt")
	handleError(errRead)

	result := string(rbytes)
	if result != sortedNum {
		t.Errorf("Test NumSort failed: result not match with sortedNum")
	}
}

func TestNumSortUniq(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("numbers.txt")
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var n bool = true
	var k int = 0
	var f bool = false

	var r bool = false
	var u bool = true

	errSort := sortStrings(tableOfStings, f, r, n, k)
	handleError(errSort)

	out, err := os.Create("TestNumbers2.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, u, f, k, out)
	handleError(errWrite)

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers2.txt")
	handleError(errRead)

	result := string(rbytes)
	if result != uniqSortedNum {
		t.Errorf("Test NumSortUniq failed: result not match with uniqSortedNum")
	}
}

func TestNumSortUniqReverse(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("numbers.txt")
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var n bool = true
	var k int = 0
	var f bool = false

	var r bool = true
	var u bool = true

	errSort := sortStrings(tableOfStings, f, r, n, k)
	handleError(errSort)

	out, err := os.Create("TestNumbers3.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, u, f, k, out)
	handleError(errWrite)

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers3.txt")
	handleError(errRead)

	result := string(rbytes)
	if result != uniqReverseSortedNum {
		t.Errorf("Test NumSortUniqReverse failed: result not match with ")
	}
}

func TestNumSortFailOnText(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("text.txt")
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var n bool = true
	var k int = 0
	var f bool = false
	var r bool = true

	errSort := sortStrings(tableOfStings, f, r, n, k)
	handleError(errSort)
	if errSort != nil {
		t.Errorf("Test FAIL failed: no error")
	}
}

func TestSortFUniqReverse(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("text.txt")
	handleError(errRead)

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var n bool = false
	var k int = 0
	var f bool = true

	var r bool = true
	var u bool = true

	errSort := sortStrings(tableOfStings, f, r, n, k)
	handleError(errSort)

	out, err := os.Create("TestSort.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, u, f, k, out)
	handleError(errWrite)

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestSort.txt")
	handleError(errRead)

	result := string(rbytes)
	if result != sortedUniqReverseText {
		t.Errorf("Test NumSortUniqReverse failed: result does not match")
	}
}
