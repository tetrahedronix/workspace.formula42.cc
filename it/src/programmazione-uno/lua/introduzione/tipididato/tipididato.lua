--[[
 * Copyright (C) 2023 Giulio Carlo
 *
 * This file is part of programmazione-uno.
 *
 * programmazione-uno is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * programmazione-uno is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with programmazione-uno.  If not, see <http://www.gnu.org/licenses/>.
--]]

--[[
    Lua è un linguaggio dinamicamente tipizzato, significa che l'interprete
    assegna alle variabili o alle espressioni il tipo associato con i valori.
    Non c'è una istruzione per la definizione dei tipi da utilizzare con le
    variabili e le espressioni.
    Ci sono essenzialmente otto tipi di dato di base: nil, booleano, numero,
    stringa, userdata, funzione, thread e tabella.
    
    I valori in Lua sono detti di prima classe, cioè sono assegnabili alle
    variabili, passati come argomenti di funzione e restituiti come risultati.
]]
-- Assegnazione di un valore di tipo strina alla variabile a
a = "Hello World"
-- La funzione type restituisce la stringa del tipo associato alla variabile a
print(type(a))  -- Stampa string

-- Assegnazione di un valore di tipo numerico alla variabile a
a = 5
print(type(a))  -- Stampa number

--[[
    Il tipo nil ha un solo valore che è nil. L'assegnazione di un valore nullo
    nil a una variabile la reinizializza e la locuzione di memoria è resa di 
    nuovo disponibile. Il tipo nil rappresenta l'assenza di un valore utile.
]]
a = nil
print(type(a))  -- Stampa nil, il tipo della variabile a

--[[
    Il tipo booleano boolean ha due valori possibili: false per falso e true 
    per il valore logico vero. Si osserva che false e nil nelle espressioni 
    logiche hanno lo stesso significato, qualsiasi altro valore ha significato
    logico vero. In particolare la stringa vuota "" e l'intero 0 assumono il
    valore booleano vero nelle espressioni logiche delle strutture di controllo.
]]
if "Yes" then
    print("it's true")      -- Stampa la frase it's true
end

if nil then
    print("It's false")     -- Non viene eseguita questa istruzione
end

--[[
    Il tipo numerico number rappresenta sia numeri interi che reali (virgola
    mobile)), cioè i due sottotipi integer e float. Lo standard Lua specifica
    interi a 64 e float a 64 in doppia precisione.

    Si osserva che qualsiasi overflow sui numeri interi produce valori che si 
    "riavvolgono" secondo le regole dell'aritmetica in complemento a due.

    Lua ha regole esplicite per trattare ciascun sottotipo, ma è sempre 
    possibile procedere con le coversione automatiche fra tipi, per cui il 
    programmatore potrebbe ignorare, in prima istanza, la differenza fra i 
    sottotipi integer e float. 
]]
num = 1             -- Assegna alla variabile il valore intero 1 
print(type(num))    -- Stampa number
num = math.pi       -- Assegna alla variabile il valore float pi greco
print(type(num))    -- Stampa number

--[[
    Il tipo stringa string rappresenta sequenze di caratteri (byte) costanti.
    L'unica limitazione per una stringa è la sua lunghezza che non può eccedere
    64 bit.
]]
s = "Hello World"   -- Assegna alla variabile il valore stringa Hello World
print(type(s))      -- Stampa string

--[[
    il tipo useradata permette la memorizzazione di dati C in variabili Lua. 
    Un valore di questo tipo rappresenta un blocco di memoria grezza. 
    Ci sono due sottotipi di userdata: full useradata che è un oggetto con
    un blocco di memoria gestito da Lua e light userdata che è semplicemente
    un puntatore C al valore vero e proprio. Con questo tipo di dato sono
    poossibili le operazioni di assegnamento e testing.
]]
