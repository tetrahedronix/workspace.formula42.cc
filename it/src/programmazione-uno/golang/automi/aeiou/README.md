## Istruzioni

1. Il seguente comando compila il programma `aeiou` con il compilatore `go` e
produce l'eseguibile `aeiou`

```
go build aeiou.go

```

oppure per eseguire il programma `aeiou` senza costruire l'eseguibile:

```
go run aeiou.go

```

2. Per verificare la correttezza del programma, si consiglia di eseguire il
*testing* unitario utilizzando il seguente comando:

```
go test

```

## Comportamento atteso del programma `aeiou`

Il programma `aeiou` implementa un automa a stati finiti per verificare se
una parola contiene tutte e cinque le vocali ('a', 'e', 'i', 'o', 'u') in
ordine lessicografico. Ecco un riassunto del comportamento atteso del programma:

1. **Inizializzazione**: Il programma legge il contenuto di un file specificato
come argomento della riga di comando e lo converte in un buffer di byte.

2. **Analisi delle parole**: Le parole vengono estratte dal buffer e vengono
testate dall'automa implementato per determinare se contengono le vocali in
ordine.

3. **Automazione a stati**: L'automa ha uno stato iniziale e transizioni
definite per ogni carattere delle vocali. Le transizioni sono progettate per
accettare solo parole che contengono le vocali in ordine lessicografico.

4. **Verifica delle vocali**: L'automa inizia cercando la 'a' come prima
vocale, poi procede con 'e', 'i', 'o' e infine 'u'. Se tutte le vocali sono
trovate in ordine, la parola viene accettata.

5. **Restituzione del risultato**: Il programma determina se ogni parola nel
file soddisfa i criteri di accettazione dell'automa e stampa il risultato.

In sostanza, il programma serve a verificare se le parole contenute in un file
rispettano un certo schema di presenza delle vocali in ordine lessicografico.

## Esempi

Esempi di parole accettate dall'automa

1. `""`. La parola vuota ε non è accettata dall'automa. In una implementazione
di un automa deterministico per riconoscere le vocali in ordine lessicografico,
non è prevista l'accettazione della parola vuota ε. Questo è dovuto al fatto che
un automa deterministico opera su una serie di stati e transizioni definite,
senza ε-transizioni che consentono il passaggio diretto tra gli stati senza
input. Poiché la parola vuota non contiene alcun input, non esiste uno stato
di accettazione che possa riconoscerla. Di conseguenza, l'automa `aeiou` non
accetta tale parola e opera esclusivamente su stringhe non vuote contenenti
le vocali in ordine lessicografico. Nella sua implementazione viene aggiunto
uno *stato fittizio* che rappresenta una situazione di non-accettazione
raggiungibile dall'inizio solo se l'input è la parola vuota ε, senza 
violare la definizione di automa a stati finiti deterministico.

2. 