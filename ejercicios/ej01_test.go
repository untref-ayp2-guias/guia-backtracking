package ejercicios

import (
	"fmt"
	"testing"
)

func TestEjercicio1(t *testing.T) {
	n := 5
	k := 3

	solucion := Ejercicio1(n, k)
	fmt.Println(solucion)

	if len(solucion) != 10 {
		t.Error("La solucion no es correcta")
	}
}
