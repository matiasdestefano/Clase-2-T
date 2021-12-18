package main

import "fmt"

//Una empresa de inteligencia artificial necesita tener una funcionalidad para crear una
//estructura que represente una matriz de datos.
//Para ello requieren una estructura Matrix que tenga los métodos:

//Set:  Recibe una serie de valores de punto flotante e inicializa los valores en la estructura Matrix
//Print: Imprime por pantalla la matriz de una formas más visible (Con los saltos de línea entre filas)
//La estructura Matrix debe contener los valores de la matriz, la dimensión del alto, la dimensión del ancho, 
//si es cuadrática y cuál es el valor máximo.

type Matriz struct {
	Valores 	[3][3]int
	Filas 		int
	Columnas 	int
	EsCuadratica bool
	ValorMaximo int
}

func main() {
	matriz := Matriz{}
	matriz.Set(1, 2, 3, 4, 5, 6)
	matriz.Print()
}

func (m *Matriz) Set(valores ...int) {
	m.Filas = 3
	m.Columnas = 3

	if(m.Filas == m.Columnas) {
		m.EsCuadratica = true
	} else {
		m.EsCuadratica = false
	}

	m.llenarValores(valores)
	m.obtenerMaximo()
}

func (m Matriz) Print() {
	for i := 0; i < m.Filas; i++ {
		for j := 0; j < m.Columnas; j++ {
			fmt.Printf("%d\t", m.Valores[i][j])
		}
		fmt.Printf("\n")
	}
}

func (m *Matriz) obtenerMaximo() {

	m.ValorMaximo = m.Valores[0][0]

	for i := 0; i < m.Filas; i++ {
		for j := 0; j < m.Columnas; j++ {
			if m.Valores[i][j] > m.ValorMaximo {
				m.ValorMaximo = m.Valores[i][j]
			}
		}
	}
}

func (m *Matriz) llenarValores(valores []int) {
	k := 0
	for i := 0; i < m.Filas; i++ {
		for j := 0; j < m.Columnas; j++ {
			m.Valores[i][j] = valores[k]
			k++

			if k >= len(valores) {
				return
			}
		}
	}
}