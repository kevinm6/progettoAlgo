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

// FIX: DISABLED => consecutive `go test` runs doesn't return consistent results.
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
		"./test/input/BaseColoraStato",
		"./test/expected/BaseColoraStato",
		verbose,
	)
}

func TestBaseRegole(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseRegole",
		"./test/expected/BaseRegole",
		verbose,
	)
}

func TestBaseBlocco(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseBlocco",
		"./test/expected/BaseBlocco",
		verbose,
	)
}

func TestBaseBlocco2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseBlocco2",
		"./test/expected/BaseBlocco2",
		verbose,
	)
}

func TestBasePropaga(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePropaga",
		"./test/expected/BasePropaga",
		verbose,
	)
}

func TestBasePropagaBloccoProf(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePropagaBloccoProf",
		"./test/expected/BasePropagaBloccoProf",
		verbose,
	)
}

func TestBasePropaga2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePropaga2",
		"./test/expected/BasePropaga2",
		verbose,
	)
}

func TestBasePropaga3(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePropaga3",
		"./test/expected/BasePropaga3",
		verbose,
	)
}

func TestBasePista(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePista",
		"./test/expected/BasePista",
		verbose,
	)
}

func TestBasePista2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePista2",
		"./test/expected/BasePista2",
		verbose,
	)
}

func TestBaseBloccoOrdina(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BasePropagaBloccoOrdina",
		"./test/expected/BasePropagaBloccoOrdina",
		verbose,
	)
}

func TestBaseSpegniPiastrellaOff(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseSpegniPiastrellaOff",
		"./test/expected/BaseSpegniPiastrellaOff",
		verbose,
	)
}

func TestBaseSpegni(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseSpegni",
		"./test/expected/BaseSpegni",
		verbose,
	)
}

func TestBaseSpegni2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/BaseSpegni2",
		"./test/expected/BaseSpegni2",
		verbose,
	)
}

func TestInputVario1(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/InputVario1",
		"./test/expected/ExpectedVario1",
		verbose,
	)
}

func TestInputVario2(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/InputVario2",
		"./test/expected/ExpectedVario2",
		verbose,
	)
}

func TestInputVario3(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/InputVario3",
		"./test/expected/ExpectedVario3",
		verbose,
	)
}

func TestInputVario4(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/InputVario4",
		"./test/expected/ExpectedVario4",
		verbose,
	)
}

func TestInputVario5(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"./test/input/InputVario5",
		"./test/expected/ExpectedVario5",
		verbose,
	)
}
