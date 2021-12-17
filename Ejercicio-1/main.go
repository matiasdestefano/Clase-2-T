package main

import "fmt"
//Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

//Nombre: [Nombre del alumno]
//Apellido: [Apellido del alumno]
//DNI: [DNI del alumno]
//Fecha: [Fecha ingreso alumno]

//Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
//Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle


type Alumno struct {
	Nombre string 
	Apellido string 
	DNI int 
	FechaIngreso Fecha
}

type Fecha struct {
	Dia int
	Mes int
	Año int
}

func main() {
	alumno := Alumno{Nombre: "Matias", Apellido: "De Stefano", DNI: 40130248, FechaIngreso: Fecha{Dia: 31, Mes: 1, Año: 1997}}
	
	fmt.Printf(alumno.Detalle())
}

func (alumno Alumno) Detalle() string {
	return fmt.Sprintf("Nombre: %s\nApellido: %s\nDNI: %d\nFecha de Ingreso: %d/%d/%d", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.FechaIngreso.Dia, alumno.FechaIngreso.Mes, alumno.FechaIngreso.Año)
}