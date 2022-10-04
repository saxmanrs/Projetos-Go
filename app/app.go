package app

import "github.com/urfave/cli"

//funcao gerar com G maiusculo
func Gerar() *cli.App {

	app := cli.NewApp()
	app.Name = "Primeira Aplicação"
	app.Usage = "Busca de Ips e nome de servidores"

	return app

	app.Commands = []cli.Command{
		{
			Name: "IP",
			Usage: "Busca de IPs",
			Flags: []cli.Flag{



				
			},

		}
		
	}

}
