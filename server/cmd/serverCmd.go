// Server related command.
package cmd

import (
	"context"
	"fmt"
	"main/db"
	"main/handler"
	"net/http"
	"os"
)

func initWeb(repoPath string) error {
	if repoPath == "" {
		fmt.Println("No repoPath specified. Used current directory as repo.")
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error when get pwd: ", err)
		}
		repoPath = pwd
	}
	_, err := db.InitRepo(repoPath)
	return err
}

func startWeb(repoPath string) error {
	if repoPath == "" {
		fmt.Println("No repoPath specified. Used current directory as repo.")
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error when get pwd: ", err)
		}
		repoPath = pwd
	}
	repo, err := db.OpenRepo(repoPath)
	if err != nil {
		fmt.Println("Error when open repo: ", err)
		return err
	}
	http.Handle("/", handler.NewHttpHandler(repo))
	fmt.Println("Start server and listen port 8080")
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			fmt.Printf("Failed to start server: %v\n", err)
		}
	}()
	<-context.Background().Done()
	return nil
}
