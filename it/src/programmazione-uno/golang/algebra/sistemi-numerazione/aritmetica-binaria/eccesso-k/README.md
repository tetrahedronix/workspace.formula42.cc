# Codici in eccesso-k

I codici in eccesso-k (in inglese "excess-k codes") sono una particolare classe
di codici numerici utilizzati principalmente per la rappresentazione di numeri
in sistemi digitali e nelle comunicazioni.

In un codice in eccesso-k, ogni numero decimale `N` viene rappresentato sommando
il valore `k` alla sua rappresentazione binaria naturale su `n` bit. 
Si converte `(N+k)` in un intero senza segno su `n` bit. L'intervallo di
rappresentazione varia da `-k` a `(2^n)-k-1`, estremi inclusi.

## Applicazioni

Il codice in eccesso-k è utilizzato nella rappresentazione floating point dei
numeri reali, in particolare nell'esponente dei formati IEEE 754, lo standard
più diffuso per rappresentare numeri reali nei computer. 
Secondo la specifica IEEE 754, l'esponente viene memorizzato utilizzando una
rappresentazione in eccesso-k, dove k è spesso chiamato *bias* (polarizzazione).
I valori specifici di `k` dipendono però dal formato:

* Per i numeri a precisione singola (32 bit), l'esponente è rappresentato su
8 bit. Il bias è 127, quindi si usa un codice in eccesso-127. L'intervallo
dell'esponente varia da `-127` a `(2^8)-128=128`, tuttavia l'intervallo
effettivo per i numeri normalizzati in IEEE 754 diventa `[-126,127]`, in
quanto il valore `-127` viene "perso" per rappresentare i denormalizzati, e
il valore `128` viene usato per infinito e `NaN`: quando tutti i bit sono a 1
`(11111111)2`, cioè 255, si rappresentano valori speciali, in particolare 
l'esponente zero e i valori riservati per denormalizzati e infiniti.

* Per i numeri a precisione doppia (64 bit), l'esponente è rappresentato su 11
bit. Il bias è 1023, quindi si usa un codice in eccesso-1023. L'intervallo
dell'esponente varia da `-1022` a `+1023`, cioè `[-1022,1023]`. Anche in questo
caso due rappresentazioni sono riservate per i dernomalizzati e lo zero rispetto 
all'intervallo di variabilità effettivo di `[-1023,1024]`.

## Esempio

La rappresentazione dell'esponente -5 in un formato a precisione singola su 8
bit (e bias 127) si ottiene calcolando `-5+127=122` che codificato in binario
senza segno è `01111010`.

## Vantaggi

La rappresentazione in eccesso-k presenta diversi vantaggi:

1. Permette di rappresentare sia esponenti positivi che negativi usando solo bit
non segnati.
2. Semplifica i confronti tra numeri floating point, poiché numeri con esponenti
maggiori sono sempre maggiori (in valore assoluto).
3. Facilita l'implementazione hardware delle operazioni aritmetiche.

