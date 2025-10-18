# Spazi vettoriali

TensorFlow consente la manipolazione di spazi vettoriali in modo naturale, poiché i suoi oggetti fondamentali sono i tensori. Un tensore è un ente geometrico che costituisce una generalizzazione del vettore in uno spazio vettoriale. Esso possiede un determinato numero di indici. Quelli più comuni hanno due indici che, in un determinato sistema di riferimento, vengono individuati dalle componenti di una matrice n×n. In generale, un tensore con `k` indici è individuato tramite `n^k` componenti. È possibile poi applicare una legge di trasformazione per passare da un sistema di riferimento a un altro.

Un tensore supporta le operazioni lineari più comuni degli spazi vettoriali: somma di vettori, prodotto scalare, moltiplicazione per scalare, prodotto matriciale e trasformazioni lineari. 

Tramite il binding Go di TensorFlow si possono implementare queste operazioni su array multidimensionali di dimensione arbitraria. In particolare, tramite il pacchetto `tensorflow/op` è possibile definire e manipolare sottospazi, basi, trasformazioni lineari e operazioni come il calcolo di autovalori, autovettori o decomposizioni ortogonali, lavorando in modo simile a quanto avviene usando una libreria di algebra lineare come OpenBLAS.

TensorFlow in Go permette non solo di lavorare con spazi vettoriali, ma anche di eseguire automaticamente il calcolo delle derivate (cioè, senza scrivere esplicitamente le formule del calcolo differenziale) attraverso grafi computazionali. Questo lo rende adatto sia per l’algebra lineare, sia per il machine learning e il deep learning

## Indice

1. [`Scalare`](./scalare/) propone un semplice operazione di moltiplicazione scalare quando lo spazio vettoriale è un campo su se stesso.
2. [`Matrice`](./matrice/) usa il binding TensorFlow per creare tensori da matrici native (`[][]float32`) con `tf.NewTensor`.



