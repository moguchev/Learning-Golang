package main

import (
	"testing"
)

func TestTransformFunction(t *testing.T) {
	test1, err := transform("(2+31)/((-2.5)*4)")
	if err != nil {
		t.Error("Error not expected")
	}
	if test1 != "( 2 + 31 ) / ( ( - 2.5 ) * 4 ) " {
		t.Error(test1)
		t.Error("Do not match results")
	}

	test2, err := transform("(22-11*5)/(2.555-(-3))")
	if err != nil {
		t.Error("Error not expected")
	}
	if test2 != "( 22 - 11 * 5 ) / ( 2.555 - ( - 3 ) ) " {
		t.Error(test2)
		t.Error("Do not match result with: ", "( 22 - 11 * 5 ) / ( 2.555 - ( - 3 ) ) ")
	}

	_, e := transform("(2.2.-11*5)/(2.555-(-3))")
	if e == nil {
		t.Error("Error expected: wrong number 2.2.")
	}
}

func TestConvertInfixToPostfix(t *testing.T) {
	test1, err := convertInfixToPostfix("( 2 + 1 ) - ( 35 + 1 )")
	if err != nil {
		t.Error("Error not expected")
	}
	if test1 != "2 1 + 35 1 + -" {
		t.Error(test1)
		t.Error("Do not match with: ", "2 1 + 35 1 + -")
	}
	test2, err := convertInfixToPostfix("5 + 3 * ( - 2 )")
	if err != nil {
		t.Error("Error not expected")
	}
	if test2 != "5 3 2 ! * +" {
		t.Error(test2)
		t.Error("Do not match results")
	}
	_, e := convertInfixToPostfix(") 5 +  * ( - 2 ) )")
	if e == nil {
		t.Error("Error expected: wrong brackets")
	}
}

func TestCalculations(t *testing.T) {
	trueExpression, e := transform("(5-1)*(3/4)")
	if e != nil {
		t.Error("Error not expected")
	}

	postfix, e := convertInfixToPostfix(trueExpression)
	if e != nil {
		t.Error("Error not expected")
	}

	result, e := parsePostfix(postfix)
	if e != nil {
		t.Error("Error not expected")
	}
	if result != 3 {
		t.Error("expected 3")
	}

	trueExpression, e = transform("(2/4)-(-2)")
	if e != nil {
		t.Error("Error not expected")
	}

	postfix, e = convertInfixToPostfix(trueExpression)
	if e != nil {
		t.Error("Error not expected")
	}

	result, e = parsePostfix(postfix)
	if e != nil {
		t.Error("Error not expected")
	}
	if result != 2.5 {
		t.Error("expected 2.5")
	}

	trueExpression, e = transform("-(-2)")
	if e != nil {
		t.Error("Error not expected")
	}

	postfix, e = convertInfixToPostfix(trueExpression)
	if e != nil {
		t.Error("Error not expected")
	}

	result, e = parsePostfix(postfix)
	if e != nil {
		t.Error("Error not expected")
	}

	if result != 2 {
		t.Error("expected 2")
	}

}

func TestFailCalculations(t *testing.T) {
	_, e := transform("(2+b)/3")
	if e == nil {
		t.Error("Error expected")
	}
}
