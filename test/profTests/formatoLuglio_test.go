package main

import (
	"testing"
)

var prog = "./solution"
var verbose = true

func TestBaseColoraStato(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"../../inFiles/BaseColoraStato",
		"../../outFiles/BaseColoraStato",
		verbose,
	)
}

func TestBaseRegole(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"../../inFiles/BaseRegole",
		"../../outFiles/BaseRegole",
		verbose,
	)
}

func TestBaseBlocco(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"../../inFiles/BaseBlocco",
		"../../outFiles/BaseBlocco",
		verbose,
	)
}

func TestBasePista(t *testing.T) {
	LanciaGenericaConFileInOutAtteso(
		t,
		prog,
		"../../inFiles/BasePista",
		"../../outFiles/BasePista",
		verbose,
	)
}
