package db

import (
	"fmt"
	"testing"
)

func TestVideoDB(t *testing.T) {
	testDB, err := Create("testrepo", "")
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
