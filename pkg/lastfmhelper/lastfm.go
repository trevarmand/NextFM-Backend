package lastfmhelper

import (
	"fmt"
	"os"

	"github.com/trevarmand/lastfm-go/lastfm"
	// "github.com/trevarmand/nextfm-backend/pkg/util/log"
)

var (
	lastFmApiKey    = os.Getenv("lastfm_api_key")
	lastFmApiSecret = os.Getenv("lastfm_api_secret")
)

/**
 Returns a link users must follow in order to authorize this project to access their account data.
**/
func FetchAccountAuthURL() string {
	api := lastfm.New(lastFmApiKey, lastFmApiSecret)
	token, _ := api.GetToken() //discarding error
	authUrl := api.GetAuthTokenUrl(token)
	fmt.Println(authUrl)
	return authUrl
}
