package ejercicios

// Ej1. Encontrar todas las combinaciones posibles de un un conjunto de n
// números tomados de a k a la vez.

// Ej2. Encontrar todas las permutaciones posibles de un conjunto de n
// números.
func Ejercicio1(n, k int) [][]int {
	var solucion [][]int
	var combinacion []int
	backtracking1(n, k, combinacion, &solucion)
	return solucion
}

func backtracking1(n, k int, combinacion []int, solucion *[][]int) {
	if len(combinacion) == k { //Solucion
		*solucion = append(*solucion, combinacion)
		return
	}
	for i := 1; i <= n; i++ {
		combinacion = append(combinacion, i)           //Registrar
		backtracking1(n, k, combinacion, solucion)     //Extender
		combinacion = combinacion[:len(combinacion)-1] //Borrar
	}
}
