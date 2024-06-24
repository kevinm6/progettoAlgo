package main

import (
	// "bufio"
	// "fmt"
	"os"
)

func newPiano() *piano {
	return &piano{piastrelle: make(map[punto]*piastrella)}
}

func dfsBlocco(P *piano, p punto, visited map[punto]bool, blocco *[]*piastrella) {
	if _, exists := P.piastrelle[p]; !exists || visited[p] {
		return
	}

	visited[p] = true
	*blocco = append(*blocco, P.piastrelle[p])
	direzione := []punto{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for _, d := range direzione {
		neighbor := punto{p.x + d.x, p.y + d.y}
		dfsBlocco(P, neighbor, visited, blocco)
	}
}

// func readInputs() {
// 	scanner := bufio.NewScanner(os.Stdin)
//
// 	for scanner.Scan() {
// 		row := scanner.Text()
// 		fmt.Println("Read => ", row)
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		fmt.Fprintln(os.Stderr, "Error reading from stdin: ", err)
// 	}
// }

func quit() {
	// Termina l'esecuzione del programma
	//  viene chiamata quando viene letto `q` da stdin
	os.Exit(0)
}
