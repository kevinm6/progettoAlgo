package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

var prog = "./978578_manca_kevin"
var verbose = false

func LanciaGenericaConFileInOutAtteso(t *testing.T, prog string, inputFile string, expectedFile string, verbose bool) {
	input, err := os.ReadFile(inputFile)
	if err != nil {
		t.Fatalf("Failed to read input file %s: %v", inputFile, err)
	}

	expected, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read expected output file %s: %v", expectedFile, err)
	}

	cmd := exec.Command(prog)
	cmd.Stdin = bytes.NewReader(input)

	var out bytes.Buffer
	cmd.Stdout = &out

	err = cmd.Run()
	if err != nil {
		t.Fatalf("Command failed: %v", err)
	}

	if verbose {
		t.Logf("Input: %s", input)
		t.Logf("Expected Output: %s", expected)
		t.Logf("Actual Output: %s", out.String())
	}

	if out.String() != string(expected) {
		t.Fatalf("Output does not match expected output\nGot:\n%s\nExpected:\n%s", out.String(), string(expected))
	}
}

// DISABLED => consecutive `go test` runs doesn't return consistent results.
// func TestBuildProject(t *testing.T) {
// 	cmd := exec.Command("go", "build")
// 	err := cmd.Run()
// 	if err != nil {
// 		t.Errorf("Build failed with error: %v", err)
// 	}
// }

func TestBaseColoraStato(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BaseColoraStato",
		"./test/myTests/expected/BaseColoraStato",
		verbose,
	)
}

func TestBaseRegole(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BaseRegole",
		"./test/myTests/expected/BaseRegole",
		verbose,
	)
}

func TestBaseBlocco(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BaseBlocco",
		"./test/myTests/expected/BaseBlocco",
		verbose,
	)
}

func TestBasePropaga(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BasePropaga",
		"./test/myTests/expected/BasePropaga",
		verbose,
	)
}

func TestBasePropaga2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BasePropaga2",
		"./test/myTests/expected/BasePropaga2",
		verbose,
	)
}

func TestBasePista(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BasePista",
		"./test/myTests/expected/BasePista",
		verbose,
	)
}

func TestBaseBloccoOrdina(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BasePropagaBloccoOrdina",
		"./test/myTests/expected/BasePropagaBloccoOrdina",
		verbose,
	)
}

func TestBaseSpegni(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BaseSpegni",
		"./test/myTests/expected/BaseSpegni",
		verbose,
	)
}

func TestBaseSpegni2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/myTests/input/BaseSpegni2",
		"./test/myTests/expected/BaseSpegni2",
		verbose,
	)
}
