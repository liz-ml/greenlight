package validator

//necessário pois não basta receber e decodificar o json

import (
	"regexp" //suporta para trabalhar com Expressões regulares (validar padrões de texto, formato de e-mail)
)

	//EmailRX começa com letra maiúscula para ser exportada
var (
	EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9")
)

type Validator struct {
	Errors map[string]string
}

func New() *Validator { //função construtora que retorna um ponteiro
	//cria uma instância de validator (nil) e retorna seu endereço de memória
	//tentar inserir dados em um mapa nil causa panic, por isso o contrutor New()
	return &Validator{Errors: make(map[string]string)}
}

func (v *Validator) Valid() bool {
	return len(v.Errors) == 0 // se o tamanho for 0 retorna true
}

func (v *Validator) AddError(key, message string) {
	// consultar um mapa retorna primeiro o valor daquela chave e depois um boolean
	if _, exists := v.Errors[key]; !exists { // o _ descarta o texto do valor
		v.Errors[key] = message
	}
}

func (v *Validator) Check(ok bool, key, message string) {
	if !ok {
		v.AddError(key, message)
	}
}


func In(value string, list ...string) bool {
	for i := range list {
		if value == list[i] {
			return true
		}
	}

	return false
}

func Matches(value string, rx *regexp.Regexp) bool {
	return rx.MatchString(value)
}

func Unique(values []string) bool {
	uniqueValues := make(map[string]bool) //go não permite chaves duplciadas
	for _, value := range values {
		uniqueValues[value] = true
	}
	return len(values) == len(uniqueValues)
}