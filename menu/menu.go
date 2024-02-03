package menu

import (
	"Salve_Crud_GO/service"
	variav "Salve_Crud_GO/var"
	"fmt"
)

func Menu() {
	for {
		Estru()
		fmt.Println("\nMenu \n Opções : \n",
			"1) Gravar Dados \n 2) Ler Dados\n 3) Fim")
		Estru()
		fmt.Println("\nESCOLHA  : ")

		fmt.Scanln(&variav.Men)
		switch variav.Men {
		case 1:
			service.Escolha(1)
		case 2:

		case 3:
			return
		default:
			continue
		}
	}
}

func Estru() {
	for i := 0; i < 20; i++ {
		fmt.Print("=")
	}
}
