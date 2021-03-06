// Manage repository of web.
package db

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"path"
)

// Directory structure
// repo/
//	- templates/
//	- resource/
//		- js/
//		- css/
//	- mainnet.db
type Repo struct {
	repoPath string
	resPath string
	tplPath string
	dbPath string
	DataStore *SQLiteDatastore
}

func (r *Repo) RepoPath() string {
	return r.repoPath
}

func (r *Repo) TplPath() string {
	return r.tplPath
}

func (r *Repo) ResPath() string {
	return r.resPath
}

func (r *Repo) FetchTemplate(fileNames []string) (*template.Template, error) {
	filePaths := make([]string, len(fileNames))
	for i, s := range fileNames {
		filePaths[i] = path.Join(r.tplPath, s)
	}
	return template.ParseFiles(filePaths...)
}

func InitRepo(repoPath string) (*Repo, error) {
	res, err := OpenRepo(repoPath)
	if err != nil {
		return nil, err
	}

	if fileExists(res.dbPath) {
		return nil, errors.New("datastore already exists in "+repoPath)
	}
	err = res.DataStore.InitDatabaseTables("")
	if err != nil {
		fmt.Println("Error when initialize DB: "+err.Error())
		return nil, err
	}
	return res, nil
}

func OpenRepo(repoPath string) (*Repo, error) {
	res := &Repo{
		repoPath: repoPath,
	}
	res.dbPath = path.Join(repoPath, "mainnet.db")
	if !res.check() {
		return nil, errors.New("invalid repo "+repoPath)
	}

	var err error
	res.DataStore, err = CreateDB(res.dbPath, "")
	if err != nil {
		fmt.Println("Error when create DB: "+err.Error())
		return nil, err
	}
	return res, nil
}

func (r *Repo) check() bool {
	r.resPath = path.Join(r.repoPath, "resource")
	r.tplPath = path.Join(r.repoPath, "templates")
	return directoryExists(r.resPath) && directoryExists(r.tplPath)
}

func directoryExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func fileExists(filePath string) bool {
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}
