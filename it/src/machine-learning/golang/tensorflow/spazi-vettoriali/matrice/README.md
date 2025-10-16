## Contesto e richiamo teorico

Questo esempio mostra come creare e ispezionare tensori TensorFlow direttamente in Go, utilizzando il binding fornito dal pacchetto `github.com/wamuir/graft/tensorflow`. Il programma illustra il legame tra i tipi nativi Go (come `[][]float32`) e la rappresentazione interna dei tensori nella libreria TensorFlow, con deduzione automatica di forma e tipo.

## Analisi del codice

Dopo aver definito due matrici in Go, si costruiscono i corrispondenti tensori mediante il costruttore `tf.NewTensor`. Tale funzione analizza il valore fornito e genera un tensore di forma e tipo coerenti con i dati di input. In seguito, il codice stampa i valori, la forma e il tipo del tensore risultante.

> ```go
> MatrixA := [][]float32{
> 	{2.0, 1.0, 5.0},
> 	{1.0, 3.0, 2.0},
> }
> 
> MatrixB := [][]float32{
> 	{3.0, 4.0},
> 	{-1.0, 2.0},
> 	{2.0, 1.0},
> }
> 
> tensorA, err := tf.NewTensor(MatrixA)
> tensorB, err := tf.NewTensor(MatrixB)
> 
> fmt.Println("Tensor A:")
> fmt.Println(tensorA.Value())
> fmt.Println("Shape:", tensorA.Shape())
> fmt.Println("DType:", tensorA.DataType())
> 
> fmt.Println("Tensor B:")
> fmt.Println(tensorB.Value())
> fmt.Println("Shape:", tensorB.Shape())
> fmt.Println("DType:", tensorB.DataType())
> ```
> 
> (*Il frammento mostra la creazione dei tensori a partire da matrici Go e la successiva ispezione di valore, forma e tipo di dato.*)

Il metodo `NewTensor(value any)` è il costruttore principale: crea un `tf.Tensor` deducendo forma e tipo dai dati forniti. Il tipo risultante (`DType`) è una costante enumerativa interna che mappa i tipi di Go ai corrispondenti tipi TensorFlow, ad esempio `TF_FLOAT`, `TF_DOUBLE`, `TF_INT32` e così via.

> ```go
> const (
> 	Float  DataType = C.TF_FLOAT
> 	Double DataType = C.TF_DOUBLE
> 	Int32  DataType = C.TF_INT32
> 	Bool   DataType = C.TF_BOOL
> 	...
> )
> ```
> 
> (*Il frammento mostra un estratto dell’enumerazione dei tipi supportati nel binding Go.*)

## Osservazioni

Il valore di un tensore (`Value()`) viene convertito nel corrispondente valore Go. Tuttavia, non tutti i tipi TensorFlow sono attualmente supportati nel binding, e la conversione di tipi non gestiti può generare un panic. È pertanto consigliabile utilizzare tipi primitivi comuni (`float32`, `float64`, `int32`) per garantire stabilità e compatibilità.

## Applicazioni e spunti

Questo approccio consente di integrare TensorFlow in pipeline Go, mantenendo il controllo tipico del linguaggio e la possibilità di costruire dinamicamente grafi computazionali. Costituisce la base per sistemi di calcolo numerico, addestramento o inferenza di modelli ML direttamente all’interno di un ecosistema Go.
