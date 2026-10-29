package main

import (
	"fmt"
	"net/http"
)
//go não é OO, não tem classes, por isso vincular uma função a um struct é a única forma de aplicar os conceito de OO, como encapsulatemtno, polimorfismo, etc
//receiver é um parâmetro especial colocado antes do nome de uma função que a associa a um tipo específico, geralmente struct
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	
	js := `{"status": "available", "environment": %q, "version" : %q}`
	//Sprintf cria e retorna uma string formatada na memória
	js = fmt.Sprintf(js, app.config.env, version)

	w.Write([]byte(js))

}
