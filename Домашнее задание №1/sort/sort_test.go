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
	if errRead != nil {
		t.Errorf(errRead.Error())
	}
	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var options Opts = Opts{
		Numeric:  true,
		Key:      0,
		FoldCase: false,
		Reverse:  false,
		Unique:   false,
	}
	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		t.Errorf(errSort.Error())
	}

	out, err := os.Create("TestNumbers.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, options, out)
	if errWrite != nil {
		t.Errorf(errWrite.Error())
	}

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	result := string(rbytes)
	if result != sortedNum {
		t.Errorf("Test NumSort failed: result not match with sortedNum")
	}
}

func TestNumSortUniq(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("numbers.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var options Opts = Opts{
		Numeric:  true,
		Key:      0,
		FoldCase: false,
		Reverse:  false,
		Unique:   true,
	}
	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		t.Errorf(errSort.Error())
	}

	out, err := os.Create("TestNumbers2.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, options, out)
	if errWrite != nil {
		t.Errorf(errWrite.Error())
	}

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers2.txt")
	if errRead != nil {
		t.Errorf(err.Error())
	}

	result := string(rbytes)
	if result != uniqSortedNum {
		t.Errorf("Test NumSortUniq failed: result not match with uniqSortedNum")
	}
}

func TestNumSortUniqReverse(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("numbers.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var options Opts = Opts{
		Numeric:  true,
		Key:      0,
		FoldCase: false,
		Reverse:  true,
		Unique:   true,
	}

	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		t.Errorf(errSort.Error())
	}

	out, err := os.Create("TestNumbers3.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, options, out)
	if errWrite != nil {
		t.Errorf(errWrite.Error())
	}

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestNumbers3.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	result := string(rbytes)
	if result != uniqReverseSortedNum {
		t.Errorf("Test NumSortUniqReverse failed: result not match with ")
	}
}

func TestNumSortFailOnText(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("text.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}

	var options Opts = Opts{
		Numeric:  true,
		Key:      0,
		FoldCase: false,
		Reverse:  true,
		Unique:   false,
	}

	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		t.Errorf(errSort.Error())
	}
}

func TestSortFUniqReverse(t *testing.T) {
	// читаем данные из файла
	bytes, errRead := ioutil.ReadFile("text.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	// разбиваем на строки
	lines := strings.Split(string(bytes), "\n")
	tableOfStings := make([][]string, len(lines))

	for i, line := range lines {
		tableOfStings[i] = make([]string, 1)
		tableOfStings[i][0] = line
	}
	var options Opts = Opts{
		Numeric:  false,
		Key:      0,
		FoldCase: true,
		Reverse:  true,
		Unique:   true,
	}

	errSort := sortStrings(tableOfStings, options)
	if errSort != nil {
		t.Errorf(errSort.Error())
	}

	out, err := os.Create("TestSort.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer out.Close()

	errWrite := writeStrings(tableOfStings, options, out)
	if errWrite != nil {
		t.Errorf(errWrite.Error())
	}

	// читаем данные из файла
	rbytes, errRead := ioutil.ReadFile("TestSort.txt")
	if errRead != nil {
		t.Errorf(errRead.Error())
	}

	result := string(rbytes)
	if result != sortedUniqReverseText {
		t.Errorf("Test NumSortUniqReverse failed: result does not match")
	}
}
