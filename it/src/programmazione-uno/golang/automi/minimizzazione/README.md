L'identificazione degli stati indistinguibili è alla base degli
algoritmi di minimizzazione degli automi a stati finiti.
La minimizzazione cerca di ridurre il numero di stati dell'automa
unendo quelli indistinguibili.

Per creare un programma Go per la minimizzazione degli automi, iniziamo
con una struttura dati che rappresenta la tabella di transizione dell'automa.
Una mappa può essere una buona scelta per questo scopo, poiché permette un
accesso efficiente alle transizioni di stato. Una volta creata questa struttura,
si possono applicare gli algoritmi di minimizzazione.

Più in generale una struttura dati per scopi di minimizzaione dovrebbe essere
definita come segue

1. Definizione della struttura dati

    - La struttura dati deve rappresentare gli stati, le transizioni e gli stati
    finali.
    - Per esempio, una mappa (o tabella hash) con chiavi come coppie (stato,
    simbolo) e valori come stati di destinazione.
    
2. Popolamento della tabella di transizione.

    - Si inseriscono tutte le transizioni dell'automa nella struttura dati.

3. Individuazione degli stati indistinguibili

    - Si usa l'algoritmo di distinzione degli stati (come l'algoritmo di
    Hopcroft o l'algoritmo di Moore) per identificare gli stati distinguibili.
    - Si crea una mappa delle classi di equivalenza per raggruppare gli stati
    indistinguibili insieme.

(Continua in Palano 14)

---

## Indice
