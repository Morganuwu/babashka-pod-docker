package test

import (
	"strings"
	"testing"
)

func TestInitFunction(t *testing.T) {
	passed := true
	printResult("InitFunction", passed)
	if !passed {
		t.Errorf("Falló test de ejemplo")
	}
}

func TestAdditionLogic(t *testing.T) {
	result := 2 + 3
	expected := 5
	passed := result == expected
	printResult("AdditionLogic", passed)
	if !passed {
		t.Errorf("incorrecta")
	}
}

func TestSubtractionLogic(t *testing.T) {
	result := 10 - 4
	expected := 6
	passed := result == expected
	printResult("SubtractionLogic", passed)
	if !passed {
		t.Errorf("incorrecta")
	}
}

func TestMultiplicationLogic(t *testing.T) {
	result := 3 * 4
	expected := 12
	passed := result == expected
	printResult("MultiplicationLogic", passed)
	if !passed {
		t.Errorf("Multiplicación incorrecta")
	}
}

func TestDivisionLogic(t *testing.T) {
	result := 20 / 4
	expected := 5
	passed := result == expected
	printResult("DivisionLogic", passed)
	if !passed {
		t.Errorf("División incorrecta")
	}
}

func TestModuloLogic(t *testing.T) {
	result := 7 % 3
	expected := 1
	passed := result == expected
	printResult("ModuloLogic", passed)
	if !passed {
		t.Errorf("Módulo incorrecto")
	}
}

func TestStringParser(t *testing.T) {
	s := "hola"
	passed := s == "hola"
	printResult("StringParser", passed)
	if !passed {
		t.Errorf("String incorrecto")
	}
}

func TestStringContains(t *testing.T) {
	s := "hello world"
	passed := strings.Contains(s, "world")
	printResult("StringContains", passed)
	if !passed {
		t.Errorf("String no contiene subcadena")
	}
}

func TestStringLength(t *testing.T) {
	s := "test"
	passed := len(s) == 4
	printResult("StringLength", passed)
	if !passed {
		t.Errorf("Longitud incorrecta")
	}
}

func TestBooleanFlag(t *testing.T) {
	ok := true
	passed := ok
	printResult("BooleanFlag", passed)
	if !passed {
		t.Errorf("Booleano falso")
	}
}

func TestToggleFlag(t *testing.T) {
	flag := false
	flag = !flag
	passed := flag
	printResult("ToggleFlag", passed)
	if !passed {
		t.Errorf("Toggle fallido")
	}
}

func TestArrayLength(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	passed := len(arr) == 4
	printResult("ArrayLength", passed)
	if !passed {
		t.Errorf("Longitud de array incorrecta")
	}
}

func TestArraySum(t *testing.T) {
	arr := []int{1, 2, 3}
	sum := 0
	for _, v := range arr {
		sum += v
	}
	passed := sum == 6
	printResult("ArraySum", passed)
	if !passed {
		t.Errorf("incorrecta")
	}
}

func TestMapKeys(t *testing.T) {
	m := map[string]int{"a": 1, "b": 2}
	passed := len(m) == 2
	printResult("MapKeys", passed)
	if !passed {
		t.Errorf("Map tiene keys incorrectas")
	}
}

func TestMapValue(t *testing.T) {
	m := map[string]int{"a": 1}
	passed := m["a"] == 1
	printResult("MapValue", passed)
	if !passed {
		t.Errorf("Valor de map incorrecto")
	}
}

func TestNilCheck(t *testing.T) {
	var ptr *int
	passed := ptr == nil
	printResult("NilCheck", passed)
	if !passed {
		t.Errorf("Nil check falló")
	}
}

func TestLoopIteration(t *testing.T) {
	count := 0
	for i := 0; i < 5; i++ {
		count++
	}
	passed := count == 5
	printResult("LoopIteration", passed)
	if !passed {
		t.Errorf("Loop iteration falló")
	}
}

func TestConditionTrue(t *testing.T) {
	val := 10
	passed := val > 5
	printResult("ConditionTrue", passed)
	if !passed {
		t.Errorf("Condición falló")
	}
}

func TestConditionFalse(t *testing.T) {
	val := 3
	passed := val < 5
	printResult("ConditionFalse", passed)
	if !passed {
		t.Errorf("Condición falló")
	}
}

func TestGitPushSimulation(t *testing.T) {
	cmd := "git push origin main"
	passed := strings.Contains(cmd, "push")
	printResult("GitPushSimulation", passed)
	if !passed {
		t.Errorf("Simulación de push fallida")
	}
}

func TestGitPullSimulation(t *testing.T) {
	cmd := "git pull origin main"
	passed := strings.Contains(cmd, "pull")
	printResult("GitPullSimulation", passed)
	if !passed {
		t.Errorf("Simulación de pull fallida")
	}
}

// func TestErrorHandler(t *testing.T) {
// 	passed := false
// 	printResult("ErrorHandler", passed)
// 	if !passed {
// 		t.Errorf("Falla intencional")
// 	}
// }
