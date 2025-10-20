# Contesto del Progetto per Gemini

Questo documento fornisce una panoramica del progetto per l'agente Gemini, facilitando la comprensione e l'interazione con la codebase.

## Panoramica del Progetto

Questo è un progetto Go (`go 1.24.5`) focalizzato sull'apprendimento automatico (Machine Learning), in particolare sull'esplorazione e l'implementazione di concetti di algebra lineare come spazi vettoriali, vettori e matrici. Il progetto utilizza i binding Go per TensorFlow, in particolare la libreria `github.com/wamuir/graft v0.10.0`, per manipolare tensori che generalizzano i concetti di vettori e matrici.

La struttura del progetto suggerisce un'organizzazione modulare, con una sottodirectory `tensorflow` dedicata a esercizi e implementazioni relative agli spazi vettoriali e all'uso di TensorFlow in questo contesto.

**Tecnologie Principali:**
*   **Linguaggio:** Go (versione 1.24.5)
*   **Machine Learning:** TensorFlow (tramite binding Go)
*   **Librerie Go:** `github.com/wamuir/graft` (binding TensorFlow), `google.golang.org/protobuf` (indiretto)

## Compilazione ed Esecuzione

Essendo un progetto Go, i comandi standard di Go dovrebbero essere utilizzati per la compilazione e l'esecuzione.

*   **Compilazione:** Per compilare il modulo principale o un'applicazione specifica, navigare nella directory contenente il file `main.go` (o il modulo desiderato) ed eseguire:
    ```bash
    go build .
    ```
*   **Esecuzione:** Per eseguire un'applicazione Go, navigare nella directory appropriata ed eseguire:
    ```bash
    go run .
    ```
*   **Test:** Per eseguire i test (se presenti), navigare nella directory del modulo ed eseguire:
    ```bash
    go test ./...
    ```

Al momento, non sono stati identificati file `main.go` specifici o test. Si consiglia di esplorare la sottodirectory `tensorflow` per esempi o applicazioni specifiche.

## Convenzioni di Sviluppo

Il progetto segue le convenzioni standard di Go. Si raccomanda di aderire alle seguenti pratiche:

*   **Formattazione del Codice:** Utilizzare `gofmt` per mantenere uno stile di codice coerente.
    ```bash
    go fmt ./...
    ```
*   **Linting:** Utilizzare strumenti di linting Go (come `golangci-lint` se configurato) per identificare potenziali problemi.
*   **Struttura del Progetto:** Mantenere una chiara separazione delle responsabilità tra i pacchetti e le directory.
*   **Documentazione:** Scrivere commenti chiari e concisi per funzioni, tipi e variabili esportate.
*   **Gestione delle Dipendenze:** Le dipendenze sono gestite tramite `go.mod` e `go.sum`.
    ```bash
    go mod tidy
    ```
