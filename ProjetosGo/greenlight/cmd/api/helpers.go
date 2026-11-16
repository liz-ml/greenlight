package main

import (
	"errors"
	"net/http"
	"strconv"
	"github.com/julienschmidt/httprouter"
	"encoding/json"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	//retorna tudo em string, por isso é necessário converter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}

		return id, nil
	
}



//helper para ajusar a mandar respostas. Pega o destino w, o status HTTP para enviar, os dados para o Json e um header map contendo cabeçalhos HTTP adicionais que queremos incluir na resposta
//inteface aceita diferentes tipos de dados
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//codifica os dados para Json, retornando erro caso falhe. Retorna dois valores: []byte, error
	js, err := json.Marshal(data)

	if err != nil {
		return err
	}
// nova linha para visualizar melhor
	js = append(js, '\n')


//http.Header é um tipo baseado em map[string][]string, assim cada cabeçalho pode ter uma lista de valores. esse bloco percorre essa lista, em que key é o nome do cabeçalho, value o slice de valores e no final esses valores sao colocados na resposta, 
// substituindo valores antigos para aquela chave
	for key, value := range headers {
		w.Header()[key] = value
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst interface{}) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		var syntaxError *json.SyntaxError
		var unmarshalTypeError *json.UnmarshalTypeError
		var invalidUnmarshalError *json.InvalidUnmarshalError
		switch {

		case errors.As(err, &syntaxError):
			return fmt.Errorf("body contains badly-formed JSON (at character %d)", syntaxError.Offset)

		case errors.Is(err, io.ErrUnexpectedEOF):
			return errors.New("body contains badly-formed JSON")

		case errors.As(err, &unmarshalTypeError):
			if unmarshalTypeError.Field != "" {
				return fmt.Errorf("body contains incorrect JSON type for field %q", unmarshalTypeError.Field)
			}
			return fmt.Errorf("body contains incorrect JSON type (at character %d)", unmarshalTypeError.Offset)

		case errors.Is(err, io.EOF):
			return errors.New("body must not be empty")

		case errors.As(err, &invalidUnmarshalError):
			panic(err)
		default:
			return err
		}
	}	

	return nil
}

