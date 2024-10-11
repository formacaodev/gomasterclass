package utils

import (
	"fmt"
	"strings"
)

func Executar(secao string, funcoes ...func()) {
	for i, funcao := range funcoes {
		fmt.Printf("\n\n>>> %s - Exercício %d <<<\n", strings.ToUpper(secao), i+1)
		funcao()
	}
}
