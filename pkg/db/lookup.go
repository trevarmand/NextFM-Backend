package db

import (
	"fmt"

	"github.com/trevarmand/nextfm-backend/pkg/util/log"
)

type Artist struct {
	Name        string
	SpotifyLink string
	dbId        int
	lastFmId    int
}

func (apc *AwsPsqlConnection) GetArtist(artistName string) *Artist {
	result, err := apc.dbConn.Query("select * from music.artist;")

	logLookupError("GetArtist", "Failed to execute select statement", err)

	result.Next()
	var name string
	var id int
	var link string
	result.Scan(&id, &name, &link)
	fmt.Println("GetArtistSqlIface", name, id, link)

	return &Artist{
		Name:        name,
		SpotifyLink: link,
		dbId:        id,
	}
}

func logLookupError(srcFunc, msg string, err error) {
	log.LogError("db:lookup:"+srcFunc, msg, err)
}
