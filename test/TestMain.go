package test

import (
	"fmt"
	"os"
	"testing"
)

var totalTests, passedTests, failedTests int

func printResult(testName string, passed bool) {
	totalTests++
	status := "FAIL"
	if passed {
		status = "PASS"
		passedTests++
	} else {
		failedTests++
	}
	fmt.Printf("| %-20s | %-5s |\n", testName, status)
}

// TestMain ejecuta todos los tests y luego imprime un resumen
func TestMain(m *testing.M) {
	fmt.Println("| TEST NAME            | STATUS |")
	fmt.Println("|---------------------|--------|")
	exitVal := m.Run()
	fmt.Println("\n| TOTAL TESTS | PASSED | FAILED |")
	fmt.Printf("| %-11d | %-6d | %-6d |\n", totalTests, passedTests, failedTests)
	os.Exit(exitVal)
}
