package db

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"sync"
)

type VideoDB struct {
	db   *sql.DB
	lock *sync.Mutex
}

func NewVideoDB(db *sql.DB, lock *sync.Mutex) *VideoDB {
	return&VideoDB{
		db:   db,
		lock: lock,
	}
}

type Video struct {
	Index int `json:"index"`
	VideoName string `json:"videoName"`
	BlockNum int `json:"blockNum"`
	Poster string `json:"poster"`
	VideoLength int64 `json:"videoLength"`
	Created int64 `json:"created"`
	Updated int64 `json:"updated"`
}

func (v *Video) ToJson() ([]byte, error) {
	return json.Marshal(v)
}

func (vs *VideoDB) GetLast() *Video {
	vs.lock.Lock()
	defer vs.lock.Unlock()
	stm := `select * From table where ind = (select max(ind) from videos))`
	videos := vs.handleQuery(stm)
	if len(videos) > 1 {
		fmt.Println("Error when fetch the last video in db: get more than one video.")
		return nil
	} else if len(videos)==0{
		return nil
	} else {
		return videos[0]
	}
}

func (vs *VideoDB) Add(video *Video) error {
	vs.lock.Lock()
	defer vs.lock.Unlock()
	fmt.Println("DB: Call add video.")
	tx, err := vs.db.Begin()
	if err != nil {
		return err
	}
	//create table videos (index integer primary key not null, name text not null, blockNum integer not null, poster text, created integer not null, updated integer not null)
	stm := `insert into videos(ind, videoName, blockNum, poster, videoLength, created, updated) values(?,?,?,?,?,?,?)`
	stmt, err := tx.Prepare(stm)
	if err != nil {
		fmt.Printf("DB: Error in prepare: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		video.Index,
		video.VideoName,
		video.BlockNum,
		video.Poster,
		video.VideoLength,
		video.Created,
		video.Updated,
	)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (vs *VideoDB) List(query string) []*Video {
	vs.lock.Lock()
	defer vs.lock.Unlock()
	q := "select * from videos"
	if query != "" {
		q += " where " + query
	}
	q += " order by updated desc;"
	return vs.handleQuery(q)
}

func (vs *VideoDB) handleQuery(stm string)[]*Video {
	list := make([]*Video, 0)
	rows, err := vs.db.Query(stm)
	if err != nil {
		fmt.Printf("DB: Error in video query: %s\n%v\n", stm, err)
		return list
	}
	for rows.Next() {
		var index int
		var videoName string
		var blockNum int
		var poster string
		var videoLength int64
		var created int64
		var updated int64
		if err := rows.Scan(&index, &videoName, &blockNum, &poster, &videoLength, &created, &updated); err != nil {
			fmt.Printf("DB: Error in db scan: %s\n", err)
			continue
		}
		list = append(list, &Video{
			Index:       index,
			VideoName:   videoName,
			BlockNum:    blockNum,
			Poster:      poster,
			VideoLength: videoLength,
			Created:     created,
			Updated:     updated,
		})
	}
	return list
}