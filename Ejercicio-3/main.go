package main

type Producto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p producto)
}

type producto struct {
	TipoDeProducto string
	Nombre         string
	Precio         float64
}

const (
	Pequeño = "Pequeño"
	Mediano = "Mediano"
	Grande  = "Grande"
)

type tienda struct {
	listaDeProductos []producto
}

func main() {

}

func nuevoProducto(tipoDeProducto string, nombre string, precio float64) producto {
	var productoAgregado producto
	productoAgregado.TipoDeProducto = tipoDeProducto
	productoAgregado.Nombre = nombre
	productoAgregado.Precio = precio
	return productoAgregado
}

func nuevaTienda() Ecommerce {
	var t tienda
	return &t
}

func (t *tienda) Total() float64 {
	var total float64
	for _, v := range t.listaDeProductos {
		total += v.Precio
	}

	return total
}

func (p producto) CalcularCosto() float64 {
	if p.TipoDeProducto == Pequeño {
		return p.Precio
	} else if p.TipoDeProducto == Mediano {
		return p.Precio + p.Precio*0.03
	} else if p.TipoDeProducto == Grande {
		return p.Precio + p.Precio*0.06 + 2500
	} else {
		return 0
	}
}

func (t *tienda) Agregar(p producto) {
	t.listaDeProductos = append(t.listaDeProductos, p)
}
