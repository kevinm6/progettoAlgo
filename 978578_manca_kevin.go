package main

import (
	"bufio"
	"fmt"
	"os"
)

// import "fmt"

type piastrella struct {
	x, y      int
	color     string
	intensità int
	accesa    bool
	// stato stato_piastrella
}

type piano struct {
	piastrelle map[punto]*piastrella
	ps         piastrella
	regole     string
}

// type stato_piastrella struct {
// 	intensità int // > num => > intensità  (impostare magari un limite a 5 ?)
// 	accesa    bool
// }

type punto struct {
	x, y int
}

func main() {
	/*
	 * Il programma deve leggere dallo standard input una sequenza di righe (separate
	 * da a-capo), ciascuna delle quali corrisponde a una linea della prima colonna
	 * della Tabella 1, dove x, y, kᵢ sono interi e α, αᵢ β, s sono stringhe finite
	 * di lunghezza arbitraria. I vari elementi sulla riga sono separati da uno
	 * spazio. Quando una riga è letta, viene eseguita l’operazione associata; le
	 * operazioni di stampa sono effettuate sullo standard output, e ogni operazione
	 * deve iniziare su una nuova riga.
	 *
	 * esempio di regole di esecuzione
	 * 2g + 1b → z
	 * lg + 2b → w
	 * 1b + 1r → y
	 * 2b + 1r → g
	 * 1b + 1g + 1r → t.
	 */

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		row := scanner.Text()
		fmt.Println("Read => ", row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin: ", err)
	}
}

func esegui(p piano, s string) {

}

func colora(p piano, x int, y int, alpha string) {
	// INPUT: C x y alpha
	// colora(x, y, α) Colora Piastrella(x, y) di colore α, qualunque sia lo
	// stato di Piastrella(x, y) prima dell’operazione.
	P := punto{x, y}
	if tile, exists := p.piastrelle[P]; exists {
		// se la piastrella esiste
		tile.color = alpha
	} else {
		p.piastrelle[P] = &piastrella{x: x, y: y, color: alpha, intensità: 0}
	}
}

func spegni(p piano, x int, y int) {
	// INPUT: S x y
	// Spegne Piastrella(x, y). Se Piastrella(x, y) è gi`a spenta, non fa nulla.
}

func regola(p piano, r string) {
	// INPUT: r r
	// Definisce la regola di propagazione k1α1 + k2α2 + · · · + knαn → β e la
	// inserisce in fondo all’elenco delle regole.
}

func stato(p piano, x int, y int) {
	// INPUT: ? x y
	// Stampa e restituisce il colore e l’intensità di
	// Piastrella(x, y). Se Piastrella(x, y) è spenta, non stampa nulla.
	P := punto{x, y}
	tile := p.piastrelle[P]
	if tile.accesa {
		fmt.Printf("%s %d", tile.color, tile.intensità)
	}
}

func stampa(p piano) {
	// INPUT: s
	// Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
}

func blocco(x int, y int) {
	// INPUT: b x y
	// Calcola le stampa a somma delle intensità delle piastrelle contenute nel
	// blocco di appartenenza di Piastrella(x, y). Se Piastrella(x, y) è spenta,
	// restituisce 0.
	P := newPiano()
	p := punto{x, y}
	visited := make(map[punto]bool)
	blocco := []*piastrella{}
	dfsBlocco(P, p, visited, &blocco)
	fmt.Println()
}

func bloccoOmog(x int, y int) {
	// INPUT: B x y
	// Calcola e stampa la somma delle intensità delle piastrelle contenute nel
	// blocco omogeneo di appartenenza di Piastrella(x, y). Se Piastrella(x, y)
	// è spenta, restituisce 0.
}

func propaga(x int, y int) {
	// INPUT: p x y
	// Applica a Piastrella(x, y) la prima regola di propagazione applicabile
	// dell’elenco, ricolorando la piastrella. Se nessuna regola è applicabile, non
	// viene eseguita alcuna operazione.
}

func propagaBlocco(x int, y int) {
	// INPUT: P x y
	// Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
}

func ordina(p piano) {
	// INPUT: o
	// Ordina l’elenco delle regole di propagazione in base al consumo delle
	// regole stesse: la regola con consumo maggiore diventa l’ultima
	// dell’elenco. Se due regole hanno consumo uguale mantengono il loro ordine
	// relativo.
}

func pista(p piano, x int, y int, s string) {
	// INPUT: t x y s
	// Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di
	// direzioni s, se tale pista è definita. Altrimenti non stampa nulla.
}

func lung(x1 int, y1 int, x2 int, y2 int) {
	// INPUT: L x y
	// Determina la lunghezza della pista piu` breve che parte da Piastrella(x1, y1) e
	// arriva in Piastrella(x2, y2). Altrimenti non stampa nulla.
}
