package handlers

import (
	"adn/db"
	"adn/models"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

var matrix [6][6]string
var contMutan int

type Request struct {
	Adn []string `json:"adn"`
}

type Response struct {
	Response string `json:"message"`
	Status   int    `json:"status"`
}

func GetAdn(rw http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(rw, "Listado de adn")

	var isMuttan int

	var datos []models.Adn = db.GetAdn()

	fmt.Println(datos[1].Mutante)

	for i := 0; i < len(datos); i++ {
		if datos[i].Mutante {
			isMuttan++
		}
	}

	message := "Cant Muttan " + strconv.Itoa(isMuttan)

	fmt.Println(message)
}

func CreateAdn(rw http.ResponseWriter, r *http.Request) {

	//Obtener Registro
	decoder := json.NewDecoder(r.Body)

	adnr := &Request{}

	if err := decoder.Decode(&adnr); err != nil {
		panic(err)
	} else {
		fillMatriz(adnr)
		validateMutan()
	}

	if contMutan > 1 {
		adn := models.Adn{}
		adn.Secuencia = adnr.Adn
		adn.Mutante = true

		data := Response{
			"Is mutan",
			http.StatusOK,
		}

		db.SaveAdn(&adn)
		sendData(rw, data, http.StatusOK)
	} else {
		sendError(rw, http.StatusForbidden)
	}

}

func fillMatriz(adn *Request) {

	tamaño := len(adn.Adn)

	for k := 0; k < tamaño; k++ {
		cadena := strings.Split(adn.Adn[k], "")
		for i := k; i < k+1; i++ {
			for j := 0; j < len(matrix); j++ {
				matrix[i][j] = cadena[j]
			}
		}

	}

}

func validateMutan() {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix); j++ {

			leA := matrix[i][j]

			// Vertical
			if i <= len(matrix)-3 && leA == matrix[i+1][j] && leA == matrix[i+2][j] &&
				leA == matrix[i+3][j] {
				contMutan++
			}

			// Horisontal
			if j <= len(matrix)-3 && leA == matrix[i][j+1] && leA == matrix[i][j+2] &&
				leA == matrix[i][j+3] {
				contMutan++
			}

			//Obliuca
			if j <= len(matrix)-3 && i <= len(matrix)-3 && leA == matrix[i+1][j+1] &&
				leA == matrix[i+3][j+2] && leA == matrix[i+3][j+3] {
				contMutan++
			}

		}
	}

}
