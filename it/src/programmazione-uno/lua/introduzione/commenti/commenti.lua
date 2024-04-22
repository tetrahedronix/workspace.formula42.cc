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
    In Lua i blocchi di commenti si costruiscono con la combinazione di simboli 
    --[[ per l'apertura e il doppio simbolo ] per la chiusura. I commenti
    non sono nidificabili nel blocco.
]]

-- Un commento in linea inizia con la combinazione di simboli --
-- Tutto ciò che viene scritto dopo questi simboli viene ignorato
-- dall'interprete.

--[[
    I commenti sono utilissimi per poter descrivere il comportamento di un 
    programma oppure per disabilitare linee di codice, aggiungere note per
    altri sviluppatori. Osservando questi principi, i commenti possono essere
    utilizzati per migliorare la leggibilità del codice, rendendolo più facile
    da comprendere e mantenere. Tuttavia è importante utilizzare i commenti
    in modo appropriato: un uso eccessivo di commenti può rendere il codice
    più difficile da leggere e da mantenere.
]]
function Hello()
    print("Hello World")
    print("Hello Jack!")
end

--[[ 
    Rimuovendo il primo trattino - della linea di codice seguente, si disabilita
    la chiamata alla funzione Hello e il programma non stamperà le frasi di
    benvenuto. Infatti, l'interprete considera ---[ come un commento
    in linea.
]]
---[[
    Hello()
--]]