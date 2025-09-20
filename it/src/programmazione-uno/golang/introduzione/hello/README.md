## hello.go

### Contesto e richiamo teorico

Il linguaggio **Go** è un linguaggio general-purpose, fortemente tipizzato, con supporto per **garbage collection** e programmazione concorrente. I programmi sono organizzati in pacchetti e moduli, favorendo la gestione modulare delle dipendenze. Go è **compilato**, il che significa che il codice sorgente (`.go`) viene tradotto in un eseguibile tramite il compilatore e il linker, operazioni accessibili tramite il comando `go` con sottocomandi come `run` o `build`.

### Svolgimento dettagliato

Un programma Go consiste di uno o più **pacchetti**, identificati dalla dichiarazione iniziale `package`. Per un programma eseguibile si usa:

```go
package main
```

I pacchetti standard, come `fmt` per la gestione dell'input/output, si importano con:

```go
import "fmt"
```

La funzione **main()** rappresenta il punto d’ingresso del programma:

```go
func main() {
    var 세계 string = "Hello 세계! I'm Unicode"
    fmt.Println(세계)
}
```

Note operative:

* In Go i punti e virgola non sono necessari se ogni istruzione è su una riga diversa.
* Go supporta Unicode, quindi variabili e stringhe possono contenere caratteri internazionali.
* La funzione `fmt.Println()` stampa valori separati da spazi e termina con un ritorno a capo.

### Osservazioni

* Go impone uno **stile di formattazione rigido**, applicabile automaticamente tramite `gofmt`.
* La gestione degli **import** può essere semplificata con `goimports`, che aggiunge o rimuove pacchetti inutilizzati e formatta il codice.

### Applicazioni e spunti

Questo schema base è sufficiente per comprendere la struttura di un programma Go semplice, come il classico "Hello World", e costituisce la base per sviluppi più complessi in applicazioni Web, enterprise o strumenti a linea di comando.

