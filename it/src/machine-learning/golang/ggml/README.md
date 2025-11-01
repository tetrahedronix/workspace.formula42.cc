# GGML

[![zread](https://img.shields.io/badge/Ask_Zread-_.svg?style=flat&color=00b0aa&labelColor=000000&logo=data%3Aimage%2Fsvg%2Bxml%3Bbase64%2CPHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTYiIHZpZXdCb3g9IjAgMCAxNiAxNiIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTQuOTYxNTYgMS42MDAxSDIuMjQxNTZDMS44ODgxIDEuNjAwMSAxLjYwMTU2IDEuODg2NjQgMS42MDE1NiAyLjI0MDFWNC45NjAxQzEuNjAxNTYgNS4zMTM1NiAxLjg4ODEgNS42MDAxIDIuMjQxNTYgNS42MDAxSDQuOTYxNTZDNS4zMTUwMiA1LjYwMDEgNS42MDE1NiA1LjMxMzU2IDUuNjAxNTYgNC45NjAxVjIuMjQwMUM1LjYwMTU2IDEuODg2NjQgNS4zMTUwMiAxLjYwMDEgNC45NjE1NiAxLjYwMDFaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00Ljk2MTU2IDEwLjM5OTlIMi4yNDE1NkMxLjg4ODEgMTAuMzk5OSAxLjYwMTU2IDEwLjY4NjQgMS42MDE1NiAxMS4wMzk5VjEzLjc1OTlDMS42MDE1NiAxNC4xMTM0IDEuODg4MSAxNC4zOTk5IDIuMjQxNTYgMTQuMzk5OUg0Ljk2MTU2QzUuMzE1MDIgMTQuMzk5OSA1LjYwMTU2IDE0LjExMzQgNS42MDE1NiAxMy43NTk5VjExLjAzOTlDNS42MDE1NiAxMC42ODY0IDUuMzE1MDIgMTAuMzk5OSA0Ljk2MTU2IDEwLjM5OTlaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik0xMy43NTg0IDEuNjAwMUgxMS4wMzg0QzEwLjY4NSAxLjYwMDEgMTAuMzk4NCAxLjg4NjY0IDEwLjM5ODQgMi4yNDAxVjQuOTYwMUMxMC4zOTg0IDUuMzEzNTYgMTAuNjg1IDUuNjAwMSAxMS4wMzg0IDUuNjAwMUgxMy43NTg0QzE0LjExMTkgNS42MDAxIDE0LjM5ODQgNS4zMTM1NiAxNC4zOTg0IDQuOTYwMVYyLjI0MDFDMTQuMzk4NCAxLjg4NjY0IDE0LjExMTkgMS42MDAxIDEzLjc1ODQgMS42MDAxWiIgZmlsbD0iI2ZmZiIvPgo8cGF0aCBkPSJNNCAxMkwxMiA0TDQgMTJaIiBmaWxsPSIjZmZmIi8%2BCjxwYXRoIGQ9Ik00IDEyTDEyIDQiIHN0cm9rZT0iI2ZmZiIgc3Ryb2tlLXdpZHRoPSIxLjUiIHN0cm9rZS1saW5lY2FwPSJyb3VuZCIvPgo8L3N2Zz4K&logoColor=ffffff)](https://zread.ai/ggml-org/ggml)

Dalla documentazione [ggml](https://huggingface.co/blog/introduction-to-ggml).

La libreria [ggml](https://github.com/ggerganov/ggml) è un framework open source per il machine learning scritto in C e C++. È concepita come un'alternativa leggera e minimale a librerie più complesse come PyTorch e Tensorflow. 

La sua popolarità è cresciuta rapidamente grazie all'utilizzo in progetti molto noti come **llama.cpp** e **whisper.cpp**. Ggml si distingue per alcune caratteristiche fondamentali che ne fanno una scelta ideale per applicazioni portabili ad alte prestazioni:

* **Minimalismo** - Il nucleo della libreria è composto da soli cinque file sorgente che implementano le funzionalità essenziali. Il supporto GPU è opzionale e può essere aggiunto tramite moduli separati.
* **Compilazione semplificata**: Non richiede strumenti complessi o dipendenze esterne. Ad esempio, senza il supporto GPU, è sufficiente un compilatore standard come GCC o Clang.
* **Leggerezza** - Il binario risultante occupa circa 1MB su disco, un ordine di grandezza inferiore rispetto a centinaia di megabyte richiesti da librerie come PyTorch.
* **Ampia compatibilità** - Supporta le principali architetture hardware, incluse X86_64, ARM, Apple Silicon e GPU NVIDIA.
* **Quantizzazione dei tensori** - I tensori possono essere quantizzati per ridurre l'uso di memoria a runtime, con un approccio simile alla compressione JPEG. In molti casi, questa tecnica consente un aumento della velocità di esecuzione.
* **Efficienza** - L'overhead associato alla gestione dei tensori e all'esecuzione delle operazioni matematiche è estremamente ridotto, garantendo buone prestazioni anche su sistemi con risorse limitate.

È importante tenere presenti alcuni limiti e potenziali difficoltà nell'utilizzo di ggml rispetto ad altre librerie di machine learning:

* **Supporto parziale delle operazioni** - Non tutte le operazioni sui tensori sono disponibili per ogni backend. Ad esempio, alcune funzioni possono essere eseguite su CPU ma non ancora su CUDA.
* **Maggiore complessità di sviluppo** - L'utilizzo di ggml può risultare più impegnativo rispetto ad altri framework, poiché richiede una comprensione più approfondita della programmazione a basso livello e della gestione esplicita della memoria.
* **Evoluzione continua del progetto** - Essendo una libreria in rapido sviluppo, il codice subisce aggiornamenti frequenti che talvolta possono introdurre modifiche incompatibili con le versioni precedenti.

GGML è una libreria di basso livello. È fondamentale capire i concetti e i termini che la caratterizzano per avere un maggior controllo sulle prestazioni.

## Glossario

### ggml_context

È una struttura che funziona da "contenitore" per gli oggetti principali gestiti da gglm, come tensori, grafi e, opzionalmente, i relativi dati. 

>```c++
>     struct ggml_context;
>```

### ggml_cgraph

Rappresenta un grafo computazionale, cioè una sequenza ordinata di operazioni da eseguire. In pratica, definisce il flusso di calcolo che verrà passato al backend (per esempio CPU o GPU) per l'esecuzione.

>```c++
> // computation graph
> struct ggml_cgraph {
>   int size;
>   int n_nodes;
>   int n_leafs;
> 
>   struct ggml_tensor ** nodes;
>   struct ggml_tensor ** grads;
>   struct ggml_tensor ** leafs;
> 
>   struct ggml_hash_set visited_hash_set;
> 
>   enum ggml_cgraph_eval_order order;
> };
>```

### ggml_backend

È l'interfaccia responsabile dell'esecuzione dei grafi computazionali. Ggml supporta diversi backend, tra cui CPU (predefinito), CUDA, Metal (per Apple Silicon), Vulkan, RPC e altri.

>```c++
> struct ggml_backend {
>    ggml_guid_t guid;
> 
>    struct ggml_backend_i iface;
>    ggml_backend_context_t context;
> };
>```

### ggml_backend_buffer_type

Definisce un tipo di buffer di memoria associato a un determinato backend. Può essere considerato come un allocatore di memoria specifico per un ambiente d'esecuzione. Ad esempio, per eseguire calcoli su GPU è necessario allocare la memoria sulla GPU stessa tramite il relativo `buffer_type` (spesso abbreviato come `buft`)

>```c++
> struct ggml_backend_buffer_type_i {
>         const char *          (*GGML_CALL get_name)        (ggml_backend_buffer_type_t buft);
>         // allocate a buffer of this type
>         ggml_backend_buffer_t (*GGML_CALL alloc_buffer)    (ggml_backend_buffer_type_t buft, size_t size);
>         // tensor alignment
>         size_t                (*GGML_CALL get_alignment)   (ggml_backend_buffer_type_t buft);
>         // max buffer size that can be allocated
>         size_t                (*GGML_CALL get_max_size)    (ggml_backend_buffer_type_t buft);
>         // data size needed to allocate the tensor, including padding
>         size_t                (*GGML_CALL get_alloc_size)  (ggml_backend_buffer_type_t buft, const struct ggml_tensor * tensor);
>         // check if tensor data is in host memory
>         bool                  (*GGML_CALL is_host)         (ggml_backend_buffer_type_t buft);
>     };
>```

* `ggml_backend_buffer_type`: rappresenta un tipo di buffer. Si può pensare come un allocatore di memoria connesso a ciascun `ggml_backend`. Per esempio, se si vuole compiere calcoli su GPU, occorre allocare memoria sulla GPU per mezzo di `buffer_type` (si solito abbreviato con `buft`).

### ggml_backend_buffer

Rappresenta un buffer di memoria allocato e gestito da un backend specifico (es. CPU, GPU). Contiene i dati dei tensori e altre informazioni necessarie per le operazioni di calcolo.

>```c++
> typedef struct ggml_backend_buffer * ggml_backend_buffer_t;
>```

### ggml_gallocr

È un allocatore di memoria specifico per i grafi computazionali (`ggml_cgraph`). Gestisce l'allocazione dei buffer necessari per l'esecuzione delle operazioni definite nel grafo, ottimizzando l'uso della memoria ed evitando riallocazioni.

>```c++
> typedef struct ggml_gallocr * ggml_gallocr_t;
>```

### ggml_backend_sched

È lo scheduler che consente l'utilizzo congiunto di più dispositivi backend. Gestisce l'allocazione dei buffer di calcolo, l'assegnazione e la copia dei tensori tra i backend. Può distribuire i calcoli su hardware diversi (ad esempio, GPU e CPU) quando si gestiscono modelli di grandi dimensioni o più GPU. Lo scheduler può anche assegnare automaticamente le operazioni non supportate dalla GPU alla CPU, garantendo un utilizzo ottimale delle risorse e la compatibilità.

>```c++
> typedef struct ggml_backend_sched * ggml_backend_sched_t;
>```

## Indice

1. [`Hello`](./hello/): Un'applicazione di esempio che dimostra l'utilizzo base di ggml nel calcolo matriciale.

