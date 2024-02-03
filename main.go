package main

import (
	"Salve_Crud_GO/dados"
	"Salve_Crud_GO/menu"
	"Salve_Crud_GO/service"
	variav "Salve_Crud_GO/var"
)

func main() {
	dados.Connect()
	menu.Menu()
	service.Escolha(variav.Opt)

}
