package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"
)

var prog = "./978578_manca_kevin"
var verbose = true

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
		t.Fatalf("Output does not match expected output\nGot: %s\nExpected: %s", out.String(), string(expected))
	}
}

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
