package cmd

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"os"
)
const localhost = "http://localhost:8080"
type cmdsMap map[string]func() error
func Run() error {
	appCmd := kingpin.New("cephweb",
		"cephweb is a set of tools that help you to run and manage a video on demand web based on ceph storage")
	cmds := make(cmdsMap)
	initCmd := appCmd.Command("init", "Initialize the repository for web.")
	initRepoPath := initCmd.Arg("repo", "The path of repository.\n" +
		"cephweb would use current path as repository if not specified.\n" +
		"Note that there should be {resource, templates} directories within repository.").String()
	cmds[initCmd.FullCommand()] = func() error {
		return initWeb(*initRepoPath)
	}

	startCmd := appCmd.Command("start", "Start web server.")
	startRepoPath := startCmd.Arg("repo","The path of repository.\n" +
		"cephweb would use current path as repository if not specified.\n" +
		"Note that there should be {resource, templates} directories within repository.").String()
	cmds[startCmd.FullCommand()] = func() error {
		return startWeb(*startRepoPath)
	}

	addVideoCmd := appCmd.Command("addVideo", "Add a video to server.")
	cmds[addVideoCmd.FullCommand()] = func() error {
		return sendRequest()
	}

	cmd := kingpin.MustParse(appCmd.Parse(os.Args[1:]))
	for key, value := range cmds {
		if key == cmd {
			return value()
		}
	}

	return nil
}