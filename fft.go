package main

import (
	"fmt"
	"log"
	"math"
	"math/cmplx"
	"net/http"

	"github.com/gorilla/mux"
)

var mag float64
var convertir_a_cadena string
var respuesta string

// Defino los tipos de Datos de mi aplicación.

func homeLink(w http.ResponseWriter, r *http.Request) {
	x := []float64{0, 1, 2, 3, 4, 5, 6, 7}
	y := make([]complex128, len(x))

	ditfft2(x, y, len(x), 1)

	respuesta = "["
	for _, c := range y {
		if len(respuesta) > 1 {
			respuesta = respuesta + ","
		}
		fmt.Printf("%2.4f\n", c)
		mag := math.Sqrt(real(c)*real(c) + imag(c)*imag(c))
		mag_s := fmt.Sprintf("%v", mag)
		respuesta = respuesta + mag_s
		fmt.Println("Magnitud", mag)
		fase := math.Atan(real(c)/imag(c)) * 180 / math.Pi
		fmt.Println("Fase", fase)
	}
	respuesta = respuesta + "]"

	fmt.Println(respuesta)

	code := 200 // preparando cabeceras
	fmt.Println(code)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	fmt.Fprintf(w, "%s", respuesta)

}

func ditfft2(x []float64, y []complex128, n, s int) {
	if n == 1 {
		y[0] = complex(x[0], 0)
		return
	}

	ditfft2(x, y, n/2, 2*s)
	ditfft2(x[s:], y[n/2:], n/2, 2*s)

	for k := 0; k < n/2; k++ {
		//Rect devuelve el número complejo x con coordenadas polares r, θ calcula la magnitud y el angulo
		tf := cmplx.Rect(1, -2*math.Pi*float64(k)/float64(n)) * y[k+n/2]
		y[k], y[k+n/2] = y[k]+tf, y[k]-tf
	}

}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))

}
