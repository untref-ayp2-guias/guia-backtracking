# Guía 10 - Backtracking
## Implementaciones de las diapositivas

En las siguientes carpetas se encuentran las implementaciones de algunos problemas utilizando backtracking:

- `/ratmaze`, problema del laberinto, resuelto con backtracking
- `/cambio`, problema del cambio de moneda, resuelto con backtracking
- `/sudoku`, resolución de un sudoku clásico, utilizando backtracking
- `/reinas`, resolución del problema de las n reinas, utilizando backtracking

- `/main`, ejecuciones adicionales con fines interactivos

## Ejercicios

En la carpeta `/ejercicios` encontrarás los esqueletos de la implementación para las siguientes consignas.

### Backtracking

Usando los conceptos de backtracking resuelva los siguientes ejercicios. 

0. En los ejercicios del laberinto y sudoku investigar la estructura de datos que soporta la solución, e identificar las operaciones Factible, esSolución, Extender, Registrar y Borrar. Tomar como ejemplo el código de cambio de moneda y N-reinas que está comentado.
   
1. Dado un valor y 3 dados, queremos encontrar todas las combinaciones que superen o igualen a ese valor dado.
   
    > Ejemplo: Si el valor es 15, las combinaciones pueden ser: (6, 6, 6), (6, 5, 6), (5, 5, 5), (5, 4, 6)

2. Problema de la mochila (Knapsack). Este es uno de los 21 problemas de [NP completos](https://es.wikipedia.org/wiki/NP-completo) de Richard Karp, y ya [mencionado en 1896](https://doi.org/10.1112%2Fplms%2Fs1-28.1.486) por Mathews, G. B. Consiste en un problema de optimización combinatoria, donde se espera poder llenar una "mochila" con un peso limitado, por una cantidad de objetos, cada uno con un peso y valor específico, máximizando el valor total almacenado. Hay una simplificación del mismo, denominado problema de la mochila 0-1 donde un objeto puede estar o no dentro de la mochila, por completo, es decir no se puede fraccionar
