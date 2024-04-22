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
    All'interno dei programmi le variabili globali sono inizializzate a nil ma
    non c'è bisogno di dichiararle, semplicemente basta utilizzare 
    l'identificatore
]]

--[[
    La variabile a è globale e non è stata dichiarata, il suo valore è nil.
]]    
print(a)    -- Stampa nil
a = 10
print(a)    -- Stampa 10

-- L'assegnamento di nil reinizializza la variabile
a = nil     -- A questo punto la locuzione di memoria è di nuovo disponibile.

-- Crea una nuova variabile globale a
print(a)