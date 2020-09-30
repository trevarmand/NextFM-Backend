package main

import (
	"fmt"

	"github.com/trevarmand/nextfm-backend/pkg/lastfmhelper"
	"github.com/trevarmand/nextfm-backend/pkg/util/aws"
	"github.com/trevarmand/nextfm-backend/pkg/db"
	// "github.com/trevarmand/nextfm-backend/pkg/util/log"

)

func main() {
	fmt.Println("Starting NextFm Backend")

	sess, err := aws.GetSession()
	if err != nil {
		fmt.Println("Failed to establish session, exiting")
		return
	}
	fmt.Println(sess)


	db.GetConnection()
	lastfmhelper.FetchAccountAuthURL()
}
