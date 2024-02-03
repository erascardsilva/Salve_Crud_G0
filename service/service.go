package service

import (
	"bufio"
	"fmt"
	"os"
	variav "salve/var"
)

func Escolha(op int) {
	switch op {
	case 1:
		Ler()
	case 2:

	default:

	}

}

func Ler() {
	fmt.Println("==================================")
	fmt.Println("Digite aqui o texto :")
	le := bufio.NewReader(os.Stdin)
	variav.Text, _ = le.ReadString('\n')
}
