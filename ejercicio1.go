package main

import "fmt"

func main() {
	var notas [6][4]float64
	var promedios []float64
	var notasAltas []float64
	var notasBajas []float64
	var sumaGeneral float64

	for i := 0; i < 6; i++ {
		fmt.Printf("\nEstudiante %d\n", i+1)

		for j := 0; j < 4; j++ {
			fmt.Printf("Ingrese la nota de la materia %d: ", j+1)
			fmt.Scan(&notas[i][j])

			for notas[i][j] < 0 || notas[i][j] > 100 {
				fmt.Println("Nota invalida. Ingrese una nota entre 0 y 100:")
				fmt.Scan(&notas[i][j])
			}
		}
	}

	for i := 0; i < 6; i++ {
		var suma float64
		mayor := notas[i][0]
		menor := notas[i][0]

		for j := 0; j < 4; j++ {
			suma += notas[i][j]

			if notas[i][j] > mayor {
				mayor = notas[i][j]
			}

			if notas[i][j] < menor {
				menor = notas[i][j]
			}
		}

		promedio := suma / 4

		promedios = append(promedios, promedio)
		notasAltas = append(notasAltas, mayor)
		notasBajas = append(notasBajas, menor)

		sumaGeneral += promedio
	}

	fmt.Println("\n------ RESULTADOS ------")

	for i := 0; i < len(promedios); i++ {
		fmt.Printf("\nEstudiante %d\n", i+1)
		fmt.Printf("Promedio: %.2f\n", promedios[i])
		fmt.Printf("Nota mas alta: %.2f\n", notasAltas[i])
		fmt.Printf("Nota mas baja: %.2f\n", notasBajas[i])
	}

	promedioGeneral := sumaGeneral / float64(len(promedios))

	fmt.Printf("\nPromedio general de la clase: %.2f\n", promedioGeneral)
}
