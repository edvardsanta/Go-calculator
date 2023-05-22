package main

import (
	"testing"
)

func TestCalculte_ShouldAdd(t *testing.T) {
	//arrange
	expected := 6.0

	//act
	result := calculate(2.5, 3.5, "+")

	//assert
	if result != expected {
		t.Errorf("Addition test failed: expected %.2f, got %.2f", expected, result)
	}
}
func TestCalculte_ShouldSubtract(t *testing.T) {
	//arrange
	expected := 3.0

	//act
	result := calculate(5.0, 2.0, "-")

	//assert
	if result != expected {
		t.Errorf("Subtraction test failed: expected %.2f, got %.2f", expected, result)
	}
}
func TestCalculte_ShouldMultiply(t *testing.T) {
	//arrange
	expected := 6.0

	//act
	result := calculate(2.0, 3.0, "*")

	//assert
	if result != expected {
		t.Errorf("Multiplication test failed: expected %.2f, got %.2f", expected, result)
	}
}

func TestCalculte_ShouldDivide(t *testing.T) {
	//arrange
	expected := 5.0

	//act
	result := calculate(10.0, 2.0, "/")

	//assert
	if result != expected {
		t.Errorf("Division test failed: expected %.2f, got %.2f", expected, result)
	}
}

func TestCalculate_WhenSymbolIsNothing(t *testing.T) {
	//arrange
	expected := 0.0

	//act
	result := calculate(10.0, 2.0, "")

	//assert
	if result != expected {
		t.Errorf("Symbol nothing test failed: expected %.2f, got %.2f", expected, result)
	}
}

func TestPower(t *testing.T) {
	//act
	expected := 8.0

	//act
	result := power(2.0, 3.0)

	//assert
	if result != expected {
		t.Errorf("Power test failed: expected %.2f, got %.2f", expected, result)
	}
}

func TestAll(t *testing.T) {
	t.Run("Addition", TestCalculte_ShouldAdd)
	t.Run("Subtraction", TestCalculte_ShouldSubtract)
	t.Run("Multiplication", TestCalculte_ShouldMultiply)
	t.Run("Division", TestCalculte_ShouldDivide)
	t.Run("Symbol nothing", TestCalculate_WhenSymbolIsNothing)
	t.Run("TestPower", TestPower)
}
