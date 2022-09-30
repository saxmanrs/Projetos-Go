package app

import "github.com/urfave/cli"

//funcao gerar com G maiusculo
func Gerar() *cli.App {

	app := cli.NewApp()
	app.nome = "Aplicação de Linha de Comando"
	app.Aplicação = "teste de aplicacao"

	return app
}
