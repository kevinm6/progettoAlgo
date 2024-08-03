/*
---
Title         : Progetto Laboratorio di Algoritmi & Strutture Dati @ Unimi
Author        : Kevin Manca
Matricola     : 978578
Data Consegna : 10/07/2024
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

// piano: rappresenta il piano, cioè un insieme di piastrelle e regole
type Piano struct {
	piastrelle map[punto]*Piastrella // mappa che rappresenta l'insieme delle piastrelle, { Punto: ->Piastrella }
	regole     []*Regola             // insieme di regole
}

// Type alias: piano --puntatore-> Piano
// Utile per rispettare le segnature delle funzioni e per passare
// un type Piano per indirizzo (nonostante sia comunque una struttura composta da altri indirizzi)
// Inoltre aiuta ad evitare errori nel codice dovuti ai puntatori
type piano *Piano

// Piastrella: rappresenta una piastrella del piano
type Piastrella struct {
	x, y      int    // coordinate che identificano la piastrella
	colore    string // colore della piastrella
	intensità int    // intensità della piastrella (0: spenta, >1: accesa)
}

// punto: Rappresenta un vertice con coordinate (x, y)
type punto struct {
	x, y int
}

// Regola: regola
type Regola struct {
	istruzioneCompleta string         // stringa della regola completa
	colore             string         // nuovo colore della regola
	consumo            int            // numero tot di piastrelle a cui è stata applicata la regola
	valColore          map[string]int // mappa il colore col relativo valore
}

// elemRegola : istruzioni della regola (formato: k₁α₁ + k₂α₂ + ... + kₙαₙ → β)
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

	// Crea e inizializza un piano e le regole
	p := Piano{
		piastrelle: make(map[punto]*Piastrella),
		regole:     make([]*Regola, 0),
	}

	// Leggo da stdin, non effettuo un controllo dell'input,
	//   in quanto da specifica è garantito corretto
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		row := scanner.Text()
		esegui(&p, row)
	}

	// Stampo un errore esclusivamente nel caso ci sia stato un problema con lo scanner
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading from stdin: ", err)
	}
}

//
/************************************* ESEGUI *******************************************/
func esegui(p piano, s string) {
	flags := strings.Fields(s) // restituisce la stringa in input come array di stringhe separate da spazi
	var x, y int               // inizializzo le variabili che rappresentano le coordinate

	if len(flags) > 2 {
		x, _ = strconv.Atoi(flags[1])
		y, _ = strconv.Atoi(flags[2])
	}

	switch flags[0] {
	case "C": // Blocco Omogeneo
		alpha := flags[3]
		intensità, _ := strconv.Atoi(flags[4])
		colora(p, x, y, alpha, intensità)

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
		blocco(p, x, y, true)

	case "b": // Blocco
		blocco(p, x, y, false)

	case "p": // Propaga
		fmt.Println("call propaga")
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
		fmt.Println("Input non riconosciuto!")
		return
	}
}

// Colora Piastrella(x, y) di colore α,
// qualunque sia lo stato di Piastrella(x, y) prima dell’operazione.
func colora(p piano, x int, y int, alpha string, i int) {
	coordinate := punto{x, y}
	if piastrella, ok := p.piastrelle[coordinate]; ok && piastrella.intensità == 0 {
		piastrella.intensità = i
		piastrella.colore = alpha
	} else {
		p.piastrelle[coordinate] = &Piastrella{x: coordinate.x, y: coordinate.y, colore: alpha, intensità: i}
	}
}

// Spegne Piastrella(x, y).
// Se Piastrella(x, y) è già spenta, non fa nulla.
func spegni(p piano, x int, y int) {
	P := punto{x, y}
	// FIX mancava il bool per verificare che la piastrella esistesse
	//		perciò un input tipo `S 9 9` mandava in crash il programma per un puntatore non valido
	if tile, ok := p.piastrelle[P]; ok && tile.intensità > 0 {
		tile.intensità = 0
	}
}

// Definisce la regola di propagazione e la inserisce in fondo all’elenco delle regole.
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
func stato(p piano, x int, y int) (string, int) {
	P := punto{x, y}

	if piastrella, ok := p.piastrelle[P]; ok && piastrella.intensità > 0 {
		fmt.Printf("%s %d\n", piastrella.colore, piastrella.intensità)
		return piastrella.colore, piastrella.intensità
	}
	return "", 0
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
func blocco(p piano, x int, y int, omogeneo bool) {
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

	fmt.Println(risultatoSomma)
}

// Applica a Piastrella(x, y) la prima regola di propagazione applicabile
// dell’elenco, ricolorando la piastrella.
// Se nessuna regola è applicabile, non viene eseguita alcuna operazione.
func propaga(p piano, x int, y int) {
	vertice := punto{x, y}

	intensità := 1

	/* TODO
	Ha davvero senso che mi ricalcolo ogni volta tutti i colori intorno piuttosto che farlo un unica volta? NO
	*/
	for _, regolaDaValidare := range p.regole {
		// Applico solo la prima regolaValida valida e incremento il suo consumo
		if regolaValida := verificaRegola(p, regolaDaValidare, vertice, &intensità); regolaValida != nil {
			colora(p, x, y, regolaValida.colore, intensità)
			regolaValida.consumo++
			break
		}
	}
}

// Propaga il colore sul blocco di appartenenza di Piastrella(x, y).
func propagaBlocco(p piano, x int, y int) {
	vertice := punto{x, y}

	_, ok := p.piastrelle[vertice]
	if !ok {
		return // Se la piastrella è spenta/non esiste, non proseguo
	}

	/* FIX
	DA EVITARE che si rifaccia il controllo del blocco dopo che si è applicata una regola, ma è da fare
	sul piano orgiginale prima che la regola sia stata applicata (altrimenti LOGICAMENTE CAMBIA)
	*/

	visite := make(map[punto]bool)
	blocco := make(map[punto]*Piastrella)
	// Calcolo le piastrelle del blocco con la `DFS` che mi aggiorna la mappa `blocco`
	dfs(p, vertice, visite, blocco, false, nil)

	if len(blocco) > 0 {
		for vertice, piastrella := range blocco {
			if piastrella == nil {
				break
			}

			coordinate := punto{vertice.x, vertice.y}
			intensità := 1
			for _, regolaDaValidare := range p.regole {
				if regolaValida := verificaRegola(p, regolaDaValidare, coordinate, &intensità); regolaValida != nil {
					fmt.Println(regolaValida)
					colora(p, coordinate.x, coordinate.y, regolaValida.colore, p.piastrelle[vertice].intensità)
					regolaValida.consumo++
					return
				}
			}
		}
	}
}

// Ordina l’elenco delle regole di propagazione in base al consumo delle
// regole stesse: la regola con consumo maggiore diventa l’ultima dell’elenco.
// Se due regole hanno consumo uguale mantengono il loro ordine relativo.
func ordina(p piano) {
	// NOTE:
	// SliceStable, come riportato sui [docs di go](https://pkg.go.dev/sort#SliceStable)
	//   non modifica l'ordine originale
	sort.SliceStable(p.regole, func(i, j int) bool {
		consumoI := p.regole[i].consumo
		consumoJ := p.regole[j].consumo
		// fmt.Println("i =>", p.regole[i].istruzioneCompleta, " (", consumoI, ")")
		// fmt.Println("j =>", p.regole[j].istruzioneCompleta, " (", consumoJ, ")")
		return (consumoI < consumoJ) // && i < j || consumoI < consumoJ
	})
}

// Stampa la pista che parte da Piastrella(x, y) e segue la sequenza di
// direzioni s, se tale pista è definita. Altrimenti non stampa nulla.
func pista(p piano, x int, y int, s string) {
	vertice := punto{x, y}
	seqDirezioni := strings.Split(s, ",")

	pistaDaStampare := ""
	calcolaPista(p, vertice, seqDirezioni, &pistaDaStampare)

	if pistaDaStampare != "" {
		fmt.Printf("[\n%s\n]\n", pistaDaStampare)
	}
}

// Determina la lunghezza della pista piu` breve che parte da Piastrella(x1, y1) e
// arriva in Piastrella(x2, y2).
// Altrimenti non stampa nulla.
func lung(p piano, x1 int, y1 int, x2 int, y2 int) {
	verticeOrig := punto{x1, y1}
	verticeDest := punto{x2, y2}

	lunghezza := calcolaPistaBreve(p, verticeOrig, verticeDest)

	if lunghezza > 0 {
		fmt.Println(lunghezza)
	}
}

//
/******************************** UTILITY FUNCTIONS **************************************/
//

// Funzione creata per effettuare un wrap di un operazione usata
// in numerosi punti del codice ed evitare ripetizioni.
func calcolaDeltaVertice(origine punto, deltaX int, deltaY int) (destinazione punto) {
	destinazione = punto{origine.x + deltaX, origine.y + deltaY}

	return destinazione
}

// DFS utilizzata dalle funzioni `blocco()`, `propagaBlocco()` per ottenere le piastrelle
// Utilizzo un check-nil per capire cosa calcolare.
// Non restituisce nulla, proprio per il fatto di essere versatile ed effettuare le modifiche
// sui valori originali passati per parametro quando necessario
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
		nuovoVertice := calcolaDeltaVertice(vertice, direzione.x, direzione.y)

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

// Funzione che calcola le piastrelle circonvicine alla Piastrella(x, y)
// Restituisce la mappa `vicine`
func piastrelleCirconvicine(p piano, vertice punto, colori map[string]int) (vicine map[punto]*Piastrella) {
	// TODO controllare se è più conveniente limitare le funzioni di questa e non
	// restituire un valore in base alla funzione ma magari splittarla in un
	// altra funzione e usare questa solo per la BFS
	vicine = make(map[punto]*Piastrella)

	// dirc := []punto{
	// 	{vertice.x - 1, vertice.y},
	// 	{vertice.x + 1, vertice.y},
	// 	{vertice.x, vertice.y - 1},
	// 	{vertice.x, vertice.y + 1},
	// 	{vertice.x - 1, vertice.y - 1},
	// 	{vertice.x + 1, vertice.y - 1},
	// 	{vertice.x - 1, vertice.y + 1},
	// 	{vertice.x + 1, vertice.y + 1}}
	for _, direzione := range direzioni {
		nuovoVertice := calcolaDeltaVertice(vertice, direzione.x, direzione.y)

		if piastrella, ok := p.piastrelle[nuovoVertice]; ok {
			vicine[punto{piastrella.x, piastrella.y}] = piastrella
			if colori != nil {
				colori[piastrella.colore]++
			}
		}
	}
	return vicine
}

// Verifica che una regola sia applicabile in base ai colori circostanti
// return: nil se nessuna regola è valida, altrimenti un puntatore alla regola
func verificaRegola(p piano, regola *Regola, vertice punto, intensità *int) *Regola {
	fmt.Println(" stampa regola: ", regola)
	valoriColore := make(map[string]int)
	piastrelleCirconvicine(p, vertice, valoriColore)

	// Verifico che l'insieme dei valori dei colori della regola sia minore
	// del valore delle piastrelle circonvicine, altrimenti non posso applicare la regola
	for colore, val := range regola.valColore {
		if valoriColore[colore] < val {
			return nil
		}
	}

	// Controllo se la piastrella esiste ed è accesa, in quel caso l'intensità sarà la stessa
	// della piastrella, altrimenti l'intensità sarà 1 di default
	piastrella, piastrellaOk := p.piastrelle[vertice]
	if piastrellaOk {
		*intensità = piastrella.intensità
	}

	return regola
}

// Calcola la pista seguendo le direzioni, restituisce nil se la pista non è definita
func calcolaPista(p piano, vertice punto, seqDirezioni []string, pistaDaStampare *string) {
	piastrella, ok := p.piastrelle[vertice]
	if !ok || piastrella.intensità == 0 {
		return
	}

	// Aggiungo la piastrella iniziale
	*pistaDaStampare = fmt.Sprintf("%d %d %s %d", piastrella.x, piastrella.y,
		piastrella.colore, piastrella.intensità)

	// seguo le direzioni per spostarmi nel piano e aggiorno la mappa
	for i := 0; i < len(seqDirezioni); i++ {
		deltaX := direzioni[seqDirezioni[i]].x
		deltaY := direzioni[seqDirezioni[i]].y
		vertice = calcolaDeltaVertice(vertice, deltaX, deltaY)

		if altraPiastrella, ok := p.piastrelle[vertice]; ok && altraPiastrella.intensità > 0 {
			*pistaDaStampare += fmt.Sprintf("\n%d %d %s %d",
				altraPiastrella.x, altraPiastrella.y,
				altraPiastrella.colore, altraPiastrella.intensità)
		} else {
			*pistaDaStampare = ""
			return
		}
	}
}

// Calcola la pista più breve, utilizza una `BFS` che aggiorna la mappa `pistaBreve`
// return: la lunghezza della pista più breve, 0 altrimenti
func calcolaPistaBreve(p piano, verticeOrig punto, verticeDest punto) int {
	piastrellaOrig, origineOk := p.piastrelle[verticeOrig]
	piastrellaDest, destOk := p.piastrelle[verticeDest]

	// Se le piastrelle di origine e destinazione non sono accese oppure non sono valide,
	//  non posso calcolare la pista
	if (!origineOk || piastrellaOrig.intensità == 0) || (!destOk || piastrellaDest.intensità == 0) {
		return 0
	}

	// Inizializzazione
	visitate := make(map[punto]bool)
	queue := []punto{verticeOrig}
	lunghezza := make(map[punto]int)
	// Aggiungo il primo vertice di origine alla mappa delle visite
	visitate[verticeOrig] = true
	lunghezza[verticeOrig] = 1

	/* TODO
		 * Documentare le modifiche fatte
		 * lunghezza ora è una mappa perché devo tenere conto del vertice precedente
		 *	A - B - C
		 *	lunghezza[A] = 0
		 *	lunghezza[B] = lunghezza[A] + 1
		 *	lunghezza[C] = lunghezza[B] + 1
	   *  NOTE : prima aggiornavo solo un intero generico
	*/

	// BFS
	for len(queue) > 0 {
		vertice := queue[0]
		queue = queue[1:]

		adiacenti := piastrelleCirconvicine(p, vertice, nil)
		for _, piastrella := range adiacenti {
			coordinatePiastrella := punto{piastrella.x, piastrella.y}
			if !visitate[coordinatePiastrella] {
				queue = append(queue, coordinatePiastrella) // aggiorno la coda
				visitate[coordinatePiastrella] = true       // segno la piastrella attuale come visitata
				// aggiorno la lunghezza -> lunghezza dal vertice nella coda +1
				lunghezza[coordinatePiastrella] = lunghezza[vertice] + 1

				if coordinatePiastrella == verticeDest { // se il vertice attuale è quello di arrivo, mi fermo
					return lunghezza[coordinatePiastrella]
				}
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
