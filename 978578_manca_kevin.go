/*
---
Title    : Progetto Laboratorio di Algoritmi & Strutture Dati @ Unimi
Author   : Kevin Manca
Matricola: 978578
Data     : 10/07/2024
---
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

//
/************************************ STRUTTURE ******************************************/
//

// piano => rappresenta il piano, cioè un insieme di piastrelle e regole
type Piano struct {
	piastrelle map[punto]*Piastrella // mappa che rappresenta l'insieme delle piastrelle, { Punto: ->Piastrella }
	regole     []*Regola             // insieme di regole
}

// Type alias: piano --puntatore-> Piano
// Utile per rispettare le segnature delle funzioni e per passare
// un type Piano per indirizzo (nonostante sia comunque una struttura composta da altri indirizzi)
// Inoltre aiuta ad evitare errori nel codice dovuti ai puntatori
type piano *Piano

// Piastrella => Rappresenta una piastrella del piano
type Piastrella struct {
	x, y      int    // coordinate che identificano la piastrella
	colore    string // colore della piastrella
	intensità int    // intensità della piastrella (0: spenta, >1: accesa)
}

// punto => Rappresenta un vertice con coordinate (x, y)
type punto struct {
	x, y int
}

// Regola => regola
type Regola struct {
	istruzioneCompleta string // stringa della regola completa
	// elementi  []elemRegola   // insieme di istruzioni che compongono la regola
	colore    string         // nuovo colore della regola
	consumo   uint           // numero tot di piastrelle a cui è stata applicata la regola
	valColore map[string]int // mappa il colore col relativo valore
}

// elemRegola => istruzioni della regola (formato: k₁α₁ + k₂α₂ + ... + kₙαₙ → β)
type elemRegola struct {
	k     int    // kᵢ sono interi positivi la cui somma non supera 8
	alpha string // stringa dell'alfabeto tutte differenti tra loro
}

// Rappresentano le direzioni di movimento nel piano
// La mappa e la relativa chiave forniscono una veloce interpretazione delle coordinate (il valore)
// ed è utile per la mappa delle direzioni nella funzione `calcolaPista()`
var direzioni = map[string]punto{
	"OO": {-1, 0},
	"EE": {1, 0},
	"SS": {0, -1},
	"NN": {0, 1},
	"SO": {-1, -1},
	"SE": {1, -1},
	"NO": {-1, 1},
	"NE": {1, 1},
}

//
/************************************** MAIN ********************************************/
func main() {

	// Crea e inizializza un piano vuoto e le regole (vuote)
	p := Piano{
		piastrelle: make(map[punto]*Piastrella),
		regole:     make([]*Regola, 0),
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		esegui(&p, row)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin: ", err)
	}
}

//
/************************************* ESEGUI *******************************************/
func esegui(p piano, s string) {
	flags := strings.Fields(s)
	var x, y int

	// fmt.Println("flags => ", flags)
	if len(flags) > 2 {
		x, _ = strconv.Atoi(flags[1])
		y, _ = strconv.Atoi(flags[2])
	}

	switch flags[0] {
	case "C": // Blocco Omogeneo
		alpha := flags[3]
		intensity, _ := strconv.Atoi(flags[4])
		colora(p, x, y, alpha, intensity)

	case "r": // Regola
		// Le regole hanno questa forma: r x 1 a
		// ottengo una string-slice dalla stringa di partenza per rimuovere `r `, ottenendo 'x 1 a'
		r := s[2:]
		regola(p, r)

	case "?": // Stato
		stato(p, x, y)

	case "S": // Spegni
		spegni(p, x, y)

	case "s": // Stampa
		stampa(p)

	case "B": // Blocco Omogeneo
		blocco(p, x, y, true, true)

	case "b": // Blocco
		blocco(p, x, y, false, true)

	case "p": // Propaga
		propaga(p, x, y)

	case "P": // Propaga Blocco
		propagaBlocco(p, x, y)

	case "o": // Ordina
		ordina(p)

	case "t": // Pista
		sequenzaDirezioni := flags[3]
		pista(p, x, y, sequenzaDirezioni)

	case "L": // Lunghezza
		x2, _ := strconv.Atoi(flags[3])
		y2, _ := strconv.Atoi(flags[4])
		lung(p, x, y, x2, y2)

	case "q":
		quit()

	default:
		return
	}
}

// Colora Piastrella(x, y) di colore α,
// qualunque sia lo stato di Piastrella(x, y) prima dell’operazione.
func colora(p piano, x int, y int, alpha string, intensità int) {
	coordinate := punto{x, y}
	if piastrella, ok := p.piastrelle[coordinate]; ok && piastrella.intensità == 0 {
		piastrella.intensità = intensità
		piastrella.colore = alpha
	} else {
		p.piastrelle[coordinate] = &Piastrella{x: coordinate.x, y: coordinate.y, colore: alpha, intensità: intensità}
	}
}

// Spegne Piastrella(x, y).
// Se Piastrella(x, y) è già spenta, non fa nulla.
func spegni(p piano, x int, y int) {
	P := punto{x, y}
	if tile := p.piastrelle[P]; tile.intensità > 0 {
		tile.intensità = 0
		// fmt.Printf("Piastrella(%d, %d) spenta con successo", tile.x, tile.y)
	}
}

// Definisce la regola di propagazione {@type istruzioneRegola}
// inserisce in fondo all’elenco delle regole.
func regola(p piano, r string) {
	istruzioni := strings.Fields(r)
	colore := istruzioni[0]
	valori := make(map[string]int)

	for i := 1; i < len(istruzioni); i += 2 {
		val, _ := strconv.Atoi(istruzioni[i])
		alpha := istruzioni[i+1]
		valori[alpha] = val
	}

	p.regole = append(p.regole, &Regola{istruzioneCompleta: r, valColore: valori, colore: colore, consumo: 0})
}

// Stampa e restituisce il colore e l’intensità di Piastrella(x, y).
// Se Piastrella(x, y) è spenta, non stampa nulla.
func stato(p piano, x int, y int) {
	P := punto{x, y}

	if piastrella, ok := p.piastrelle[P]; ok && piastrella.intensità > 0 {
		fmt.Printf("%s %d\n", piastrella.colore, piastrella.intensità)
	}
}

// Stampa l’elenco delle regole di propagazione, nell’ordine attuale.
func stampa(p piano) {
	fmt.Println("(")
	for _, regola := range p.regole {
		elementi := strings.SplitN(regola.istruzioneCompleta, " ", 2)
		if len(elementi) == 2 {
			fmt.Printf("%s: %s\n", elementi[0], elementi[1])
		}
	}
	fmt.Println(")")
}

// Calcola e stampa la somma delle intensità delle piastrelle contenute nel
// blocco (omogeneo o non in base alla variabile omonima passata come parametro)
// di appartenenza di Piastrella(x, y).
// Se Piastrella(x, y) è spenta, restituisce 0.
func blocco(p piano, x int, y int, omogeneo bool, stampaRisultato bool) (circonvicine map[punto]bool) {
	vertice := punto{x, y}
	risultatoSomma := 0

	// verifico che la piastrella esista e non sia spenta, else ∉ blocco
	if piastrella, ok := p.piastrelle[vertice]; !ok || piastrella.intensità == 0 {
		fmt.Println(risultatoSomma)
		return
	} else {
		risultatoSomma = piastrella.intensità
	}

	visite := make(map[punto]bool)

	if omogeneo {
		blocco := make(map[punto]*Piastrella)
		dfs(p, vertice, visite, blocco, true, &risultatoSomma)
	} else {
		dfs(p, vertice, visite, nil, false, &risultatoSomma)
	}

	if stampaRisultato {
		fmt.Println(risultatoSomma)
	}
	return visite
}

// Applica a Piastrella(x, y) la prima regola di propagazione applicabile
// dell’elenco, ricolorando la piastrella.
// Se nessuna regola è applicabile, non viene eseguita alcuna operazione.
func propaga(p piano, x int, y int) {
	vertice := punto{x, y}

	// piastrella, ok := p.piastrelle[vertice]
	// if !ok {
	// 	piastrella = &Piastrella{x: vertice.x, y: vertice.y, colore: "", intensità: 1}
	// }
	// println("CALLED PROPAGA")

	// if !ok { //|| piastrella.intensità == 0 {
	// 	return
	// }
	intensità := 1

	for _, r := range p.regole {
		if regola := verificaRegola(p, r, vertice, &intensità); regola != nil {
			// fmt.Println("Trovata regola da applicare! => ", regola.colore, regola.consumo)
			// Applico solo la prima regola valida e incremento il suo consumo
			// fmt.Println(piastrella)
			// colora(p, x, y, regola.colore, p.piastrelle[vertice].intensità)
			regola.consumo++
			break
		}
	}
}

// Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(p piano, x int, y int) {
	vertice := punto{x, y}

	_, ok := p.piastrelle[vertice]
	if !ok {
		// Se la piastrella è spenta, non proseguo
		return
	}
	visite := make(map[punto]bool)
	blocco := make(map[punto]*Piastrella)
	dfs(p, vertice, visite, blocco, false, nil)

	if len(blocco) > 0 {
		for vertice, piastrella := range blocco {
			if piastrella == nil {
				break
			}

			coordinate := punto{vertice.x, vertice.y}
			intensità := 1
			for _, r := range p.regole {
				if regola := verificaRegola(p, r, coordinate, &intensità); regola != nil {
					colora(p, coordinate.x, coordinate.y, regola.colore, p.piastrelle[vertice].intensità)
					r.consumo++
				}
			}
		}
	}
}

// Ordina l’elenco delle regole di propagazione in base al consumo delle
// regole stesse: la regola con consumo maggiore diventa l’ultima dell’elenco.
// Se due regole hanno consumo uguale mantengono il loro ordine relativo.
func ordina(p piano) {
	// NOTE: SliceStable, come riportato sui [docs di go](https://pkg.go.dev/sort#SliceStable)
	//     	 non modifica l'ordine originale
	sort.SliceStable(p.regole, func(i, j int) bool {
		// fmt.Println("regole => ", p.regole[i].colore, p.regole[i].consumo)
		consumoI := p.regole[i].consumo
		consumoJ := p.regole[j].consumo
		return consumoI < consumoJ
	})
}

// Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di
// direzioni s, se tale pista è definita. Altrimenti non stampa nulla.
func pista(p piano, x int, y int, s string) {
	vertice := punto{x, y}
	seqDirezioni := strings.Split(s, ",")
	piastrellePista := make(map[punto]*Piastrella)
	calcolaPista(p, vertice, seqDirezioni, piastrellePista)

	if piastrellePista != nil {
		fmt.Println("[")
		for _, piastrella := range piastrellePista {
			fmt.Printf(
				"%d %d %s %d\n", piastrella.x, piastrella.y, piastrella.colore, piastrella.intensità)
		}
		fmt.Println("]")
	}
}

// Determina la lunghezza della pista piu` breve che parte da Piastrella(x1, y1) e
// arriva in Piastrella(x2, y2).
// Altrimenti non stampa nulla.
func lung(p piano, x1 int, y1 int, x2 int, y2 int) {
	verticeOrig := punto{x1, y1}
	verticeDest := punto{x2, y2}

	pistaBreve := make(map[punto]*Piastrella)
	lunghezza := calcolaPistaBreve(p, verticeOrig, verticeDest, pistaBreve)

	if pistaBreve != nil && lunghezza > 0 {
		fmt.Println(lunghezza)
	}
}

//
/******************************** HELPER FUNCTIONS **************************************/
//

func dfs(
	p piano,
	vertice punto,
	visite map[punto]bool,
	blocco map[punto]*Piastrella,
	omogeneo bool,
	sum *int,
) {
	visite[vertice] = true
	colore := p.piastrelle[vertice].colore

	for _, direzione := range direzioni {
		nuovoVertice := punto{vertice.x + direzione.x, vertice.y + direzione.y}

		if circonvicina, ok := p.piastrelle[nuovoVertice]; ok && !visite[nuovoVertice] && circonvicina.intensità > 0 {
			if !omogeneo || (omogeneo && circonvicina.colore == colore) {
				if blocco != nil {
					blocco[vertice] = circonvicina
				}
				if sum != nil {
					*sum += circonvicina.intensità
				}
				dfs(p, nuovoVertice, visite, blocco, omogeneo, sum)
			}
		}
	}
}

// Restituisce le piastrelle circonvicine alla Piastrella(punto.x, punto.y)
func piastrelleCirconvicine(p piano, vertice punto, vicine map[punto]*Piastrella) {
	for _, direzione := range direzioni {
		nuovoVertice := punto{vertice.x + direzione.x, vertice.y + direzione.y}

		if piastrella, ok := p.piastrelle[nuovoVertice]; ok {
			vicine[nuovoVertice] = piastrella
		}
	}
	return
}

// Calcola e mappa il numero di colori dell'intorno
func calcolaColoriCirconvicine(p piano, vertice punto) map[string]int {
	colori := make(map[string]int)
	circonvicine := make(map[punto]*Piastrella)
	piastrelleCirconvicine(p, vertice, circonvicine)

	for _, piastrella := range circonvicine {
		colori[piastrella.colore]++
	}
	return colori
}

// Verifica che una regola sia applicabile in base ai colori circostanti
func verificaRegola(p piano, regola *Regola, vertice punto, intensità *int) *Regola {
	valoriColore := calcolaColoriCirconvicine(p, vertice)

	for colore, val := range regola.valColore {
		if valoriColore[colore] < val {
			return nil
		}
	}

	// Controllo se la piastrella esiste ed è accesa, altrimenti l'intensità sarà 1 di default
	piastrella, piastrellaOk := p.piastrelle[vertice]
	if piastrellaOk {
		*intensità = piastrella.intensità
	}

	return regola
}

// Calcola la pista seguendo le direzioni
func calcolaPista(p piano, vertice punto, seqDirezioni []string, piastrellePista map[punto]*Piastrella) {
	piastrella, ok := p.piastrelle[vertice]
	if !ok || piastrella.intensità == 0 {
		return
	}

	// Aggiungo la piastrella iniziale
	piastrellePista[vertice] = piastrella

	for i := 0; i < len(seqDirezioni); i++ {
		deltaX := direzioni[seqDirezioni[i]].x
		deltaY := direzioni[seqDirezioni[i]].y
		vertice = punto{vertice.x + deltaX, vertice.y + deltaY}
		// fmt.Println("CALLED ON => ", vertice, direzione)
		if altraPiastrella, ok := p.piastrelle[vertice]; ok && altraPiastrella.intensità > 0 {
			piastrellePista[vertice] = altraPiastrella
		} else {
			break
		}
	}
}

// func calcolaPistaBreve(p piano, verticeOrig punto, verticeDest punto, pistaBreve map[punto]*Piastrella) int {
// 	piastrellaOrig, origineOk := p.piastrelle[verticeOrig]
// 	piastrellaDest, destOk := p.piastrelle[verticeDest]
//
// 	// If the origin or destination tiles are not valid or not lit, return nil.
// 	if (!origineOk || piastrellaOrig.intensità == 0) || (!destOk || piastrellaDest.intensità == 0) {
// 		return 0
// 	}
//
// 	// Initialization
// 	visitate := make(map[punto]bool)
// 	precedenti := make(map[punto]punto)
// 	queue := []punto{verticeOrig}
// 	visitate[verticeOrig] = true
//
// 	// BFS with a queue
// 	for len(queue) > 0 {
// 		vertice := queue[0]
// 		queue = queue[1:]
//
// 		if vertice == verticeDest {
// 			// Reconstruct the shortest path
// 			piastrelle := []Piastrella{}
// 			for v := vertice; v != verticeOrig; v = precedenti[v] {
// 				piastrelle = append([]Piastrella{*p.piastrelle[v]}, piastrelle...)
// 			}
// 			piastrelle = append([]Piastrella{*piastrellaOrig}, piastrelle...)
//
// 			for _, piast := range piastrelle {
// 				puntoPiastrella := punto{piast.x, piast.y}
// 				pistaBreve[puntoPiastrella] = &piast
// 			}
// 			return len(piastrelle)
// 		}
//
// 		adiacenti := make(map[punto]*Piastrella)
// 		piastrelleCirconvicine(p, vertice, adiacenti)
// 		for _, piastrella := range adiacenti {
// 			coordinatePiastrella := punto{piastrella.x, piastrella.y}
// 			if !visitate[coordinatePiastrella] {
// 				queue = append(queue, coordinatePiastrella)
// 				visitate[coordinatePiastrella] = true
// 				precedenti[coordinatePiastrella] = vertice
// 			}
// 		}
// 	}
//
// 	return 0
// }

func calcolaPistaBreve(p piano, verticeOrig punto, verticeDest punto, pistaBreve map[punto]*Piastrella) int {
	piastrellaOrig, origineOk := p.piastrelle[verticeOrig]
	piastrellaDest, destOk := p.piastrelle[verticeDest]

	// Se le piastrelle di origine e destinazione non sono accese oppure non sono valide,
	//  non posso calcolare la pista
	if (!origineOk || piastrellaOrig.intensità == 0) || (!destOk || piastrellaDest.intensità == 0) {
		return 0
	}

	// Inizializzazione
	visitate := make(map[punto]bool)
	precedenti := make(map[punto]punto)
	queue := []punto{verticeOrig}
	visitate[verticeOrig] = true

	// BFS
	for len(queue) > 0 {
		vertice := queue[0]
		queue = queue[1:]

		if vertice == verticeDest {
			// Reconstruct the shortest path
			piastrelle := []Piastrella{}
			for v := vertice; v != verticeOrig; v = precedenti[v] {
				piastrelle = append([]Piastrella{*p.piastrelle[v]}, piastrelle...)
			}
			piastrelle = append([]Piastrella{*piastrellaOrig}, piastrelle...)

			// for _, piast := range piastrelle {
			// 	vertPiastrella := punto{piast.x, piast.y}
			// 	pistaBreve[vertPiastrella] = &piast
			// }
			return len(piastrelle)
		}

		adiacenti := make(map[punto]*Piastrella)
		piastrelleCirconvicine(p, vertice, adiacenti)
		for _, piastrella := range adiacenti {
			coordinatePiastrella := punto{piastrella.x, piastrella.y}
			if !visitate[coordinatePiastrella] {
				queue = append(queue, coordinatePiastrella)
				visitate[coordinatePiastrella] = true
				precedenti[coordinatePiastrella] = vertice
			}
		}
	}
	return 0
}

// Termina l'esecuzione del programma
// viene chiamata quando viene letto 'q' da stdin
func quit() {
	os.Exit(0)
}
