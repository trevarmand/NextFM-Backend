package db

func (apc *AwsPsqlConnection) GetArtistUrl(artistName string) string {
	apc.dbConn.Exec("SELECT * from artist")
	return "nil"
}
