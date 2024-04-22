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

-- definisce la funzione fact che calcola ricorsivamente il fattoriale di
-- un numero.
function fact(numero)
    if numero == 0 then
        return 1
    else
        return numero * fact(numero - 1)
    end
end

print("enter a number: ")
a = io.read("*n")   -- legge un numero dallo standard input
print(fact(a))
--[[
    *Chunks*
    Ogni pezzo di codice che l'interprete Lua esegue, per esempio un file 
    oppure una singola linea di codice in modalità interativa, è chiamato 
    chunk.

    Un chunk è una istruzione oppure una combinazione di istruzioni e funzioni
    e non c'è limite alla sua dimensione.

    Caricare il file sorgente fattoriale.lua con il comando

    lua -i fattoriale.lua

    L'interprete esegue il chunk contenuto nel file e poi entra in modalità 
    interattiva. Questo comportamento dell'interprete risulta utile quando
    si intende effettuare un debugging.
]]
