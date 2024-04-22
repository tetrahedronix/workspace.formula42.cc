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
    La funzione dofile esegue i chunk presenti nel file specificato come
    argomento. Se questa funzione è chiamata senza argomento allora viene
    esaminato lo standard input (stdin). Vengono restituiti i valori prodotti
    dal chunk contenuto nel file altrimenti il codice di errore.

    Eseguire il seguente codice dalla riga di comando

    lua dofile.lua

    oppure in modalità interattiva. Il programma hello.lua aspetta l'inserimento
    dell'input poi stampa la frase "Hello World!" seguita dal testo inserito
    dall'utente.
]]
dofile("hello.lua")