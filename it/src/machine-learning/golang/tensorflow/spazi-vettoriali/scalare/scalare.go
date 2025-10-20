package main

/*
#cgo LDFLAGS: -L/usr/local/lib -ltensorflow -ltensorflow_framework -Wl,-rpath,/usr/local/lib
*/
import "C"
import (
	"fmt"

	tf "github.com/wamuir/graft/tensorflow"
	"github.com/wamuir/graft/tensorflow/op"
)

func main() {

	// Si definiscono due scalari Alfa, Beta reali
	// Uno scalare è un elemento di un campo K, in questo caso il campo dei
	// numeri reali e il campo Z2. I numeri reali sono rappresentati come
	// valori floating point a 32 bit; per questioni di semplicità il campo Z2 è
	// invece rappresentato con un int32, ma le sue operazioni non sono definite.
	var (
		ScalareAlfa float32 = 3.8
		ScalareBeta int32   = 1
	)

	// Si crea un tensore chiamando il costruttore NewTensor() della libreria
	// TensorFlow, e passando come argomento il dato da associargli, in questo
	// caso il numero scalare contenuto nella variabile ScalareAlfa
	tensorA, err := tf.NewTensor(ScalareAlfa)

	if err != nil {
		panic(err)
	}

	// Si crea un altro tensore per lo scalare ScalareBeta
	tensorB, err := tf.NewTensor(ScalareBeta)

	if err != nil {
		panic(err)
	}

	// Semplici istruzioni di debug: si stampano gli attributi del tensore:
	// * Il suo tipo, cioè 1 corrispondente a FLOAT32
	// * Il valore, cioè il numero reale rappresentato in floating point
	// * La sua forma, corrispondente alla dimensione della rappresentazione interna
	// * Il rank, ossia il numero di indici
	fmt.Println("Scalare A")
	fmt.Println("Tipo:", tensorA.DataType())
	fmt.Println("Valore:", tensorA.Value())
	fmt.Println("Forma:", tensorA.Shape())
	fmt.Println("Rank:", len(tensorA.Shape()))
	// Stesse istruzioni di debug per lo scalare Beta
	fmt.Println("Scalare B")
	fmt.Println("Tipo:", tensorB.DataType())
	fmt.Println("Valore:", tensorB.Value())
	fmt.Println("Forma:", tensorB.Shape())
	fmt.Println("Rank:", len(tensorB.Shape()))

	// Esempio di utilizzo della funzione MoltPerScalare
	productTensor, err := MoltPerScalare(tensorA, tensorA)
	if err != nil {
		panic(err)
	}
	fmt.Println("Prodotto Scalare (Tensore): A×A", productTensor.Value())
}

// MoltPerScalare calcola il prodotto scalare di due tensori scalari.
// Accetta due tensori (*tf.Tensor) e restituisce un nuovo tensore
// che rappresenta il loro prodotto.
func MoltPerScalare(a, b *tf.Tensor) (*tf.Tensor, error) {
	// Crea un nuovo grafo per l'operazione
	g := tf.NewGraph()
	if g == nil { // Added check for NewGraph
		return nil, fmt.Errorf("tf.NewGraph() ha restituito un grafo nullo")
	}
	// Definisce le operazioni nel grafo
	s := op.NewScopeWithGraph(g) // s is the main scope

	// Create sub-scopes for unique naming of Const operations
	scopeA := s.SubScope("ConstA")
	scopeB := s.SubScope("ConstB")
	scopeMul := s.SubScope("MulProduct") // Also for Mul

	// Converte *tf.Tensor in tf.Output usando op.Const all'interno dei sub-scope
	outputA := op.Const(scopeA, a)
	outputB := op.Const(scopeB, b)

	// Esegue la moltiplicazione usando op.Mul all'interno del sub-scope
	product := op.Mul(scopeMul, outputA, outputB)

	// Aggiunge un controllo per gli errori accumulati nello scope
	if err := s.Err(); err != nil {
		return nil, fmt.Errorf("errore durante la costruzione del grafo nello scope: %w", err)
	}

	// Finalizza lo scope per ottenere un grafo
	options := &tf.SessionOptions{}
	sess, err := tf.NewSession(g, options)
	if err != nil {
		return nil, fmt.Errorf("errore nella creazione della sessione: %w", err)
	}
	if sess == nil { // Added check for NewSession
		return nil, fmt.Errorf("tf.NewSession() ha restituito una sessione nulla senza errore")
	}
	defer sess.Close()

	// Esegue il grafo
	// Inizializza una mappa vuota per i feed e una slice vuota per i target
	results, err := sess.Run(nil, []tf.Output{product}, nil)
	if err != nil {
		return nil, fmt.Errorf("errore nell'esecuzione della sessione: %w", err)
	}

	if len(results) == 0 {
		return nil, fmt.Errorf("nessun risultato ottenuto dall'esecuzione della sessione")
	}

	return results[0], nil
}

/*
## Note aggiuntive

### Prodotto per Scalare in TensorFlow con Graft

Questa implementazione della funzione `MoltPerScalare` dimostra come eseguire
un'operazione di prodotto per scalare (moltiplicazione elemento per elemento
tra tensori scalari) utilizzando il binding Go di TensorFlow, `graft`.

#### Configurazione del Grafo

TensorFlow opera su un grafo computazionale. Per definire le operazioni,
seguiamo questi passaggi:

1.  **Creazione del Grafo Principale (`*tf.Graph`)**: Si inizia con
`tf.NewGraph()`. Questo oggetto rappresenta la struttura del grafo su cui
verranno aggiunte le operazioni.

2.  **Creazione dello Scope (`*op.Scope`)**: `op.NewScopeWithGraph(g)`
crea uno scope associato al grafo principale. Lo scope è un costruttore di
operazioni che facilita l'aggiunta di nodi al grafo.

3.  **Aggiunta di Operazioni (`op.Const`, `op.Mul`)**:

    *   `op.Const(scope, tensor)`: Converte un `*tf.Tensor` Go (un tensore
		concreto) in un'operazione costante nel grafo TensorFlow, restituendo
		un `tf.Output` (un riferimento simbolico al risultato dell'operazione).

    *   `op.Mul(scope, outputX, outputY)`: Definisce un'operazione di
		 moltiplicazione tra due `tf.Output` nel grafo.

    *   **Nomi Univoci**: Per evitare conflitti di nomi nel grafo, è
		fondamentale garantire che ogni operazione abbia un nome univoco.
		Utilizziamo `s.SubScope("NomeUnico")` per creare sottoscopi con nomi
		specifici, garantendo che le operazioni all'interno di tali sottoscopi
		siano automaticamente namespaced.

4.  **Controllo Errori dello Scope (`s.Err()`)**: Le funzioni `op.*` non
	restituiscono errori direttamente, ma li registrano nello scope. È
	cruciale controllare `s.Err()` dopo aver aggiunto tutte le operazioni per
	rilevare eventuali problemi durante la costruzione del grafo.

5.  **Finalizzazione del Grafo (`s.Finalize()`)**: Dopo aver definito tutte le
	operazioni, `s.Finalize()` converte lo scope in un `*tf.Graph` eseguibile.

#### Esecuzione del Grafo con una Sessione

Una volta configurato il grafo, è necessario eseguirlo tramite una sessione:

1.  **Creazione della Sessione (`*tf.Session`)**:
	`tf.NewSession(graph, options)` crea una sessione per eseguire il grafo.
	`options` (`*tf.SessionOptions`) permette configurazioni avanzate del
	runtime di TensorFlow (es. target, configurazione GPU). Se non specificato,
	si usano le opzioni predefinite.

2.  **Esecuzione delle Operazioni (`sess.Run()`)**:

    *   `sess.Run(feeds, fetches, targets)`: Questo è il metodo principale per
		eseguire il grafo.

    *   `feeds` (`map[tf.Output]*tf.Tensor`): Utilizzato per alimentare valori concreti a placeholder nel grafo. In questo caso, non
		avendo placeholder, viene passata una mappa vuota.

    *   `fetches` (`[]tf.Output`): Una slice di `tf.Output` i cui valori devono
		essere recuperati dopo l'esecuzione. Qui recuperiamo il risultato del
		prodotto.

    *   `targets` (`[]*tf.Operation`): Una slice di operazioni da eseguire
		senza recuperarne l'output. In questo caso, non avendo operazioni
		specifiche da eseguire in questo modo, viene passata una slice vuota.

3.  **Chiusura della Sessione (`defer sess.Close()`)**: È fondamentale chiudere
	la sessione per rilasciare le risorse. `defer` assicura che venga chiamata
	alla fine della funzione.

#### Rappresentazione dei Tensori Scalari

Nel binding Go di TensorFlow (`graft`), un tensore scalare è rappresentato da
un `*tf.Tensor` la cui `Shape` è una slice vuota (`[]int64{}`) o un array di
dimensione 0. Di conseguenza, il suo `Rank` (ottenuto tramite
`len(tensor.Shape())`) è 0. Questo riflette il fatto che uno scalare non ha
dimensioni o assi, a differenza di vettori o matrici.

#### Gestione degli Errori

È buona pratica includere controlli degli errori a ogni passaggio critico
(creazione del grafo, costruzione dello scope, finalizzazione, creazione
della sessione, esecuzione) per garantire la robustezza del programma.



*/
