package tests

import (
	"testing"
	"github.com/Kanna727/GoLearn/src"
	"github.com/Kanna727/GoLearn/utils"
)

func TestHelloWorld(t *testing.T) {
	expected := "Hello World"
	actual := utils.GetConsoleOutput(src.HelloWorld)
	if actual != expected {
		t.Errorf("HelloWorld() = %v, expected %v", actual, expected)
	}
}

func TestValues(t *testing.T) {
	expected := "GoLang\n1+1 = 2\n7.0/3.0 = 2.3333333333333335\nfalse\ntrue\nfalse\n"
	actual := utils.GetConsoleOutput(src.Values)
	if actual != expected {
		t.Errorf("Values() = %v, expected %v", actual, expected)
	}
}

func TestVariables(t *testing.T) {
	expected := "initial 1 2 true 0 short\n"
	actual := utils.GetConsoleOutput(src.Variables)
	if actual != expected {
		t.Errorf("Variables() = %v, expected %v", actual, expected)
	}
}

func TestConstans(t *testing.T) {
	expected := "constant\n6e+11\n600000000000\n-0.28470407323754404\n"
	actual := utils.GetConsoleOutput(src.Constants)
	if actual != expected {
		t.Errorf("Constants() = %v, expected %v", actual, expected)
	}
}

func TestFor(t *testing.T) {
	expected := "1\n2\n3\n7\n8\n9\nloop\n1\n3\n5\n"
	actual := utils.GetConsoleOutput(src.For)
	if actual != expected {
		t.Errorf("For() = %v, expected %v", actual, expected)
	}
}