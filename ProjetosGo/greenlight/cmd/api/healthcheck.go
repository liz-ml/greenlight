package main

import (
	"net/http"
)
//go não é OO, não tem classes, por isso vincular uma função a um struct é a única forma de aplicar os conceito de OO, como encapsulatemtno, polimorfismo, etc
//receiver é um parâmetro especial colocado antes do nome de uma função que a associa a um tipo específico, geralmente struct
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){
	data := map[string]string{
		"status": "available",
		"environment": app.config.env,
		"version": version,
	}
	
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process yout request", http.StatusInternalServerError)
	}
}
