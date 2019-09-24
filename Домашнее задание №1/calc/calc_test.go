package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	testPlus1 string = "2+3"
	testPlus2 string = "1+2+3+4+2.5"
	testPlus3 string = "(1+2)+3+(4+2.5)"
	testPlus4 string = "(1+(2+(3+(4+(2.5)))))"

	testMinus1 string = "2-3"
	testMinus2 string = "1-2-3-4-2.5"
	testMinus3 string = "(1-2)-3-(4-2.5)"
	testMinus4 string = "(1-(2-(3-(4-(2.5)))))"

	testUnMinus1 string = "1-(-2)"
	testUnMinus2 string = "-1-(-(-5-1))"

	testMult1 string = "2*3"
	testMult2 string = "2*3*4*0.5"
	testMult3 string = "2+3*4"

	testDiv1 string = "1/4"
	testDiv2 string = "100/2/5/10"
	testDiv3 string = "100/2+100/4-100/5"

	testComplex1 string = "(-(2-5))*((2+0.5)/(-0.5))"
	testComplex2 string = "-(2*3+4/2-(5/(2+3)))"

	testFail1 string = "(2+3)-4)"
	testFail2 string = "2+(b-4)"
)

func TestFailMode(t *testing.T) {
	_, err := calculate(testFail1)
	require.Error(t, err)

	_, err = calculate(testFail2)
	require.Error(t, err)
}
func TestComplexMode(t *testing.T) {
	result, _ := calculate(testComplex1)
	require.Equal(t, result, -15.)

	result, _ = calculate(testComplex2)
	require.Equal(t, result, -7.)
}
func TestDivMode(t *testing.T) {
	result, _ := calculate(testDiv1)
	require.Equal(t, result, 0.25)

	result, _ = calculate(testDiv2)
	require.Equal(t, result, 1.)

	result, _ = calculate(testDiv3)
	require.Equal(t, result, 55.)
}

func TestMultMode(t *testing.T) {
	result, _ := calculate(testMult1)
	require.Equal(t, result, 6.)

	result, _ = calculate(testMult2)
	require.Equal(t, result, 12.)

	result, _ = calculate(testMult3)
	require.Equal(t, result, 14.)
}

func TestUnMinusMode(t *testing.T) {
	result, _ := calculate(testUnMinus1)
	require.Equal(t, result, 3.)

	result, _ = calculate(testUnMinus2)
	require.Equal(t, result, -7.)
}

func TestMinusMode(t *testing.T) {
	result, _ := calculate(testMinus1)
	require.Equal(t, result, -1.)

	result, _ = calculate(testMinus2)
	require.Equal(t, result, -10.5)

	result, _ = calculate(testMinus3)
	require.Equal(t, result, -5.5)

	result, _ = calculate(testMinus4)
	require.Equal(t, result, 0.5)
}
func TestPlusMode(t *testing.T) {
	result, _ := calculate(testPlus1)
	require.Equal(t, result, 5.)

	result, _ = calculate(testPlus2)
	require.Equal(t, result, 12.5)

	result, _ = calculate(testPlus3)
	require.Equal(t, result, 12.5)

	result, _ = calculate(testPlus4)
	require.Equal(t, result, 12.5)
}

func TestFailCalculations(t *testing.T) {
	_, e := splitTokens("(2+b)/3")
	if e == nil {
		t.Error("")
	}
}
