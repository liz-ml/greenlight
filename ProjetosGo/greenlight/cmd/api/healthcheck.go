package main

import (
	"fmt"
	"net/http"
)
//go não é OO, não tem classes, por isso vincular uma função a um struct é a única forma de aplicar os conceito de OO, como encapsulatemtno, polimorfismo, etc
//receiver é um parâmetro especial colocado antes do nome de uma função que a associa a um tipo específico, geralmente struct
func (app *application) HealthcheckHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "status: available\n")
	fmt.Fprintf(w, "enviroment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

