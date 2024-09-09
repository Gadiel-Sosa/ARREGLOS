package main

import (
	"fmt"
)

type Ventas struct {
	ventas        [][]int
	departamentos []string
}

func NewVentas() *Ventas {
	return &Ventas{
		ventas: [][]int{
			{5000, 4000, 6000, 7000, 5500, 8000, 6500, 7200, 6800, 5900, 7300, 7700}, // Ropa
			{3000, 3500, 2500, 4000, 3200, 3700, 3400, 3900, 3600, 3300, 3100, 3200}, // Deportes
			{7000, 7500, 8000, 8500, 7800, 8200, 7900, 8600, 8300, 8000, 7800, 8100}, // Juguetes
		},
		departamentos: []string{"Ropa", "Deportes", "Juguetes"},
	}
}

func (v *Ventas) insertarVentas(departamento string, mes int, monto int) {
	if !v.departamentoValido(departamento) {
		fmt.Println("Departamento no válido.")
		return
	}
	if mes < 0 || mes > 11 {
		fmt.Println("Mes no válido. Debe estar entre 0 y 11.")
		return
	}

	fmt.Println("Estado antes de insertar ventas:")
	v.mostrarVentas()

	indice := v.indiceDepartamento(departamento)
	v.ventas[indice][mes] = monto
	fmt.Printf("Ventas de %s en mes %d actualizadas a %d.\n", departamento, mes+1, monto)

	fmt.Println("Estado después de insertar ventas:")
	v.mostrarVentas()
}

func (v *Ventas) buscarVentas(departamento string, mes int) int {
	if !v.departamentoValido(departamento) {
		fmt.Println("Departamento no válido.")
		return -1
	}
	if mes < 0 || mes > 11 {
		fmt.Println("Mes no válido. Debe estar entre 0 y 11.")
		return -1
	}

	indice := v.indiceDepartamento(departamento)
	return v.ventas[indice][mes]
}

func (v *Ventas) eliminarVentas(departamento string, mes int) {
	if !v.departamentoValido(departamento) {
		fmt.Println("Departamento no válido.")
		return
	}
	if mes < 0 || mes > 11 {
		fmt.Println("Mes no válido. Debe estar entre 0 y 11.")
		return
	}

	fmt.Println("Estado antes de eliminar ventas:")
	v.mostrarVentas()

	indice := v.indiceDepartamento(departamento)
	v.ventas[indice][mes] = 0
	fmt.Printf("Ventas de %s en mes %d eliminadas.\n", departamento, mes+1)

	fmt.Println("Estado después de eliminar ventas:")
	v.mostrarVentas()
}

func (v *Ventas) mostrarVentas() {
	for i, departamento := range v.departamentos {
		fmt.Printf("Ventas en %s:\n", departamento)
		for mes := 0; mes < 12; mes++ {
			fmt.Printf("Mes %d: %d\n", mes+1, v.ventas[i][mes])
		}
		fmt.Println() // Línea en blanco para separar departamentos
	}
}

func (v *Ventas) buscarVentaEspecifica(monto int) [][2]interface{} {
	var resultados [][2]interface{}
	for i, departamento := range v.departamentos {
		for mes := 0; mes < 12; mes++ {
			if v.ventas[i][mes] == monto {
				resultados = append(resultados, [2]interface{}{departamento, mes})
			}
		}
	}
	return resultados
}

func (v *Ventas) departamentoValido(departamento string) bool {
	for _, dep := range v.departamentos {
		if dep == departamento {
			return true
		}
	}
	return false
}

func (v *Ventas) indiceDepartamento(departamento string) int {
	for i, dep := range v.departamentos {
		if dep == departamento {
			return i
		}
	}
	return -1
}

func main() {
	ventas := NewVentas()
	ventas.insertarVentas("Ropa", 0, 5500)
	ventas.insertarVentas("Deportes", 1, 3200)
	ventas.insertarVentas("Juguetes", 2, 7500)

	montoBuscado := 5000
	resultados := ventas.buscarVentaEspecifica(montoBuscado)
	if len(resultados) > 0 {
		fmt.Printf("Ventas de %d encontradas en:\n", montoBuscado)
		for _, resultado := range resultados {
			departamento := resultado[0].(string)
			mes := resultado[1].(int)
			fmt.Printf("Departamento: %s, Mes: %d\n", departamento, mes+1)
		}
	} else {
		fmt.Printf("No se encontraron ventas con el monto %d.\n", montoBuscado)
	}

	ventas.eliminarVentas("Deportes", 1)
}
