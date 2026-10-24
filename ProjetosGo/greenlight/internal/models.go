package internal

import (
	"log"
)


type Config struct {
	//configuração que o programa aceita
	Port int
	Env string
}

type Application struct { //agrupa dependências globais
	//padrao em aplicações Go (db pool), é passada para os handlers como "receiver" (app *application) para que eles tenham acesso a tudo que precisam sem usar variáveis globais
	
	Config Config
	Logger *log.Logger
}
