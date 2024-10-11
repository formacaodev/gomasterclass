package funcoes

import (
	"fmt"
	"strings"
)

func testeEmail() {
	email := "leonardo@formacao.dev"
	usuario, dominio := partesDoEmail(email)

	fmt.Println("Usuário:", usuario)
	fmt.Println("Domínio:", dominio)
}

func partesDoEmail(email string) (usuario, dominio string) {
	partes := strings.Split(email, "@")
	usuario = partes[0]
	dominio = partes[1]
	return usuario, dominio
}
