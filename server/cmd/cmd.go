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
	startNoCeph := startCmd.Arg("noCeph", "Whether to run server without ceph. " +
		"It would be false by default.").Bool()
	cmds[startCmd.FullCommand()] = func() error {
		return startWeb(*startRepoPath, *startNoCeph)
	}

	addVideoCmd := appCmd.Command("addVideo", "Add a video to server.")
	cmds[addVideoCmd.FullCommand()] = func() error {
		//return sendRequest()
		return nil
	}

	cephCmd := appCmd.Command("ceph", "ceph related commands. " +
		"Note that these commands are only used for test.")
	cephTestCmd := cephCmd.Command("test", "Test whether ceph works or not.")
	cmds[cephTestCmd.FullCommand()] = func() error {
		return touchCeph()
	}

	videoCmd := appCmd.Command("video", "Video related commands.")
	videoAddCmd := videoCmd.Command("add", "Add a new video.")
	videoAddPath := videoAddCmd.Arg("path", "Path of the video.").Required().String()
	videoAddPoster := videoAddCmd.Arg("poster", "Poster of the video.").Required().String()
	videoAddName := videoAddCmd.Arg("name", "Name of the video.").Required().String()
	videoAddLength := videoAddCmd.Arg("length", "Length of the video in seconds.").Required().Int()
	cmds[videoAddCmd.FullCommand()] = func() error {
		return addVideo(*videoAddPath, *videoAddPoster, *videoAddName, *videoAddLength)
	}

	cmd := kingpin.MustParse(appCmd.Parse(os.Args[1:]))
	for key, value := range cmds {
		if key == cmd {
			return value()
		}
	}

	return nil
}