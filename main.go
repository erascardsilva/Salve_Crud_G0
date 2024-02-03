package main

import (
	"salve/dados"
	"salve/menu"
	"salve/service"
	variav "salve/var"
)

func main() {
	dados.Connect()
	menu.Menu()
	service.Escolha(variav.Opt)

}
