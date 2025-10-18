# Spazi vettoriali

TensorFlow consente la manipolazione di spazi vettoriali in modo naturale, poiché i suoi oggetti fondamentali sono i tensori. Un [tensore](https://www.treccani.it/enciclopedia/tensore_(Enciclopedia-della-Matematica)/) è un ente geometrico che costituisce una generalizzazione del vettore in uno [spazio vettoriale](https://www.treccani.it/enciclopedia/spazio-vettoriale_(Enciclopedia-della-Matematica)/). Esso possiede un determinato numero di indici. Quelli più comuni hanno due indici che, in un determinato sistema di riferimento, vengono individuati dalle componenti di una matrice `n×n`. In generale, un tensore con `k` indici è individuato tramite `n^k` componenti. È possibile poi applicare una legge di trasformazione per passare da un sistema di riferimento a un altro.

Un tensore supporta le operazioni lineari più comuni degli spazi vettoriali: somma di vettori, prodotto scalare, moltiplicazione per scalare, prodotto matriciale e trasformazioni lineari. 

Tramite il [binding Go](https://pkg.go.dev/github.com/wamuir/graft@v0.10.0/tensorflow#pkg-overview) di TensorFlow si possono implementare queste operazioni su array multidimensionali di dimensione arbitraria. In particolare, tramite il pacchetto `tensorflow/op` è possibile definire e manipolare sottospazi, basi, trasformazioni lineari e operazioni come il calcolo di autovalori, autovettori o decomposizioni ortogonali, lavorando in modo simile a quanto avviene usando una libreria di algebra lineare come OpenBLAS.

TensorFlow in Go permette non solo di lavorare con spazi vettoriali, ma anche di eseguire automaticamente il calcolo delle derivate (cioè, senza scrivere esplicitamente le formule del calcolo differenziale) attraverso grafi computazionali. Questo lo rende adatto sia per l’algebra lineare, sia per il machine learning e il deep learning

## Indice

1. [`Scalare`](./scalare/) propone un semplice operazione di moltiplicazione scalare quando lo spazio vettoriale è un campo su se stesso.
2. [`Matrice`](./matrice/) usa il binding TensorFlow per creare tensori da matrici native (`[][]float32`) con `tf.NewTensor`.

## Note implementative

Un vettore può essere considerato un tensore di ordine 1, una matrice un tensore di ordine 2, e più in generale un tensore di ordine `n` è un oggetto che può essere rappresentato come una collezione di numeri indicizzati da `n` indici.

TensorFlow (come molte librerie di calcolo numerico) utilizza termini ispirati all’algebra lineare, ma li riassegna in un contesto computazionale. In questo senso, il termine **rank** indica l’ordine del tensore, cioè il numero di assi (dimensioni) della struttura dati.
Si ha quindi che:

* uno scalare ha rank 0
* un vettore ha rank 1
* una matrice ha rank 2
* un tensore 3D ha rank 3

In algebra lineare, invece, il **rango** di una matrice rappresenta la dimensione dello spazio vettoriale generato dalle sue righe o colonne, cioè il numero massimo di vettori linearmente indipendenti. I due concetti, dunque, non coincidono: il primo è *strutturale* (numero di assi del contenitore), il secondo è *algebrico* (dimensione del sottospazio generato).

Nel binding Go, una slice `[][]float32` rappresenta una matrice (tensore di rango 2 in senso TensorFlow), ma la funzione `Shape()` restituisce le dimensioni strutturali, non la dimensione vettoriale né il rango algebrico.

Come specificato anche nella documentazione ufficiale del binding di TensorFlow:

> **Note:** The rank of a tensor is not the same as the rank of a matrix. The rank of a tensor is the number of indices required to uniquely select each element of the tensor. Rank is also known as “order,” “degree,” or “ndims.”

