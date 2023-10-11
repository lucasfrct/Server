package gitcommand

import (
	"fmt"
	"strings"
	"time"

	"github.com/lucasfrct/servertools/pkg/modules/command"
)

func GitAdd(filePath string) (gitString string) {
	return fmt.Sprintf(`git add %s`, filePath)
}

func GitAddAll() (gitString string) {
	return `git add .`
}

func GitCommit() (commited string) {
	return fmt.Sprintf(`git commit -m 'auto-commit-%v'`, time.Now().Format(time.RFC3339Nano))
}

func GitPush() (commited string) {
	return `git pull; git push`
}

func GitSync() (commited string) {
	return `git pull --all; git pull --tags; git pull -p; git fetch --prune; git fetch upstream --prune; git clean -f -x -d -n`
}

func GitProcessToCommit(modifiedArquives []string) (commited string) {
	for i := range modifiedArquives {
		commited += GitAdd(modifiedArquives[i]) + "; "
	}
	commited = fmt.Sprintf(`git fetch --all; %s; %s; %s; git reset --hard; %s`, commited, GitCommit(), GitPush(), GitSync())
	return
}

func GitListFiles() string {
	return `git status -s`
}

func GitListFilesModified(cmd string) []string {

	archives := strings.Split(cmd, "\n")
	modifiedArquives := []string{}

	for i := range archives {

		var trimedArchive string = strings.TrimSpace(archives[i])
		if len(trimedArchive) <= 1 {
			continue
		}

		var index string = strings.ToLower(strings.TrimSpace(trimedArchive[:2]))
		if !strings.Contains(index, "mm") && !strings.Contains(index, "??") && !strings.Contains(index, "m") && !strings.Contains(index, "a") {
			continue
		}

		modifiedArquives = append(modifiedArquives, strings.TrimSpace(trimedArchive[2:]))
	}

	return modifiedArquives
}

func GitPull(pathSource string) (cmdStr string) {
	cmd := "git pull"
	cmdStr, err := command.Exec(pathSource, cmd)

	if err != nil {
		panic(err)
	}
	return
}
