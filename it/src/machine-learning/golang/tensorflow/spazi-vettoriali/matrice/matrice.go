package main

/*
#cgo LDFLAGS: -L/usr/local/lib -ltensorflow -ltensorflow_framework -Wl,-rpath,/usr/local/lib
*/
import "C"

import (
	"fmt"

	tf "github.com/wamuir/graft/tensorflow"
	// Il pacchetto op definisce le funzioni per aggiungere operazioni TensorFlow
	// a un grafo di tipo Graph
	// "github.com/wamuir/graft/tensorflow/op"
)

func main() {
	// Si definiscono i dati come una slice di slice (matrice) in Go.
	// Può essere array, slice, int, float32, ecc.
	MatrixA := [][]float32{
		{2.0, 1.0, 5.0},
		{1.0, 3.0, 2.0},
	}

	MatrixB := [][]float32{
		{3.0, 4.0},
		{-1.0, 2.0},
		{2.0, 1.0},
	}

	// Per creare un tensore Tensor occorre chiamare la funzione (il costruttore)
	// NewTensor() e passare come argomento una slice o un array multidimensionale
	// che rappresenta la matrice.
	// NewTensor() deduce automaticamente la forma (2x3) e il tipo (float32) dal
	// contesto
	tensorA, err := tf.NewTensor(MatrixA)
	tensorB, err := tf.NewTensor(MatrixB)

	if err != nil {
		panic(err)
	}

	// Il binding Go di TensorFlow deduce la forma (Shape) e il tipo di dato
	// (DType) del Tensor direttamente dal tipo e dalla struttura del valore Go
	// fornito.
	fmt.Println("Tensor A:")
	fmt.Println(tensorA.Value())
	fmt.Println("Shape:", tensorA.Shape())
	fmt.Println("DType:", tensorA.DataType())

	fmt.Println("Tensor B:")
	fmt.Println(tensorB.Value())
	fmt.Println("Shape:", tensorB.Shape())
	fmt.Println("DType:", tensorB.DataType())

}

/*
## Note aggiuntive

### Il costruttore

`tf.NewTensor(value any)`` è la funzione principale. Accetta un valore Go di qualsiasi
tipo (`any`). Se il valore è una struttura dati compatibile (come array, slice,
int, ecc.) NewTensor crea un tf.Tensor con la forma e il tipo di dato (DType)
dedotti dalla struttura del valore Go.

### Le strutture dati

Per creare una matrice si definisce una variabile di tipo `[][]float32` che
corrisponde a una matrice 2x3 in cui ogni elemento è un `float32`. Occorre
prestare attenzione al tipo: DataType nel binding Go è una costante enumerativa,
e i tipi ammessi sono i seguenti

const (
	Float        DataType = C.TF_FLOAT
	Double       DataType = C.TF_DOUBLE
	Int32        DataType = C.TF_INT32
	Uint32       DataType = C.TF_UINT32
	Uint8        DataType = C.TF_UINT8
	Int16        DataType = C.TF_INT16
	Int8         DataType = C.TF_INT8
	String       DataType = C.TF_STRING
	Complex64    DataType = C.TF_COMPLEX64
	Complex      DataType = C.TF_COMPLEX
	Int64        DataType = C.TF_INT64
	Uint64       DataType = C.TF_UINT64
	Bool         DataType = C.TF_BOOL
	Qint8        DataType = C.TF_QINT8
	Quint8       DataType = C.TF_QUINT8
	Qint32       DataType = C.TF_QINT32
	Bfloat16     DataType = C.TF_BFLOAT16
	Qint16       DataType = C.TF_QINT16
	Quint16      DataType = C.TF_QUINT16
	Uint16       DataType = C.TF_UINT16
	Complex128   DataType = C.TF_COMPLEX128
	Half         DataType = C.TF_HALF
	Float8e5m2   DataType = C.TF_FLOAT8_E5M2
	Float8e4m3fn DataType = C.TF_FLOAT8_E4M3FN
	Int4         DataType = C.TF_INT4
	Uint4        DataType = C.TF_UINT4
)

Il metodo Value del tensore converte il tipo sottostante C associato al tensore
al corrispondente valore Go. Per ora non tutti i tipi del tensore sono
supportati nel binding Go, pertanto questa funzione potrebbe determinare panic
se viene incontra un DataType non supportato.

*/
