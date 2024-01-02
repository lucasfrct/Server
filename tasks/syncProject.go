package tasks

import (
	"fmt"

	"github.com/lucasfrct/servertools/pkg/modules/gitcommand"
)

func SyncProject(source, dest string) {
	err := Copy(source, dest)
	if err != nil {
		panic("Erro ao tentar copiar arquivos")
	}

	fmt.Println(fmt.Sprintf("* Arquivos copiados: (%s) -> (%s)", source, dest))
	resp := CommitFilesModified(dest)
	fmt.Println(fmt.Sprintf("* Projeto Commitado (%s): %s", dest, resp))
	resp = gitcommand.GitPull(dest)
	fmt.Println(fmt.Sprintf("* Projeto sincronizado (%s): %s", dest, resp))

}
