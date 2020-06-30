// Server related command.
package cmd

import (
	"context"
	"fmt"
	"github.com/ceph/go-ceph/rados"
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

func startWeb(repoPath string, startHost string, noCeph bool) error {
	if repoPath == "" {
		fmt.Println("No repoPath specified. Used current directory as repo.")
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error when get pwd: ", err)
		}
		repoPath = pwd
	}

	// open repo and database
	repo, err := db.OpenRepo(repoPath)
	if err != nil {
		fmt.Println("Error when open repo: ", err)
		return err
	}
	var conn *rados.Conn
	conn = nil
	if !noCeph {
		fmt.Println("Connecting ceph.")
		// open connection to ceph
		conn, err = rados.NewConn()
		if err != nil {
			fmt.Println("Error when create ceph connection: ", err)
			return err
		}
		err = conn.ReadDefaultConfigFile()
		if err != nil {
			fmt.Println("Error when read ceph config: ", err)
			return err
		}
		err = conn.Connect()
		if err != nil {
			fmt.Println("Error when connect to ceph: ", err)
			return err
		}
		defer conn.Shutdown()
	} else {
		fmt.Println("Running without ceph.")
	}
	http.Handle("/", handler.NewHttpHandler(repo, conn, startHost))
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
