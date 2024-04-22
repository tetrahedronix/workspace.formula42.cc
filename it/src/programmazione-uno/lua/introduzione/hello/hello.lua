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
    È possibile eseguire i programmi con l'interprete lua in modalità
    interattiva: omettendo gli argomenti alla riga dei comandi si attiva
    l'interprete interattivo. In questa modalità ogni istruzione è eseguita
    immediatamente dopo il carattere di fine linea. Per terminare l'interazione
    usare i caratteri di controllo di fine file, CTROL+D oppure usare la 
    funzione lua os.exit(). È anche possibile usare l'interprete come una
    calcolatrice in linea specificando le espressioni matematiche.
]]
print("Hello World")

--[[
    Per risolvere espressioni matematiche nel chunk di un file sorgente occorre
    trattarle come argomento della funzione print
]]
print(math.pi/4)
--[[
    Ogni linea inserita nella modalità interattiva è trattata come un chunk 
    completo oppure una espressione. Comunque se l'interprete riconosce che
    la linea inserita non è sintatticamente completa allora aspetterà che venga
    specificato il resto del codice nella linea successiva
]]
print((math.pi / 4 +
20 - 4) * 3)