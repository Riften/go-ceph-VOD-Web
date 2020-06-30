package db

import (
	"fmt"
	"os"
	"path"
	"testing"
)

func TestVideoDB(t *testing.T) {
	pwd, _ := os.Getwd()
	t.Log(pwd)
	_, err := os.Stat("testrepo")
	if err != nil {
		t.Error(err)
		return
	}

	testDB, err := CreateDB(path.Join("testrepo", "mainnet.db"), "")
	if err != nil {
		t.Error(err)
		return
	}
	err = testDB.InitDatabaseTables("")
	if err != nil {
		t.Error(err)
		return
	}

	err = testDB.Videos.Add(&Video{
		Index:       0,
		VideoName:   "test1",
		BlockNum:    0,
		Poster:      "test1",
		VideoLength: 0,
		Created:     0,
		Updated:     0,
	})
	if err != nil {
		t.Error(err)
		return
	}

	for _,v := range testDB.Videos.List("") {
		fmt.Println(v.VideoName)
	}
}
