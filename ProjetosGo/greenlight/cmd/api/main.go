package main

import ( "flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

//as structs passadas como receiver), devem estar no mesmo local dos métodos
type config struct {
	//configuração que o programa aceita
	port int
	env string
}

type application struct { //agrupa dependências globais
	//padrao em aplicações Go (db pool), é passada para os handlers como "receiver" (app *application) para que eles tenham acesso a tudo que precisam sem usar variáveis globais
	
	config config
	logger *log.Logger
}
func main () {
	var cfg config
	// liga a flag -port a variavel cfg.port
	// usar flag quando queremos que seja possível passar parâ
	//metros na CLI sem alterar o código fonte
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env,"env", "development", "Environment (development|staging|production)")

	flag.Parse() //lê os argumentos passados na CLI
	// go run main.go -h ou --help exibe ajuda

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	app := &application{ //passar um ponteiro permite que todos os handlers recebam *application como receiver, compartilhem o mesmo estado, sem copiar a struct inteira.
//usar ponteiro quando: structs grandes, com slices, maps, necessidade de compartilhar/modificar o mesmo objeto em vários lugares
//não usar quando a struct deve ser tratada como snapshot (cópia independente)
		config: cfg,
		logger: logger,
	}

	//mux := http.NewServeMux() //router padrão que mapeia paths para handlers (função que recebe http.ResponseWriter e *http.Request; permite criar múltilplos routers independentes (versionamento de API); + fácil de testar



	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		Handler: app.routes(),
		IdleTimeout: time.Minute,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", app.config.env, srv.Addr)

	err := srv.ListenAndServe()
	logger.Fatal(err)

}



